package database

import (
	"database/sql"
	"github.com/gorilla/schema"
	"reflect"
)

func SQLDecoder() *schema.Decoder {
	decoder := schema.NewDecoder()
	SchemaRegisterSQLNulls(decoder)
	decoder.IgnoreUnknownKeys(true)
	return decoder
}

func SchemaRegisterSQLNulls(d *schema.Decoder) {
	nullString := sql.NullString{}
	nullBool := sql.NullBool{}
	nullInt32 := sql.NullInt32{}
	nullInt64 := sql.NullInt64{}
	nullFloat64 := sql.NullFloat64{}

	d.RegisterConverter(nullString, ConvertSQLNullString)
	d.RegisterConverter(nullBool, ConvertSQLNullBool)
	d.RegisterConverter(nullInt32, ConvertSQLNullInt32)
	d.RegisterConverter(nullInt64, ConvertSQLNullInt64)
	d.RegisterConverter(nullFloat64, ConvertSQLNullFloat64)
}

func ConvertSQLNullString(value string) reflect.Value {
	v := sql.NullString{}
	if err := v.Scan(value); err != nil {
		return reflect.Value{}
	}

	return reflect.ValueOf(v)
}

func ConvertSQLNullBool(value string) reflect.Value {
	v := sql.NullBool{}
	if err := v.Scan(value); err != nil {
		return reflect.Value{}
	}

	return reflect.ValueOf(v)
}

func ConvertSQLNullInt64(value string) reflect.Value {
	v := sql.NullInt64{}
	if err := v.Scan(value); err != nil {
		return reflect.Value{}
	}

	return reflect.ValueOf(v)
}

func ConvertSQLNullInt32(value string) reflect.Value {
	if value == "" {
		return reflect.ValueOf(sql.NullInt32{
			Valid: false,
		})
	}
	v := sql.NullInt32{}
	if err := v.Scan(value); err != nil {
		return reflect.Value{}
	}

	return reflect.ValueOf(v)
}

func ConvertSQLNullFloat64(value string) reflect.Value {
	v := sql.NullFloat64{}
	if err := v.Scan(value); err != nil {
		return reflect.Value{}
	}

	return reflect.ValueOf(v)
}
