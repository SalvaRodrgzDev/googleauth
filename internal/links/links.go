package links

import (
	database "googleauth/internal/pkg/db/migrations/postgres"
	"googleauth/internal/users"
	"strconv"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    users.User
}

func (link Link) Save() int64 {

	newLink := Link{Title: link.Title, Address: link.Address}

	result := database.Connection().Create(&newLink)

	if result.Error != nil {
		panic(result.Error)
	}

	id, _ := strconv.ParseInt(newLink.ID, 10, 64)

	return id
}
