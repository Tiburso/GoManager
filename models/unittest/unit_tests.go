package unittest

import (
	"testing"

	"github.com/Tiburso/GoManager/models/db"
	"github.com/stretchr/testify/assert"
)

func AssertExists[T any](t *testing.T, model T) T {
	db := db.DB

	res := db.Limit(1).Find(model)
	assert.NoError(t, res.Error)
	assert.Greater(t, res.RowsAffected, int64(0),
		"Expected to find at least one %T, but found none", model)

	return model
}

func AssertNotExists[T any](t *testing.T, model T) {
	db := db.DB

	res := db.Limit(1).Find(model)
	assert.NoError(t, res.Error)
	assert.Equal(t, int64(0), res.RowsAffected,
		"Expected to find no %T, but found at least one", model)
}
