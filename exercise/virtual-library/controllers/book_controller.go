package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"virtual-library/models"
	"virtual-library/services"
)

type BookController struct {
	bookService services.BookService
}

func NewBookController(bookService services.BookService) *BookController {
	return &BookController{bookService: bookService}
}

// CreateBook creates a new book
// @Summary Create a new book
// @Description Create a new book with the provided details
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.Book true "Book object"
// @Success 201 {object} models.Book
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /books [post]
func (bc *BookController) CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	createdBook, err := bc.bookService.CreateBook(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, createdBook)
}

// ListBooks retrieves all books
// @Summary Get all books
// @Description Retrieve a list of all books
// @Tags books
// @Produce json
// @Success 200 {array} models.Book
// @Failure 500 {object} models.ErrorResponse
// @Router /books [get]
func (bc *BookController) ListBooks(c *gin.Context) {
	books, err := bc.bookService.ListBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to retrieve books"})
		return
	}

	c.JSON(http.StatusOK, books)
}

// GetBook retrieves a book by ID
// @Summary Get a book by ID
// @Description Retrieve a book by its ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /books/{id} [get]
func (bc *BookController) GetBook(c *gin.Context) {
	bookID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid book ID"})
		return
	}

	book, err := bc.bookService.GetBook(uint(bookID))
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// UpdateBook updates a book
// @Summary Update a book
// @Description Update a book with the provided details
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body models.Book true "Book object"
// @Success 200 {object} models.Book
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /books/{id} [put]
func (bc *BookController) UpdateBook(c *gin.Context) {
	bookID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid book ID"})
		return
	}

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	book.ID = uint(bookID)
	updatedBook, err := bc.bookService.UpdateBook(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	c.JSON(http.StatusOK, updatedBook)
}

// DeleteBook deletes a book
// @Summary Delete a book
// @Description Delete a book by its ID
// @Tags books
// @Param id path int true "Book ID"
// @Success 204
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /books/{id} [delete]
func (bc *BookController) DeleteBook(c *gin.Context) {
	bookID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid book ID"})
		return
	}

	err = bc.bookService.DeleteBook(uint(bookID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to delete book"})
		return
	}

	c.Status(http.StatusNoContent)
}
