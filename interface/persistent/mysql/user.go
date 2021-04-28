package mysql

import "database/sql"

type UserImpl struct {
	cli *sql.DB
}

func NewUser(cli *sql.DB) *UserImpl {
	return &UserImpl{
		cli: cli,
	}
}
