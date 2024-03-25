package repositories_test

import (
	"database/sql"
	"time"
	"virtual-library/models"
	"virtual-library/repositories"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _ = Describe("BookRepository", func() {
	var (
		db       *gorm.DB
		mock     sqlmock.Sqlmock
		bookRepo repositories.BookRepository
	)

	BeforeEach(func() {
		var (
			dbMock *sql.DB
			err    error
		)

		dbMock, mock, err = sqlmock.New()
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		db, err = gorm.Open(postgres.New(postgres.Config{
			Conn: dbMock,
		}), &gorm.Config{})
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		bookRepo = repositories.NewBookRepository(db)
	})

	AfterEach(func() {
		db.Migrator().DropTable(&models.Book{})
		sqlDB, _ := db.DB()
		sqlDB.Close()
	})

	Describe("FindAll", func() {
		It("should return all books", func() {
			book1 := models.Book{
				ID:          1,
				Title:       "Book 1",
				Author:      "Author 1",
				PublishedAt: time.Now(),
			}
			book2 := models.Book{
				ID:          2,
				Title:       "Book 2",
				Author:      "Author 2",
				PublishedAt: time.Now(),
			}
			book3 := models.Book{
				ID:          3,
				Title:       "Book 3",
				Author:      "Author 3",
				PublishedAt: time.Now(),
			}

			rows := sqlmock.NewRows([]string{"id", "title", "author", "published_at", "created_at", "updated_at", "deleted_at"}).
				AddRow(1, book1.Title, book1.Author, book1.PublishedAt, book1.CreatedAt, book1.UpdatedAt, book1.DeletedAt).
				AddRow(2, book2.Title, book2.Author, book2.PublishedAt, book2.CreatedAt, book2.UpdatedAt, book2.DeletedAt).
				AddRow(3, book3.Title, book3.Author, book3.PublishedAt, book3.CreatedAt, book3.UpdatedAt, book3.DeletedAt)

			mock.ExpectQuery(`SELECT \* FROM "books" WHERE "books"."deleted_at" IS NULL`).WillReturnRows(rows)

			books, err := bookRepo.FindAll()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(books).To(gomega.HaveLen(3))
			gomega.Expect(books[0]).To(gomega.Equal(book1))
			gomega.Expect(books[1]).To(gomega.Equal(book2))
			gomega.Expect(books[2]).To(gomega.Equal(book3))

			gomega.Expect(mock.ExpectationsWereMet()).NotTo(gomega.HaveOccurred())
		})
	})

	Describe("Create", func() {
		It("should create a new book", func() {
			book := models.Book{
				Title:       "New Book",
				Author:      "New Author",
				PublishedAt: time.Now(),
			}

			mock.ExpectBegin()
			mock.ExpectQuery(`INSERT INTO "books" (.+) RETURNING "id"`).
				WithArgs(book.Title, book.Author, book.PublishedAt, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			mock.ExpectCommit()

			err := bookRepo.Create(&book)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(book.ID).To(gomega.Equal(uint(1)))

			gomega.Expect(mock.ExpectationsWereMet()).NotTo(gomega.HaveOccurred())
		})
	})

	Describe("FindByID", func() {
		It("should find a book by ID", func() {
			book := models.Book{
				ID:          1,
				Title:       "Book 1",
				Author:      "Author 1",
				PublishedAt: time.Now(),
			}

			rows := sqlmock.NewRows([]string{"id", "title", "author", "published_at", "created_at", "updated_at", "deleted_at"}).
				AddRow(book.ID, book.Title, book.Author, book.PublishedAt, book.CreatedAt, book.UpdatedAt, book.DeletedAt)

			mock.ExpectQuery(`SELECT \* FROM "books" WHERE "books"."id" = \$1 AND "books"."deleted_at" IS NULL ORDER BY "books"."id" LIMIT \$2`).
				WithArgs(book.ID, 1).
				WillReturnRows(rows)

			foundBook, err := bookRepo.FindByID(book.ID)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(foundBook.Title).To(gomega.Equal(book.Title))

			gomega.Expect(mock.ExpectationsWereMet()).NotTo(gomega.HaveOccurred())
		})
	})

	Describe("Update", func() {
		It("should update a book", func() {
			book := models.Book{
				ID:          1,
				Title:       "Book 1",
				Author:      "Author 1",
				PublishedAt: time.Now(),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}

			mock.ExpectBegin()
			mock.ExpectExec(`UPDATE "books" SET (.+) WHERE (.+)`).
				WithArgs("Updated Book", "Updated Author", book.PublishedAt, sqlmock.AnyArg(), sqlmock.AnyArg(), nil, book.ID).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			book.Title = "Updated Book"
			book.Author = "Updated Author"

			err := bookRepo.Update(&book)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			gomega.Expect(mock.ExpectationsWereMet()).NotTo(gomega.HaveOccurred())
		})
	})

	Describe("Delete", func() {
		It("should delete a book", func() {
			bookID := uint(1)

			mock.ExpectBegin()
			mock.ExpectExec(`UPDATE "books" SET "deleted_at"=\$1 WHERE "books"."id" = \$2 AND "books"."deleted_at" IS NULL`).
				WithArgs(sqlmock.AnyArg(), bookID).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			err := bookRepo.Delete(bookID)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			gomega.Expect(mock.ExpectationsWereMet()).NotTo(gomega.HaveOccurred())
		})
	})
})
