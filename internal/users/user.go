package users

import (
	"database/sql"
)

type User struct {
	Id       uint
	Login    string
	Password string
	Name     sql.NullString
	Surname  sql.NullString
	Birthday sql.NullString
	Mail     sql.NullString
	Phone    sql.NullString
}
