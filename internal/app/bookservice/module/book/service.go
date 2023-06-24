package book

import (
	"errors"
	"net/http"
	"time"

	"github.com/tanveerprottoy/book-service/internal/app/bookservice/module/book/dto"
	"github.com/tanveerprottoy/book-service/internal/app/bookservice/module/book/entity"
	"github.com/tanveerprottoy/book-service/internal/pkg/constant"
	"github.com/tanveerprottoy/book-service/pkg/data/sqlxpkg"
	"github.com/tanveerprottoy/book-service/pkg/response"
)

type Service struct {
	repository sqlxpkg.Repository[entity.Book]
}

func NewService(r sqlxpkg.Repository[entity.Book]) *Service {
	s := new(Service)
	s.repository = r
	return s
}

func (s *Service) handleError(err error, w http.ResponseWriter) {
	if err.Error() == "sql: no rows in result set" {
		response.RespondError(http.StatusNotFound, errors.New(constant.NotFound), w)
		return
	}
	response.RespondError(http.StatusBadRequest, err, w)
}

func (s *Service) readOneInternal(id string, w http.ResponseWriter) (entity.Book, error) {
	return s.repository.ReadOne(id)
}

func (s *Service) Create(d *dto.CreateUpdateBookDto, w http.ResponseWriter, r *http.Request) {
	// convert dto to entity
	b := entity.Book{}
	b.Title = d.Title
	b.PublicationYear = d.PublicationYear
	n := time.Now().UnixMilli()
	b.CreatedAt = n
	b.UpdatedAt = n
	err := s.repository.Create(&b)
	if err != nil {
		s.handleError(err, w)
		return
	}
	response.Respond(http.StatusCreated, response.BuildData(d), w)
}

func (s *Service) ReadMany(limit, page int, w http.ResponseWriter, r *http.Request) {
	offset := limit * (page - 1)
	d, err := s.repository.ReadMany(limit, offset)
	if err != nil {
		s.handleError(err, w)
		return
	}
	m := make(map[string]any)
	m["items"] = d
	m["limit"] = limit
	m["page"] = page
	response.Respond(http.StatusOK, response.BuildData(m), w)
}

func (s *Service) ReadOne(id string, w http.ResponseWriter, r *http.Request) {
	b, err := s.readOneInternal(id, w)
	if err != nil {
		s.handleError(err, w)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(b), w)
}

func (s *Service) Update(id string, d *dto.CreateUpdateBookDto, w http.ResponseWriter, r *http.Request) {
	b, err := s.readOneInternal(id, w)
	if err != nil {
		s.handleError(err, w)
		return
	}
	b.Title = d.Title
	b.PublicationYear = d.PublicationYear
	b.UpdatedAt = time.Now().UnixMilli()
	rows, err := s.repository.Update(id, &b)
	if err != nil {
		s.handleError(err, w)
		return
	}
	if rows > 0 {
		response.Respond(http.StatusOK, response.BuildData(b), w)
		return
	}
	response.RespondError(http.StatusBadRequest, errors.New("operation was not successful"), w)
}

func (s *Service) Delete(id string, w http.ResponseWriter, r *http.Request) {
	b, err := s.readOneInternal(id, w)
	if err != nil {
		s.handleError(err, w)
		return
	}
	rows, err := s.repository.Delete(id)
	if err != nil {
		s.handleError(err, w)
		return
	}
	if rows > 0 {
		response.Respond(http.StatusOK, response.BuildData(b), w)
		return
	}
	response.RespondError(http.StatusBadRequest, errors.New("operation was not successful"), w)
}
