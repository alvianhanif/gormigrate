// +build postgresql

package gormigrate

import (
	_ "github.com/alvianhanif/gorm/dialects/postgres"
)

func init() {
	databases = append(databases, database{
		name:    "postgres",
		connEnv: "PG_CONN_STRING",
	})
}
