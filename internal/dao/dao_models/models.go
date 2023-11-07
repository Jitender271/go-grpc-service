package daomodels

import (
	"github.com/gocql/gocql"
)

type Movies struct {
	MovieID    gocql.UUID	`db:"moveie_id"`
	Name       string		`db:"name"`
	Genre      string		`db:"genre"`
	Description string		`db:"description"`
	Rating     float32		`db:"rating"`
}