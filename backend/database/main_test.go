package database

import (
	"backend/common"
	"fmt"
	"os"
	"testing"
)

func setupUnitTests() {
	fmt.Println("Setup database unit tests")
	testInit()
	ConnectDB()

	if err := loadAndExecuteSqlScript(common.CreateTestDatabaseSqlScriptPath); err != nil {
		panic(err)
	}
	if err := loadAndExecuteSqlScript(common.InsertDataIntoDatabaseSqlScriptPath); err != nil {
		panic(err)
	}
}

func tearDownUnitTests() {
	fmt.Println("Teardown after database unit tests")
	if err := loadAndExecuteSqlScript(common.DropTestDatabaseSqlScriptPath); err != nil {
		panic(err)
	}
	DisconnectDB()
}

func TestMain(m *testing.M) {
	setupUnitTests()

	exitCode := m.Run()

	tearDownUnitTests()

	os.Exit(exitCode)
}
