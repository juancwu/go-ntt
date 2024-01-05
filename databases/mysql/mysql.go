package mysql

import (
	"database/sql"

	"github.com/juancwu/go-ntt/databases"
)

var _ databases.Driver = (*Mysql)(nil)

type Mysql struct {
	db *sql.DB
}

func init() {
	databases.Register("mysql", &Mysql{})
}

func (m *Mysql) Open(url string) (databases.Driver, error) {
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}

	mx := &Mysql{db: db}

	return mx, nil
}

func (m *Mysql) Run(msg chan databases.Migration) error {
	return nil
}

func (m *Mysql) Up(source string) error {
	return nil
}

func (m *Mysql) Down(source string) error {
	return nil
}
