package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"postapi/app/models"
)

type PostDB interface {
	Open() error
	Close() error

	CreatePost(p *models.Post) (int64, error)
	GetAll() ([]*models.Post, error)
}

type DB struct {
	db *sqlx.DB
}
