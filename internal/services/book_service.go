package services

import (
	"errors"

	"github.com/idoceb00/library-service/internal/models"
	"github.com/idoceb00/library-service/internal/repositories"
)

type BookService struct {
	repo *repositories.BookRepository
}

func NewBookService(repo *repositories.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetAll() ([]models.Book, error) {
	return s.repo.FindAll()
}

func (s *BookService) GetById(id uint) (*models.Book, error) {
	book, err := s.repo.FindById(id)
	if err != nil {
		return nil, errors.New("Book not Found")
	}

	return book, nil
}

func (s *BookService) Create(book *models.Book) error {
	return s.repo.Create(book)
}

func (s *BookService) Update(id uint, data *models.Book) (*models.Book, error) {
	book, err := s.repo.FindById(id)
	if err != nil {
		return nil, errors.New("Book not found")
	}

	book.Title = data.Title
	book.Author = data.Author

	err = s.repo.Update(book)
	return book, err
}

func (s *BookService) Delete(id uint) error {
	return s.repo.Delete(id)
}
