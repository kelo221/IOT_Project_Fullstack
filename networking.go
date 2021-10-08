package main

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	chart "github.com/wcharczuk/go-chart/v2"
	"log"
	"os"
)

func handleHTTP() {

	app := fiber.New()
	app.Static("/", "./public")

	// Match all routes starting with /api
	app.Delete("/clearDatabase", func(c *fiber.Ctx) error {
		dropDatabase()
		return c.Next()
	})

	// GET /api/list
	app.Get("/getGraphData", func(c *fiber.Ctx) error {
		getPng()
		return c.Next()
	})

	err := app.Listen(":8080")
	if err != nil {
		logrus.Panicln(err)
	}
}

func getPng() {
	var b float64
	b = 1000

	ts1 := chart.ContinuousSeries{ //TimeSeries{
		Name:    "Time Series",
		XValues: []float64{10 * b, 20 * b, 30 * b, 40 * b, 50 * b, 60 * b, 70 * b, 80 * b},
		YValues: []float64{1.0, 2.0, 30.0, 4.0, 50.0, 6.0, 7.0, 88.0},
	}

	ts2 := chart.ContinuousSeries{ //TimeSeries{
		Style: chart.Style{
			StrokeColor: chart.GetDefaultColor(1),
		},

		XValues: []float64{10 * b, 20 * b, 30 * b, 40 * b, 50 * b, 60 * b, 70 * b, 80 * b},
		YValues: []float64{15.0, 52.0, 30.0, 42.0, 50.0, 26.0, 77.0, 38.0},
	}

	graph := chart.Chart{

		XAxis: chart.XAxis{
			Name:           "The XAxis",
			ValueFormatter: chart.TimeMinuteValueFormatter, //TimeHourValueFormatter,
		},

		YAxis: chart.YAxis{
			Name: "The YAxis",
		},

		Series: []chart.Series{
			ts1,
			ts2,
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		log.Fatal(err)
	}

	fo, err := os.Create("public/img/fanGraph.png")
	if err != nil {
		panic(err)
	}

	if _, err := fo.Write(buffer.Bytes()); err != nil {
		panic(err)
	}
}
