package dbutils

import (
	"database/sql"
	"fmt"

	go_mysql "github.com/go-sql-driver/mysql"
	my_mysql "github.com/ziutek/mymysql/mysql"
)

//CheckAffected checks if result.RowsAffected() are equal to the expected numeber
func CheckAffected(result sql.Result, sqlError error, expected int64) error {
	if sqlError != nil {
		return sqlError
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n != expected {
		return fmt.Errorf("bad affected count: %d != %d", n, expected)
	}
	return nil
}

//IsMySQLDuplicate checks if mysql error is ER_DUP_ENTRY mysql error
func IsMySQLDuplicate(err error) bool {
	if val, ok := err.(*my_mysql.Error); ok && val.Code == my_mysql.ER_DUP_ENTRY {
		return true
	}
	if val, ok := err.(*go_mysql.MySQLError); ok && val.Number == 1062 {
		return true
	}
	return false
}
