package links

import (
	database "googleauth/internal/pkg/db/migrations/postgres"
	"googleauth/internal/users"
	"log"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

func (link Link) Save() int64 {

	id := 0
	err := database.Conn.QueryRow(
		"INSERT INTO Links(Title,Address) VALUES($1,$2) RETURNING id",
		link.Title,
		link.Address,
	).Scan(&id)

	if err != nil {
		log.Fatal(err)
	}

	return id
}
