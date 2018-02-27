package test

import (
	"os"
	"testing"

	"nightcode/cmd"
	"nightcode/model"
)

func TestMain(m *testing.M) {
	cmd.Execute()
	model.LinkDb()

	result := m.Run()

	model.CloseDb()
	os.Exit(result)
}
