package main

import (
	"go-front-skel-001/config"
	"go-front-skel-001/routes"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	// "github.com/gofiber/template/html/v2"
	// "go-front-skel-001/components/index"
	// "github.com/gofiber/fiber/v2/middleware/adaptor"

 //        "github.com/a-h/templ"
	"github.com/gofiber/fiber/v2/middleware/favicon"
)


// func handleRoot(c *fiber.Ctx) error {
// 	// return c.SendString("api function")
// 	// return c.Render("index", fiber.Map{
// 	// 	// return c.Render("<h1>{{ .Title }}</h1>", fiber.Map{
// 	// 	"name": "My GO ap1 v1xx",
// 	// })
// 	component := common.Header("kanibal")
// 	// return component.Render(context.Background(), )
// 	return router.Get("/hello", adaptor.HTTPHandler(templ.Handler(view.Hello("John Doe"))))
// }

func main() {
    l := config.Logger()
    l.Info("starting server")
    l.Info("loading config")
    // logger.Info("Starting oven!", "degree", 375)
    // time.Sleep(10 * time.Second)
    // logger.Info("Finished baking")
    // log.SetLevel(log.DebugLevel)
    // log.Debug("Cookie 🍪")
    // log.Info("Finished baking")
    // log.Info(time.Now())
    var err = config.LoadEnv()
    if err != nil {
	    // log.Fatal("error loading config")
	    l.Fatal("configLoaded error: %v", err)
    } else {
	    l.Info("configLoaded db: " + config.Conf["db_name"].(string))
    }

    // config server

    // inject your custom logger
    // engine := html.New("./views", ".html")
    app := fiber.New( fiber.Config{
		// ErrorHandler: CustomErrorHandler,
		// Views: engine,
		Prefork:       true,
		AppName: "Skel App v0.0.1",
    })

    app.Static("/static", "./static", fiber.Static{
	Compress:      false,
	ByteRange:     false,
	Browse:        true,
	Index:         "index.html",
	CacheDuration: 10 * time.Second,
	MaxAge:        3600,
    })
    // Add Custom Tags
    app.Use(logger.New(logger.Config{
	TimeFormat: time.RFC3339,
	TimeZone: "UTC",
	Format: "${time}${cyan} INFO <REQUEST>${reset} log: ${status} - req_id: ${locals:requestid} - ${latency} - ${method} - ${path}\n",
    }))
    app.Use(requestid.New())
    app.Use(cors.New(cors.Config{
	AllowHeaders: "Origin, Content-type, Accept, Authorization",
	AllowCredentials: true,
    }))
    app.Use(favicon.New(favicon.Config{
	    File: "./static/favicon.ico",
	    URL:  "/favicon.ico",
    }))

    // app.Get("/", handleRoot )
    // app.Get("/",adaptor.HTTPHandler(templ.Handler(index.Index("John Doe"))))
    routes.Setup(app) // load api routes

    app.Use(func(c *fiber.Ctx) error {
	    return c.Status(fiber.StatusNotFound).SendString("path not found")
    })

    // returns the original router stack
    // data, _ := json.MarshalIndent(app.Stack(), "", "  ")
    // fmt.Println(string(data))

    l.Fatal(app.Listen(":3030"))
}
