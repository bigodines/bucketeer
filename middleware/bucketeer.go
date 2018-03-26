package middleware

import "github.com/labstack/echo"

type (
	Config struct {
		Experiments map[string]Experiment
	}

	Experiment struct {
		Pct  float32
		Name string
	}

	Bucket struct {
		Experiment Experiment
	}
)

const (
	HeaderBucketeer = "X-Bucketeer-Id"
)

var (
	DefaultConfig = Config{}
)

func Bucketize() echo.MiddlewareFunc {
	return BucketizeWithConfig(DefaultConfig)
}

func BucketizeWithConfig(c Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		handler := func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			id := req.Header.Get(HeaderBucketeer)
			if id != "" {
				res.Header().Set(HeaderBucketeer, id)
			}

			return next(c)
		}

		return handler
	}
}
