package entity

type Book struct {
	Id              string `db:"id" json:"id"`
	Title           string `db:"title" json:"title"`
	PublicationYear int32  `db:"publication_year" json:"publicationYear"`
	CreatedAt       int64  `db:"created_at" json:"createdAt"`
	UpdatedAt       int64  `db:"updated_at" json:"updatedAt"`
}
