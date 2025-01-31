package mockdb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockDB(t *testing.T) {
	gormDB, mock, db := New()
	defer db.Close()

	assert.NotNil(t, gormDB)
	assert.NotNil(t, mock)
	assert.NotNil(t, db)
}
