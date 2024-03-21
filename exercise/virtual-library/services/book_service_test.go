package services_test

import (
	"errors"
	"virtual-library/models"
	"virtual-library/repositories/mocks"
	"virtual-library/services"

	"github.com/golang/mock/gomock"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("BookService", func() {
	var (
		ctrl        *gomock.Controller
		mockRepo    *mocks.MockBookRepositoryInterface
		bookService services.BookService
	)

	ginkgo.BeforeEach(func() {
		ctrl = gomock.NewController(ginkgo.GinkgoT())
		mockRepo = mocks.NewMockBookRepositoryInterface(ctrl)
		bookService = services.NewBookService(mockRepo)
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

			mockRepo.EXPECT().Create(book).Return(nil)

			result, err := bookService.CreateBook(book)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.Equal(book))
		})

		ginkgo.It("should return an error if repository fails", func() {
			book := &models.Book{
				Title:  "New Book",
				Author: "New Author",
			}

			mockRepo.EXPECT().Create(book).Return(errors.New("repository error"))

			result, err := bookService.CreateBook(book)
			gomega.Expect(err).To(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeNil())
		})
	})

	ginkgo.Describe("ListBooks", func() {
		ginkgo.It("should return a list of books", func() {
			books := []models.Book{
				{ID: 1, Title: "Book 1", Author: "Author 1"},
				{ID: 2, Title: "Book 2", Author: "Author 2"},
			}

			mockRepo.EXPECT().FindAll().Return(books, nil)

			result, err := bookService.ListBooks()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.Equal(books))
		})
	})

	ginkgo.Describe("GetBook", func() {
		ginkgo.It("should return a book by ID", func() {
			bookID := uint(1)
			book := &models.Book{ID: bookID, Title: "Book 1", Author: "Author 1"}

			mockRepo.EXPECT().FindByID(bookID).Return(book, nil)

			result, err := bookService.GetBook(bookID)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.Equal(book))
		})
	})

	ginkgo.Describe("UpdateBook", func() {
		ginkgo.It("should update a book", func() {
			bookID := uint(1)
			book := &models.Book{ID: bookID, Title: "Updated Book", Author: "Updated Author"}

			mockRepo.EXPECT().FindByID(bookID).Return(book, nil)
			mockRepo.EXPECT().Update(book).Return(nil)

			result, err := bookService.UpdateBook(book)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.Equal(book))
		})
	})

	ginkgo.Describe("DeleteBook", func() {
		ginkgo.It("should delete a book", func() {
			bookID := uint(1)

			mockRepo.EXPECT().Delete(bookID).Return(nil)

			err := bookService.DeleteBook(bookID)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
		})
	})
})
