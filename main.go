package main

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.elastic.co/apm/module/apmechov4"
)

func main() {
	// Echo instance
	e := echo.New()

	// // Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			if strings.Contains(c.Request().URL.Path, "swagger") {
				return true
			}
			return false
		},
	}))
	// //e.Use(middleware.Recover())
	e.Use(apmechov4.Middleware())

	// Routes
	// e.POST("/test", engine.Test)
	// e.POST("/transfer/bank/", engine.TransferBank)
	// e.POST("/inquiry/account/", engine.InquiryAccount)
	// e.POST("/inquiry/transaction/", engine.InquiryTransaction)
	// e.POST("/inquiry/balance/", engine.InquiryBalance)

	// Start server
	e.Logger.Fatal(e.Start(":1425"))
}
