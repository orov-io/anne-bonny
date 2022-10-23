package middleware

import (
	"context"
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/orov-io/anne-bonny/video-streamer/model"
	"github.com/orov-io/maryread"
)

type InjectFactoryConfig struct {
	// Skipper defines a function to skip middleware.
	Skipper middleware.Skipper

	// BeforeFunc defines a function which is executed just before the middleware.
	BeforeFunc middleware.BeforeFunc
}

const factoryContextKey = "factory"

var (
	DefaultFactoryConfig = InjectFactoryConfig{
		Skipper: middleware.DefaultSkipper,
	}
)

var (
	ErrFactoryNotFound = errors.New("unable to find factory in provided context")
)

func InjectFactory() echo.MiddlewareFunc {
	return InjectFactoryWithConfig(DefaultFactoryConfig)
}

func InjectFactoryWithConfig(config InjectFactoryConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = DefaultFactoryConfig.Skipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			if config.BeforeFunc != nil {
				config.BeforeFunc(c)
			}

			dbx := maryread.MustGetDBX(c)
			factory := model.NewFactory(context.Background(), dbx)

			c.Set(factoryContextKey, factory)
			return next(c)
		}
	}
}

func GetFactory(c echo.Context) (*model.Factory, error) {
	factory, ok := c.Get(factoryContextKey).(*model.Factory)
	if !ok {
		return nil, ErrFactoryNotFound
	}

	return factory, nil
}

func MustGetFactory(c echo.Context) *model.Factory {
	factory, err := GetFactory(c)
	if err != nil {
		panic(err)
	}

	return factory
}
