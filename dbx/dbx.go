package dbx

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	once     sync.Once
	instance *sql.DB = nil
)

func Init(endpoint string, username string, password string, dbName string, maxOpenConns int, maxIdleConns int) error {
	var err error = nil

	once.Do(func() {
		address := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&multiStatements=true", username, password, endpoint, dbName)
		instance, err = sql.Open("mysql", address)
		if nil != err {
			return
		}

		instance.SetMaxOpenConns(maxOpenConns)
		instance.SetMaxIdleConns(maxIdleConns)
		instance.SetConnMaxIdleTime(time.Duration(120) * time.Second)
		instance.SetConnMaxLifetime(time.Duration(60) * time.Second)
	})

	return err
}

func New() *sql.DB {
	if nil == instance {
		return nil
	}

	return instance
}

func Close() {
	if nil != instance {
		instance.Close()
		instance = nil
	}
}

func GetData(rows *sql.Rows, object interface{}) error {
	// obtain column names
	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	// obtain required information of specific interface
	v := reflect.ValueOf(object)
	if v.Elem().Type().Kind() != reflect.Struct {
		return errors.New("illegal structure")
	}

	// declare slice of specific interface
	slice := []interface{}{}

	// declare map of SQL column addresses
	columnAddrMap := map[string]interface{}{}

	for i := 0; i < v.Elem().NumField(); i++ {
		propertyName := v.Elem().Field(i)
		columnName := v.Elem().Type().Field(i).Tag.Get("db")
		if columnName == "" {
			if v.Elem().Field(i).CanInterface() == false {
				continue
			}
			columnName = propertyName.Type().Name()
		}

		// mapping column address information
		columnAddrMap[columnName] = propertyName.Addr().Interface()
	}

	// assign values
	for _, colName := range columns {
		slice = append(slice, columnAddrMap[colName])
	}

	// perform scanning
	return rows.Scan(slice...)
}
