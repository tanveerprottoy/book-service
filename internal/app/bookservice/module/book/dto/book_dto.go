package dto

type CreateUpdateBookDto struct {
	Title           string `json:"title" validate:"required"`
	Author          string `json:"author" validate:"required"`
	PublicationYear int32  `json:"publicationYear" validate:"required"`
}
