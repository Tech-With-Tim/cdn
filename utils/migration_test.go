package utils

import (
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/stretchr/testify/require"
)

func TestMigrateUp(t *testing.T) {
	err := MigrateUp(config, "../models/migrations/")
	if err != nil {
		if err != migrate.ErrNoChange {
			require.NoError(t, err)
		}
	}
}

func TestMigrateDown(t *testing.T) {
	err := MigrateDown(config, "../models/migrations/")
	require.NoError(t, err)
}

func TestMigrateSteps(t *testing.T) {
	err := MigrateSteps(1, config, "../models/migrations/")
	require.NoError(t, err)
}
