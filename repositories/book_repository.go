package repositories

import (
	"gorm.io/gorm"
	"virtual-library/models"
)

//go:generate mockgen -source=book_repository.go -destination=mocks/book_repository_mock.go -package=mocks

type BookRepository interface {
	Create(book *models.Book) error
	FindAll() ([]models.Book, error)
	FindByID(id uint) (*models.Book, error)
	Update(book *models.Book) error
	Delete(id uint) error
}

type bookRepository struct {
	BookRepository
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *bookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) Create(book *models.Book) error {
	return r.db.Create(book).Error
}

func (r *bookRepository) FindAll() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *bookRepository) FindByID(id uint) (*models.Book, error) {
	var book models.Book
	err := r.db.First(&book, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &book, nil
}

func (r *bookRepository) Update(book *models.Book) error {
	return r.db.Save(book).Error
}

func (r *bookRepository) Delete(id uint) error {
	return r.db.Delete(&models.Book{}, id).Error
}
