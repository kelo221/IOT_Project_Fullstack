package main

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

func handleHTTP() {

	app := fiber.New()
	app.Static("/", "./public")

	app.Delete("/clearDatabase", func(c *fiber.Ctx) error {
		fmt.Println("Requested to clear collection")
		dropDatabase()
		return c.SendString("Data deleted")
	})

	app.Post("/getUserSettings", func(c *fiber.Ctx) error {
		fmt.Println("User data sent")
		//fmt.Println(string(c.Body()))

		payload := struct {
			Auto     bool `json:"auto,omitempty"`
			Pressure int  `json:"pressure,omitempty"`
			Speed    int  `json:"speed,omitempty"`
		}{}

		p := payload
		if err := c.BodyParser(&p); err != nil {
			fmt.Println(err)
			return err
		}

		newAuto = p.Auto
		newPressure = p.Pressure
		newSpeed = p.Speed

		fmt.Println(p)
		return c.SendString("Data OK")
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

	//fmt.Println(dataPayload)

	for _, s := range dataPayload {
		pressureData = append(pressureData, opts.LineData{Value: s.Pressure})
		speedData = append(speedData, opts.LineData{Value: s.Speed})
		timeData = append(timeData, s.UnixTime)
	}

	///fmt.Println("graph requested")

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

// generate random data for line chart
func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}
