package container

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if err := NewMySQL().RunMigrations(); err != nil {
		os.Exit(1)
	}

	os.Exit(m.Run())
}
