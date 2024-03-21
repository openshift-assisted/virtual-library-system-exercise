package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"virtual-library/controllers"
	"virtual-library/models"
	"virtual-library/services/mocks"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("BookController", func() {
	var (
		ctrl         *gomock.Controller
		mockService  *mocks.MockBookService
		bookCtrl     *controllers.BookController
		router       *gin.Engine
		response     *httptest.ResponseRecorder
		request      *http.Request
		requestBody  []byte
		responseBody []byte
	)

	ginkgo.BeforeEach(func() {
		ctrl = gomock.NewController(ginkgo.GinkgoT())
		mockService = mocks.NewMockBookService(ctrl)
		bookCtrl = controllers.NewBookController(mockService)
		router = gin.Default()
		response = httptest.NewRecorder()
	})

	ginkgo.AfterEach(func() {
		ctrl.Finish()
	})

	ginkgo.Describe("CreateBook", func() {
		ginkgo.It("should create a new book", func() {
			book := &models.Book{
				Title:  "New Book",
				Author: "New Author",
			}
			requestBody, _ = json.Marshal(book)
			request, _ = http.NewRequest("POST", "/books", bytes.NewBuffer(requestBody))

			mockService.EXPECT().CreateBook(book).Return(book, nil)

			router.POST("/books", bookCtrl.CreateBook)
			router.ServeHTTP(response, request)

			gomega.Expect(response.Code).To(gomega.Equal(http.StatusCreated))
			responseBody, _ = json.Marshal(book)
			gomega.Expect(response.Body.String()).To(gomega.MatchJSON(string(responseBody)))
		})

		ginkgo.It("should return an error if service fails", func() {
			book := &models.Book{
				Title:  "New Book",
				Author: "New Author",
			}
			requestBody, _ = json.Marshal(book)
			request, _ = http.NewRequest("POST", "/books", bytes.NewBuffer(requestBody))

			mockService.EXPECT().CreateBook(book).Return(nil, errors.New("service error"))

			router.POST("/books", bookCtrl.CreateBook)
			router.ServeHTTP(response, request)

			gomega.Expect(response.Code).To(gomega.Equal(http.StatusInternalServerError))
			gomega.Expect(response.Body.String()).To(gomega.ContainSubstring("Failed to create book"))
		})
	})

	ginkgo.Describe("ListBooks", func() {
		ginkgo.It("should return a list of books", func() {
			books := []models.Book{
				{ID: 1, Title: "Book 1", Author: "Author 1"},
				{ID: 2, Title: "Book 2", Author: "Author 2"},
			}

			mockService.EXPECT().ListBooks().Return(books, nil)

			request, _ = http.NewRequest("GET", "/books", nil)
			router.GET("/books", bookCtrl.ListBooks)
			router.ServeHTTP(response, request)

			gomega.Expect(response.Code).To(gomega.Equal(http.StatusOK))
			responseBody, _ = json.Marshal(books)
			gomega.Expect(response.Body.String()).To(gomega.MatchJSON(string(responseBody)))
		})
	})

	ginkgo.Describe("GetBook", func() {
		ginkgo.It("should return a book by ID", func() {
			bookID := uint(1)
			book := &models.Book{ID: bookID, Title: "Book 1", Author: "Author 1"}

			mockService.EXPECT().GetBook(bookID).Return(book, nil)

			request, _ = http.NewRequest("GET", "/books/1", nil)
			router.GET("/books/:id", bookCtrl.GetBook)
			router.ServeHTTP(response, request)

			gomega.Expect(response.Code).To(gomega.Equal(http.StatusOK))
			responseBody, _ = json.Marshal(book)
			gomega.Expect(response.Body.String()).To(gomega.MatchJSON(string(responseBody)))
		})
	})

	ginkgo.Describe("UpdateBook", func() {
		ginkgo.It("should update a book", func() {
			bookID := uint(1)
			book := &models.Book{ID: bookID, Title: "Updated Book", Author: "Updated Author"}

			mockService.EXPECT().UpdateBook(book).Return(book, nil)

			requestBody, _ = json.Marshal(book)
			request, _ = http.NewRequest("PUT", "/books/1", bytes.NewBuffer(requestBody))

			router.PUT("/books/:id", bookCtrl.UpdateBook)
			router.ServeHTTP(response, request)

			gomega.Expect(response.Code).To(gomega.Equal(http.StatusOK))
			responseBody, _ = json.Marshal(book)
			gomega.Expect(response.Body.String()).To(gomega.MatchJSON(string(responseBody)))
		})
	})

	ginkgo.Describe("DeleteBook", func() {
		ginkgo.It("should delete a book", func() {
			bookID := uint(1)

			mockService.EXPECT().DeleteBook(bookID).Return(nil)

			request, _ = http.NewRequest("DELETE", "/books/1", nil)
			router.DELETE("/books/:id", bookCtrl.DeleteBook)
			router.ServeHTTP(response, request)

			gomega.Expect(response.Code).To(gomega.Equal(http.StatusNoContent))
		})
	})
})
