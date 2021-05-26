package database

import (
	"github.com/jmoiron/sqlx"
	"log"
	"postapi/app/models"
)

func (db *DB) Open() error {
	connection, err := sqlx.Open("postgres", pgConnStr)
	if err != nil {
		return err
	}

	log.Println("Connected to Database(Postgres)")
	connection.MustExec(createSchema)

	db.db = connection

	return nil
}

func (db *DB) Close() error {
	return db.db.Close()
}

func (db *DB) CreatePost(p *models.Post) (int64, error) {
	var id int64
	err := db.db.QueryRow(InsertPostSchema, p.Title, p.Content, p.Author).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (db *DB) GetAll() ([]*models.Post, error) {
	var all []*models.Post
	if err := db.db.Select(all, "SELECT id, title, content, author FROM posts"); err != nil {
		return nil, err
	} else {
		return all, nil
	}
}
