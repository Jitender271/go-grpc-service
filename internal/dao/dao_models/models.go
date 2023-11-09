package daomodels

type Movies struct {
	MovieID     string `db:"movie_id"`
	Name        string `db:"name"`
	Genre       string `db:"genre"`
	Description string `db:"description"`
	Rating      string `db:"rating"`
}