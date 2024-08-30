package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type SellSideMetrics struct {
	AverageMMR      float64 `json:"averageMMR"`
	AverageNetChurn float64 `json:"averageNetChurn"`
	MinNetGrowth    float64 `json:"minNetGrowth"`
	Runway          float64 `json:"runway"`
}

func HandleSellSideDemoMetrics(c echo.Context) error {
	data := SellSideMetrics{
		AverageMMR:      20512.1,
		AverageNetChurn: 4515.5,
		MinNetGrowth:    1313.5,
		Runway:          8,
	}
	return c.JSON(http.StatusOK, data)
}
