package server

import (
	"context"
	"errors"
	"os"
	"os/signal"

	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/ryanseipp/otel"
)

func RunServer() (err error) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	serviceName := "go-api"
	shutdown, err := InitTracing(ctx, serviceName)
	if err != nil {
		return
	}

	defer func() {
		err = errors.Join(err, shutdown(context.Background()))
	}()

	app := fiber.New()
	app.Use(otelfiber.Middleware())

	serverError := make(chan error, 1)
	go func() {
		serverError <- app.Listen(":3000")
	}()

	select {
	case err = <-serverError:
		return
	case <-ctx.Done():
		stop()
	}

	err = app.Shutdown()
	return
}
