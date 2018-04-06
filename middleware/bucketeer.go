package middleware

import (
	"github.com/bigodines/bucketeer/lib"
	"github.com/labstack/echo"
)

const (
	HeaderBucketeer = "X-Bucketeer-Id"
)

var (
	DefaultConfig = lib.Config{
		Experiments: map[string]lib.Experiment{
			"fontrolirst": lib.Experiment{
				Weight: 1.0,
				Name:   "Control group",
			},
		},
		Parser: DefaultParser,
	}

	conf lib.Config
)

func Bucketize() echo.MiddlewareFunc {
	return BucketizeWithConfig(DefaultConfig)
}

func BucketizeWithConfig(c lib.Config) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		handler := func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			id := req.Header.Get(HeaderBucketeer)
			if id != "" {
				// don't change the headers.
				res.Header().Set(HeaderBucketeer, id)

				b, err := conf.Parser(id)
				if err != nil {
					return err
				}
				// Add bucket to the context
				c.Set("bucket", b)
			}

			return next(c)
		}

		return handler
	}
}

// Default parser
func DefaultParser(bid string) (lib.Bucket, error) {
	b := lib.Bucket{
		Experiment: conf.Experiments["control"],
	}

	return b, nil
}
