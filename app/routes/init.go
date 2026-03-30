package routes

import (
	"fmt"
	"log"
	"roadmap/pkg/humax"
	"time"

	"github.com/danielgtaylor/huma/v2/adapters/humaecho"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	if err := initHTTPServer(e); err != nil {
		log.Fatal("init http server: %w", err)
	}

	return e
}

func initHTTPServer(e *echo.Echo) error {
	// quyen@note: middleware here

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod:  true,
		LogStatus:  true,
		LogURI:     true,
		LogLatency: true,
		LogError:   true,

		LogValuesFunc: func(_ echo.Context, v middleware.RequestLoggerValues) error {
			fmt.Printf("[%s] %d %s %s %v\n",
				time.Now().Format(time.RFC3339),
				v.Status,
				v.Method,
				v.URI,
				v.Latency,
			)
			return nil
		},
	}))
	return registerHumaAPI(e)
}

func registerHumaAPI(e *echo.Echo) error {
	config := humax.DefaultConfig()

	humaAPI := humaecho.New(e, config)
	publicAPIv1 := humax.NewAPI(humaAPI, "/v1")

	if err := registerPublicAPIv1(publicAPIv1); err != nil {
		return fmt.Errorf("register public api v1: %w", err)
	}

	return nil
}
