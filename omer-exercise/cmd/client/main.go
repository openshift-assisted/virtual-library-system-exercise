package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	sw "virtual-library/client"
)

func createBook(apiClient *sw.APIClient) {
	newBook := sw.ModelsBook{
		Title:       "New Book",
		Author:      "Author Name",
		PublishedAt: "2023-06-08",
	}
	book, _, err := apiClient.BooksApi.BooksPost(context.Background(), newBook)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Book {} created successfully", &book)
}

func listBooks(apiClient *sw.APIClient) {
	books, _, err := apiClient.BooksApi.BooksGet(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.Id, book.Title, book.Author)
	}
}

func getBookDetails(apiClient *sw.APIClient, bookID int32) {
	book, response, err := apiClient.BooksApi.BooksIdGet(context.Background(), bookID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("{}", response)
	fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.Id, book.Title, book.Author)
}

func updateBook(apiClient *sw.APIClient, bookID int32) {
	updatedBook := sw.ModelsBook{
		Title:  "Updated Book",
		Author: "Updated Author",
	}
	updated, _, err := apiClient.BooksApi.BooksIdPut(context.Background(), updatedBook, bookID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Book {} updated successfully", &updated)
}

func deleteBook(apiClient *sw.APIClient, bookID int32) {
	_, err := apiClient.BooksApi.BooksIdDelete(context.Background(), bookID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Book deleted successfully")
}

func main() {
	cfg := sw.NewConfiguration()
	cfg.BasePath = "http://localhost:8080/api/v1"
	apiClient := sw.NewAPIClient(cfg)

	if len(os.Args) < 2 {
		fmt.Println("Please provide a command.")
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
	case "create":
		createBook(apiClient)
	case "list":
		listBooks(apiClient)
	case "get":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a book ID.")
			os.Exit(1)
		}
		bookID, _ := strconv.ParseInt(os.Args[2], 10, 32)
		getBookDetails(apiClient, int32(bookID))
	case "update":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a book ID.")
			os.Exit(1)
		}
		bookID, _ := strconv.ParseInt(os.Args[2], 10, 32)
		updateBook(apiClient, int32(bookID))
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a book ID.")
			os.Exit(1)
		}
		bookID, _ := strconv.ParseInt(os.Args[2], 10, 32)
		deleteBook(apiClient, int32(bookID))
	default:
		fmt.Println("Invalid command.")
		os.Exit(1)
	}
}
