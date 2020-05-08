package dbclient

import (
	"model/config"

	"github.com/YunzhanghuOpen/glog"
	"github.com/jmoiron/sqlx"
)

var (
	teacherEngine *sqlx.DB
	studentEngine *sqlx.DB
)

func init() {
	teacherEngine = initDB(config.MysqlTeacherConfig)
	studentEngine = initDB(config.MysqlStuConfig)
}

func initDB(mysql config.Mysql) *sqlx.DB {
	db, err := sqlx.Connect("mysql", mysql.Server)
	if err != nil {
		glog.Fatalf("Initialization mysql driver failed, err: %v", err)
	}
	db.SetMaxIdleConns(mysql.MaxIdle)
	db.SetMaxOpenConns(mysql.MaxOpen)
	db.SetConnMaxLifetime(mysql.ConnMaxLifetime)
	return db
}

func GetTeacherEngine() *sqlx.DB {
	return teacherEngine
}

func GetStudentEngine() *sqlx.DB {
	return studentEngine
}
