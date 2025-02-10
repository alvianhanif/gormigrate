// +build sqlite

package gormigrate

import (
	_ "github.com/alvianhanif/gorm/dialects/sqlite"
)

func init() {
	databases = append(databases, database{
		name:    "sqlite3",
		connEnv: "SQLITE_CONN_STRING",
	})
}
