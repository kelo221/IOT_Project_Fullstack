package main

import (
	"crypto/subtle"
	"encoding/hex"
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	_ "github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

func handleHTTP() {

	app := fiber.New()
	// Provide a minimal config

	app.Use(basicauth.New(basicauth.Config{
		Authorizer: func(user, pass string) bool {
			salt := []byte("salt")

			//Maybe this will prevent injections?
			user = strings.Replace(user, " ", "", -1)

			passwordHash := hex.EncodeToString(HashPassword([]byte(pass), salt))
			expectedPasswordHash := aqlToString("FOR r IN IOT_DATA_LOGIN FILTER r.username == \"" + user + "\" RETURN r.hash")

			usernameMatch := subtle.ConstantTimeCompare([]byte(passwordHash[:]), []byte((expectedPasswordHash[:]))) == 1

			if usernameMatch {
				return true
			}

			return false
		},
	}))

	app.Static("/", "./public")

	app.Delete("/clearDatabase", func(c *fiber.Ctx) error {
		fmt.Println("Requested to clear collection")
		dropDatabase()
		return c.SendStatus(200)
	})

	app.Post("/getUserSettings", func(c *fiber.Ctx) error {
		fmt.Println("User data sent")
		fmt.Println(string(c.Body()))

		if (string(c.Body())) == "" {
			return c.SendStatus(404)
		}

		handleMQTTOut(string(c.Body()))
		return c.SendStatus(200)
	})

	app.Get("/graphRender", graphRender)

	err := app.Listen(":8080")
	if err != nil {
		logrus.Panicln(err)
	}
}

func graphRender(c *fiber.Ctx) error {

	dataPayload := aql("FOR x IN IOT_DATA_SENSOR RETURN x")
	var timeData []int
	speedData := make([]opts.LineData, 0)
	pressureData := make([]opts.LineData, 0)

	for _, s := range dataPayload {
		pressureData = append(pressureData, opts.LineData{Value: s.Pressure})
		speedData = append(speedData, opts.LineData{Value: s.Speed})
		timeData = append(timeData, s.UnixTime)
	}

	currentTime := time.Now()
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Pressure and Fan Speed Data",
			Subtitle: "Last updated: " + currentTime.Format("2006-01-02 15:04:05"),
		}))

	// Put data into instance
	line.SetXAxis(timeData). //	int array of sample times
					AddSeries("Fan Speed", speedData).   //	pressure data
					AddSeries("Pressure", pressureData). //	fan speed data
					SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	err := line.Render(c)
	if err != nil {
		return err
	}

	return nil
}
