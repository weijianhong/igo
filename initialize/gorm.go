package initialize

import (
	"app/global"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	case "sqlite":
		return GormSqlite()
	default:
		return GormMysql()
	}
}
