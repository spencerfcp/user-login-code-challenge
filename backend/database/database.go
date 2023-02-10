package database

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type TableName string
type IDName string

var _scannerInterface = reflect.TypeOf((*sql.Scanner)(nil)).Elem()
var _valuerInterface = reflect.TypeOf((*driver.Valuer)(nil)).Elem()

var dateType = reflect.TypeOf((*time.Time)(nil))
var datePtrType = reflect.TypeOf(time.Time{})

func isSQLSupportedType(kind reflect.Type) bool {
	isScanner := reflect.PtrTo(kind).Implements(_scannerInterface)
	isValuer := reflect.PtrTo(kind).Implements(_valuerInterface)
	return kind.Kind() != reflect.Struct || kind == dateType || kind == datePtrType || (isScanner && isValuer)
}

func sqlFieldNames(t reflect.Type) []string {
	colNames := []string{}
	for i, fields := 0, t.NumField(); i < fields; i++ {
		field := t.Field(i)
		fieldType := field.Type
		if isSQLSupportedType(fieldType) {
			colNames = append(colNames, strings.ToLower(field.Name))
		} else {
			colNames = append(colNames, sqlFieldNames(fieldType)...)
		}
	}
	return colNames
}

func columnSql(s interface{}) string {
	val := reflect.Indirect(reflect.ValueOf(s))
	fields := sqlFieldNames(val.Type())

	columnNames := []string{}
	for _, f := range fields {
		columnNames = append(columnNames, fmt.Sprintf(`"%s"`, f))
	}

	return strings.Join(columnNames, ", ")
}

func valueSql(s interface{}) string {
	val := reflect.Indirect(reflect.ValueOf(s))
	fields := sqlFieldNames(val.Type())

	valueNames := make([]string, len(fields))
	for i, f := range fields {
		valueNames[i] = fmt.Sprintf(`:%s`, f)
	}

	return strings.Join(valueNames, ", ")
}

func Exists(q sqlx.Queryer, query string, args ...interface{}) bool {
	var exists bool
	existsQuery := fmt.Sprintf("select exists(%s)", query)
	err := sqlx.Get(q, &exists, existsQuery, args...)

	if err != nil {
		panic(err)
	}
	return exists
}

func Transaction(db *sqlx.DB, txFunc func(*sqlx.Tx) error) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback() //nolint:errcheck
			panic(p)
		}
	}()

	err = txFunc(tx)
	if err != nil {
		_ = tx.Rollback() //nolint:errcheck
		return err
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "commit failed")
	}
	return nil
}

// InsertStructReturningID is a helpful tool to insert a struct representing a table row with the new row ID returned
func InsertStructReturningID(ext sqlx.Ext, table TableName, idName IDName, s interface{}) (int, error) {
	query := fmt.Sprintf(`
		insert into %s (
			%s
		) values (
			%s
		)
		returning %s
	`, table, columnSql(s), valueSql(s), idName)

	rows, err := sqlx.NamedQuery(ext, query, s)
	if err != nil {
		return -1, err
	}
	defer rows.Close()

	if !rows.Next() {
		return -1, errors.New("no rows returned")
	}

	var id int
	err = rows.Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}
