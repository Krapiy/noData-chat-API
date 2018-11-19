package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const (
	errorTableCreate = "cannot create %s table"
)

// MysqlDB connect to mysql DB
type MysqlDB struct {
	Conn *sqlx.DB
}

// New create connect to DB and retrun client
func New(address string) (*MysqlDB, error) {
	db, err := sqlx.Open("mysql", address)
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to database")
	}

	client := &MysqlDB{db}

	err = client.createUsersTable()
	if err != nil {
		return nil, errors.Wrapf(err, errorTableCreate, "users")
	}

	err = client.createRoomsTable()
	if err != nil {
		return nil, errors.Wrapf(err, errorTableCreate, "rooms")
	}

	err = client.createMessagesTable()
	if err != nil {
		return nil, errors.Wrapf(err, errorTableCreate, "messages")
	}

	return client, nil
}

func (c *MysqlDB) createUsersTable() error {
	sqlUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			user_name VARCHAR(255) NOT NULL UNIQUE,
			password_hash VARCHAR(255) NOT NULL,
			pub_key VARCHAR(255) NOT NULL,
			config BLOB NULL,
			INDEX(user_name)
		)
	`

	_, err := c.Conn.Exec(sqlUsersTable)
	return err
}

func (c *MysqlDB) createRoomsTable() error {
	sqlRoomsTable := `
		CREATE TABLE IF NOT EXISTS rooms (
			id SERIAL PRIMARY KEY,
			room_id BIGINT UNSIGNED NOT NULL,
			user_id BIGINT UNSIGNED NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
			INDEX(room_id)
		)
	`

	_, err := c.Conn.Exec(sqlRoomsTable)
	return err
}

func (c *MysqlDB) createMessagesTable() error {
	sqlMessagesTable := `
		CREATE TABLE IF NOT EXISTS messages (
			id SERIAL PRIMARY KEY,
			user_sender_id BIGINT UNSIGNED NOT NULL,
			user_receiver_id BIGINT UNSIGNED DEFAULT 0,
			room_id BIGINT UNSIGNED DEFAULT 0,
			FOREIGN KEY (user_sender_id) REFERENCES users(id),
			FOREIGN KEY (user_receiver_id) REFERENCES users(id),
			FOREIGN KEY (room_id) REFERENCES rooms(room_id)
		)
	`
	_, err := c.Conn.Exec(sqlMessagesTable)
	return err
}