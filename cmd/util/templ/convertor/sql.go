package convertor

import "database/sql"

func ConvNullInt32(n int) sql.NullInt32 {
	return sql.NullInt32{
		Valid: true,
		Int32: int32(n),
	}
}
