package services

import (
	"virtual-library/models"
	"virtual-library/repositories"
)

//go:generate mockgen -source=book_service.go -destination=mocks/book_service_mock.go -package=mocks

type BookService interface {
	CreateBook(book *models.Book) (*models.Book, error)
	ListBooks() ([]models.Book, error)
	GetBook(id uint) (*models.Book, error)
	UpdateBook(book *models.Book) (*models.Book, error)
	DeleteBook(id uint) error
}

type bookService struct {
	BookService
	bookRepo repositories.BookRepository
}

func NewBookService(bookRepo repositories.BookRepository) *bookService {
	return &bookService{bookRepo: bookRepo}
}

func (s *bookService) CreateBook(book *models.Book) (*models.Book, error) {
	err := s.bookRepo.Create(book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (s *bookService) ListBooks() ([]models.Book, error) {
	return s.bookRepo.FindAll()
}

func (s *bookService) GetBook(id uint) (*models.Book, error) {
	return s.bookRepo.FindByID(id)
}

func (s *bookService) UpdateBook(book *models.Book) (*models.Book, error) {
	existingBook, err := s.bookRepo.FindByID(book.ID)
	if err != nil {
		return nil, err
	}
	if existingBook == nil {
		return nil, nil
	}

	existingBook.Title = book.Title
	existingBook.Author = book.Author
	existingBook.PublishedAt = book.PublishedAt

	err = s.bookRepo.Update(existingBook)
	if err != nil {
		return nil, err
	}

	return existingBook, nil
}

func (s *bookService) DeleteBook(id uint) error {
	return s.bookRepo.Delete(id)
}
