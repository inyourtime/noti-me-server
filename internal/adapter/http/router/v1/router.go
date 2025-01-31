package v1

import (
	"github.com/inyourtime/noti-me-server/internal/adapter/http/handler"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRouter(app *echo.Group, db *gorm.DB) {
	healthzHandler := handler.NewHealthzHandler(db)
	app.GET("/healthz", healthzHandler.Healthz)
}
