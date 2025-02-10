package api

import (
	"errors"
	"fmt"

	"github.com/EmreKb/fiber-boilerplate/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Bootstrap() {
	app := fiber.New(fiber.Config{
		ErrorHandler: defaultErrorHandler,
	})
	app.Use(compress.New(compress.Config{Level: compress.LevelBestSpeed}))
	app.Use(cors.New(cors.ConfigDefault))
	app.Use(healthcheck.New(healthcheck.ConfigDefault))
	app.Use(helmet.New())
	app.Use(logger.New(logger.Config{
		Format:     "${time} ${status} - ${method} ${path}\n",
		TimeFormat: "15:04:05",
	}))
	app.Use(recover.New())
	app.Use(requestid.New(requestid.ConfigDefault))
	limiterConfig := limiter.ConfigDefault
	app.Use(limiter.New(limiterConfig))

	api := app.Group("/api")
	api.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	v1 := api.Group("/v1")
	v1.Get("/ping", Ping)

	if err := app.Listen(fmt.Sprintf(":%s", config.Get("PORT"))); err != nil {
		panic(err)
	}
}

func defaultErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	return c.Status(code).JSON(ErrorResponse("", "Internal Server Error"))
}
