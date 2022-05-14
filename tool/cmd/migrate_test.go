package cmd

import (
	"log"
	"path/filepath"
	"runtime"
	"testing"
)

func TestGetModuleCmd(t *testing.T) {
	cmd := GetMigration()
	cmd.Execute()
}

func TestProcessMigrate(t *testing.T) {
	_, b, _, _ := runtime.Caller(0)
	// Root folder of this project
	Root := filepath.Join(filepath.Dir(b), "..")
	log.Println(Root)

	//Columns: C�DIGO PRECIO DIRECCI� NUMERO DEUDA PRECIO2 REPARTO
	//ProcessMigrateData(Root + "\\database\\dbfMigrateTest\\registro11.dbf")
	ProcessMigrateClientData(Root+"\\database\\dbfMigrateTest\\registro10.DBF", 10)
}
