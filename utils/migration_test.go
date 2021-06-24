package utils

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMigrateUp(t *testing.T) {
	err := MigrateUp(config, "../db/migration/")
	if err != nil {
		if err != migrate.ErrNoChange {
			require.NoError(t, err)
		}
	}
}

func TestMigrateDown(t *testing.T) {
	err := MigrateDown(config, "../db/migration/")
	require.NoError(t, err)
}

func TestMigrateSteps(t *testing.T) {
	err := MigrateSteps(1, config, "../db/migration/")
	require.NoError(t, err)
}
