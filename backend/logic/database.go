package logic

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/Ryan-Campbell-PT/Sight/backend/util"
	"github.com/go-sql-driver/mysql"
)

// this dbOnce variable makes it so no matter how many times you call the function getDatabase()
// the code inside will only run once
var (
	db     *sql.DB
	dbOnce sync.Once
)

func GetDatabase() *sql.DB {
	return getSqlDatabase()
}

func getSqlDatabase() *sql.DB {
	functionName := "getSqlDatabase/"

	dbOnce.Do(func() {
		cfg := util.GetEnvConfig()
		mySqlCfg := mysql.Config{
			User:                 cfg.User,
			Passwd:               cfg.Password,
			Net:                  "tcp",
			Addr:                 "127.0.0.1:3306",
			DBName:               cfg.Azure_Database,
			AllowNativePasswords: true,
		}
		dbObj, err := sql.Open("mysql", mySqlCfg.FormatDSN())
		if util.HandleError(functionName+"Error opening db obj: ", err) {
		}

		pingErr := dbObj.Ping()
		if util.HandleError(functionName+"Error pinging db: ", pingErr) {
		}
		db = dbObj
	})

	return db
}

func getMsSqlConnectionString() string {
	cfg := util.GetEnvConfig()

	return fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		cfg.Azure_Server, cfg.Azure_User, cfg.Azure_Password, cfg.Azure_Port, cfg.Azure_Database)
}

func getMsSqlDatabase() *sql.DB {
	functionName := "getMsSqlDatabase/"
	dbOnce.Do(func() {
		dbObj, err := sql.Open("sqlserver", getMsSqlConnectionString())
		if util.HandleError(functionName+"Error connecting to MsSql db: ", err) {
			return
		}

		db = dbObj
	})

	return db
}
