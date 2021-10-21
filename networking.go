package main

import (
	"bytes"
	"crypto/subtle"
	"encoding/hex"
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	_ "github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/template/html"
	"html/template"
	"io"
	"log"
	"strings"
	"time"
)

var currentFanspeed int
var currentPressure int

type snippetRenderer struct {
	c      interface{}
	before []func()
}

func handleHTTP() {

	var currentUser string

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
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

				if currentUser != user {
					fmt.Println(user)
					aqlNoReturn("INSERT {    user: \"" +
						user +
						"\",    time: DATE_NOW()} IN IOT_DATA_LOGS")
				}
				currentUser = user
				return true
			}

			return false
		},
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Graph": graphRender(),
		})
	})

	// Serve static assets
	app.Static("/public", "./public", fiber.Static{
		Compress: true,
	})

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

	app.Get("/userLogs", userLogs)

	app.Get("/getGaugeData", gaugeData)

	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}

func gaugeData(ctx *fiber.Ctx) error {
	return ctx.JSON(MQTTpackage)
}

func userLogs(ctx *fiber.Ctx) error {
	return ctx.JSON(getDBLogs())
}

func graphRender() template.HTML {

	dataPayload := aqlMQTT("FOR x IN IOT_DATA_SENSOR RETURN x")

	var timeString []string
	speedData := make([]opts.LineData, 0)
	pressureData := make([]opts.LineData, 0)

	for _, s := range dataPayload {
		pressureData = append(pressureData, opts.LineData{Value: s.Pressure})
		speedData = append(speedData, opts.LineData{Value: s.Speed})

	}

	for _, s := range dataPayload {
		t := time.Unix(int64(s.UnixTime), 0)
		strDate := t.Format("01-02 15:04:05")
		timeString = append(timeString, strDate)
	}

	currentTime := time.Now()

	// initialize
	line := charts.NewLine()
	line.Renderer = newSnippetRenderer(line, line.Validate)

	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Pressure and Fan Speed Data",
			Subtitle: "Last updated: " + currentTime.Format("2006-01-02 15:04:05"),
		}))

	// Put data into instance
	line.SetXAxis(timeString). //	int array of sample times
					AddSeries("Fan Speed", speedData).   //	pressure data
					AddSeries("Pressure", pressureData). //	fan speed data
					SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	// generate chart and write it to io.Writer

	var htmlSnippet template.HTML = renderToHtml(line)

	return htmlSnippet
}

func renderToHtml(c interface{}) template.HTML {
	var buf bytes.Buffer
	r := c.(Renderer)
	err := r.Render(&buf)
	if err != nil {
		log.Printf("Failed to render chart: %s", err)
		return ""
	}

	return template.HTML(buf.String())
}

// Renderer
// Any kinds of charts have their render implementation and
// you can define your own render logic easily.
type Renderer interface {
	Render(w io.Writer) error
}

func newSnippetRenderer(c interface{}, before ...func()) Renderer {
	return &snippetRenderer{c: c, before: before}
}

func (r *snippetRenderer) Render(w io.Writer) error {
	const tplName = "chart"
	for _, fn := range r.before {
		fn()
	}

	tpl := template.
		Must(template.New(tplName).
			Funcs(template.FuncMap{
				"safeJS": func(s interface{}) template.JS {
					return template.JS(fmt.Sprint(s))
				},
			}).
			Parse(baseTpl),
		)

	err := tpl.ExecuteTemplate(w, tplName, r.c)
	return err
}

var baseTpl = `
<div class="container" id="graphCont">
    <div class="item" id="{{ .ChartID }}" style="width:{{ .Initialization.Width }};height:{{ .Initialization.Height }};"></div>
</div>
{{- range .JSAssets.Values }}
   <script src="{{ . }}"></script>
{{- end }}
<script type="text/javascript">
    "use strict";
    let goecharts_{{ .ChartID | safeJS }} = echarts.init(document.getElementById('{{ .ChartID | safeJS }}'), "{{ .Theme }}");
    let option_{{ .ChartID | safeJS }} = {{ .JSON }};
    goecharts_{{ .ChartID | safeJS }}.setOption(option_{{ .ChartID | safeJS }});
    {{- range .JSFunctions.Fns }}
    {{ . | safeJS }}
    {{- end }}
</script>
`
