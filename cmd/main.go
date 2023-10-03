package main

import (
	"context"
	"fmt"
	"time"

	"github.com/caarlos0/env/v9"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"rinha-de-backend/pkg/cache"
	"rinha-de-backend/pkg/database"
	"rinha-de-backend/src/handlers"
	"rinha-de-backend/src/handlers/person_handler"
	"rinha-de-backend/src/interfaces"
	"rinha-de-backend/src/repositories/person_repository"
	"rinha-de-backend/src/services/person_service"
)

var cfg Config

type Config struct {
	Port int `env:"PORT" envDefault:"3000"`
}

func GetConfig() Config {
	return cfg
}

func CreateServer() *fiber.App {
	app := fiber.New(fiber.Config{})

	return app
}

func NewServer(app *fiber.App, lc fx.Lifecycle) *fiber.App {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			errChan := make(chan error)

			go func() {
				errChan <- app.Listen(fmt.Sprintf(":%d", cfg.Port))
			}()

			select {
			case err := <-errChan:
				return err
			case <-time.After(100 * time.Millisecond):
				return nil
			}
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})

	return app
}

func main() {
	cfg = Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fx.New(
		fx.Provide(
			CreateServer,
			fx.Annotate(
				handlers.New,
				fx.ParamTags(`group:"routes"`),
			),
			fx.Annotate(
				person_handler.New,
				fx.As(new(interfaces.Handler)),
				fx.ResultTags(`group:"routes"`),
			),
			fx.Annotate(
				person_service.New,
				fx.As(new(interfaces.PersonService)),
			),
			fx.Annotate(
				person_repository.New,
				fx.As(new(interfaces.PersonRepository)),
			),
			database.NewPostgres,
			cache.NewRedis,
		),
		fx.Invoke(
			handlers.Routes,
			NewServer,
		),
	).Run()
}
