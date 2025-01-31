package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type HealthzHandler struct {
	db *gorm.DB
}

func NewHealthzHandler(db *gorm.DB) *HealthzHandler {
	return &HealthzHandler{db: db}
}

func (h *HealthzHandler) Healthz(c echo.Context) error {
	err := h.db.Exec("SELECT 1").Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}
