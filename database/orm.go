package database

import (
	"beego-poc/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs" // Import Beego logs
	_ "github.com/mattn/go-sqlite3"       // Import the SQLite3 driver
)

func registerORM() {
	orm.RegisterModel(&models.User{})
	orm.RegisterDataBase("default", "sqlite3", "data.db")

	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		logs.Critical("Failed to sync database: %v", err)
		panic(err)
	}
	orm.NewOrm()
	logs.Info("Database synchronized")
}
