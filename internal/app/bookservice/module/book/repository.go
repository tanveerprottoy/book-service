package book

import (
	"github.com/jmoiron/sqlx"
	"github.com/tanveerprottoy/book-service/internal/app/bookservice/module/book/entity"
	"github.com/tanveerprottoy/book-service/pkg/data/sqlxpkg"
)

type Repository[T entity.Book] struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository[entity.Book] {
	r := new(Repository[entity.Book])
	r.db = db
	return r
}

func (r *Repository[T]) Create(e *entity.Book) error {
	var lastId string
	err := r.db.QueryRow("INSERT INTO Books (title, author, publication_year, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id", e.Title, e.Author, e.PublicationYear, e.CreatedAt, e.UpdatedAt).Scan(&lastId)
	if err != nil {
		return err
	}
	e.Id = lastId
	return nil
}

func (r *Repository[T]) ReadMany(limit, offset int) ([]entity.Book, error) {
	d := []entity.Book{}
	err := r.db.Select(&d, "SELECT * FROM Books LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (r *Repository[T]) ReadOne(id string) (entity.Book, error) {
	b := entity.Book{}
	err := r.db.Get(&b, "SELECT * FROM Books WHERE id = $1 LIMIT 1", id)
	if err != nil {
		return b, err
	}
	return b, nil
}

func (r *Repository[T]) Update(id string, e *entity.Book) (int64, error) {
	q := "UPDATE Books SET title = $2, author = $3, publication_year = $4, updated_at = $5 WHERE id = $1"
	res, err := r.db.Exec(q, id, e.Title, e.Author, e.PublicationYear, e.UpdatedAt)
	if err != nil {
		return -1, err
	}
	return sqlxpkg.GetRowsAffected(res), nil
}

func (r *Repository[T]) Delete(id string) (int64, error) {
	q := "DELETE FROM Books WHERE id = $1"
	res, err := r.db.Exec(q, id)
	if err != nil {
		return -1, err
	}
	return sqlxpkg.GetRowsAffected(res), nil
}
