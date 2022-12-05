package middleware

import (
	"fmt"
	"net"
	"net/http"

	"github.com/labstack/echo/v4"
)

func IpFilter(subnet *net.IPNet) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			ip := net.ParseIP(req.Header.Get("X-Real-IP"))
			if !subnet.Contains(ip) {
				return echo.NewHTTPError(http.StatusForbidden, fmt.Sprintf("ip %s is not from trusted subnet", ip.String()))
			}
			return next(c)
		}
	}
}
