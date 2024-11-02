package database

import (
	"beego-poc/models"
	"fmt"
	"strings"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs" // Import Beego logs
	_ "github.com/mattn/go-sqlite3"       // Import the SQLite3 driver
)

var sqliteParams = map[string]string{
	"_journal_mode": "WAL",
	"_synchronous":  "NORMAL",
	"_cache_size":   "-32000",
	"_foreign_keys": "ON",
	"_busy_timeout": "5000",
	"_temp_store":   "MEMORY",
	"_mmap_size":    "268435456",
}

var pragmas = []string{
	"PRAGMA journal_mode=WAL;",
	"PRAGMA synchronous=NORMAL;",
	"PRAGMA cache_size=-32000;",
	"PRAGMA foreign_keys=ON;",
	"PRAGMA temp_store=MEMORY;",
	"PRAGMA mmap_size=268435456;",
}

func buildConnectionString(dbPath string, params map[string]string) string {
	var pairs []string
	for key, value := range params {
		pairs = append(pairs, fmt.Sprintf("%s=%s", key, value))
	}
	return fmt.Sprintf("%s?%s", dbPath, strings.Join(pairs, "&"))
}

func registerORM() {
	registerModels()
	buildORM()
}

func registerModels() {
	orm.RegisterModel(&models.User{})
}

func buildORM() {
	connStr := buildConnectionString("./data.db", sqliteParams)

	orm.RegisterDataBase("default", "sqlite3", connStr)

	o := orm.NewOrm()
	for _, pragma := range pragmas {
		_, err := o.Raw(pragma).Exec()
		if err != nil {
			logs.Critical("Failed to execute pragma: %v", err)
			panic(err)
		}
	}

	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		logs.Critical("Failed to sync database: %v", err)
		panic(err)
	}
}
