package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/inyourtime/noti-me-server/internal/adapter/http/handler"
	mockdb "github.com/inyourtime/noti-me-server/internal/adapter/repository/mock_db"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealthz(t *testing.T) {
	gormDB, mock, db := mockdb.New()
	defer db.Close()

	t.Run("Healthz success", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest("GET", "/healthz", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mock.ExpectExec("SELECT 1").WillReturnResult(sqlmock.NewResult(1, 1))

		h := handler.NewHealthzHandler(gormDB)

		if assert.NoError(t, h.Healthz(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("Healthz error", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest("GET", "/healthz", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mock.ExpectExec("SELECT 1").WillReturnError(errors.New("some error"))

		h := handler.NewHealthzHandler(gormDB)

		if assert.NoError(t, h.Healthz(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})
}
