package requestid

import (
	"log/slog"

	"github.com/labstack/echo/v4"
)

func SetRequestID(base *slog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := newContext(c.Request().Context(), base.With("request_id", c.Response().Header().Get(echo.HeaderXRequestID)))
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}
