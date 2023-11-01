package server

import (
	"context"
	"log"
	"net/http"

	"connectrpc.com/connect"
	"github.com/gocql/gocql"
	v1 "github.com/slayer321/bookstore/api/gen/book/v1"
	"github.com/slayer321/bookstore/api/gen/book/v1/bookv1connect"
)

const address = "localhost:8080"

var db *gocql.Session

type bookStoreServiceServer struct {
	bookv1connect.UnimplementedInventoryServiceHandler
}

func CreateServer(d *gocql.Session) {

	mux := http.NewServeMux()
	path, handler := bookv1connect.NewInventoryServiceHandler(&bookStoreServiceServer{})
	mux.Handle(path, handler)
	db = d
	log.Printf("Listening on %s", address)
	http.ListenAndServe(address, mux)
}

func (b *bookStoreServiceServer) AddBooks(ctx context.Context, req *connect.Request[v1.Book]) (*connect.Response[v1.AddStatus], error) {
	book := req.Msg
	if db.Closed() {
		log.Println("DB session not open")
		return connect.NewResponse[v1.AddStatus](&v1.AddStatus{
			Code:    404,
			Message: "DB session not open",
		}), nil
	}
	if err := db.Query("INSERT INTO books (id, title, author, pages, publisher) VALUES (?, ?, ?, ?, ?)", gocql.TimeUUID(), book.Title, book.Author, book.Pages, book.Publisher).Exec(); err != nil {
		log.Println("Error while inserting book into database", err)
		return connect.NewResponse[v1.AddStatus](&v1.AddStatus{
			Code:    401,
			Message: "Unable to add book to database",
		}), err
	}

	log.Printf("Book added successfully: %v", book.Title)
	return connect.NewResponse[v1.AddStatus](&v1.AddStatus{
		Code:    200,
		Message: "Book added successfully",
	}), nil
}

func (b *bookStoreServiceServer) GetBooks(ctx context.Context, req *connect.Request[v1.BookTitle]) (*connect.Response[v1.Book], error) {
	bookTitle := req.Msg.Title

	var id string
	var title, author, publisher string
	var pages int32

	if err := db.Query("SELECT id, title , author, pages, publisher FROM books WHERE title = ?", bookTitle).Scan(&id, &title, &author, &pages, &publisher); err != nil {
		log.Fatal(err)
	}
	book := &v1.Book{
		Id:        id,
		Title:     title,
		Author:    author,
		Pages:     pages,
		Publisher: publisher,
	}

	log.Printf("Book retrieved successfully: %+v", book)
	return connect.NewResponse[v1.Book](book), nil
}

func (b *bookStoreServiceServer) DeleteBooks(ctx context.Context, req *connect.Request[v1.BookTitle]) (*connect.Response[v1.AddStatus], error) {
	bookTitle := req.Msg.Title

	if err := db.Query("DELETE FROM bookstore.books WHERE title = ?", bookTitle).Exec(); err != nil {
		log.Printf("Error deleting book: %v", err)
		return connect.NewResponse[v1.AddStatus](&v1.AddStatus{
			Code:    401,
			Message: "Unable to delete book from database",
		}), err
	}
	log.Printf("Book deleted successfully: %v", bookTitle)
	return connect.NewResponse[v1.AddStatus](&v1.AddStatus{
		Code:    200,
		Message: "Book deleted successfully",
	}), nil
}

func (b *bookStoreServiceServer) UpdateBooks(ctx context.Context, req *connect.Request[v1.Book]) (*connect.Response[v1.AddStatus], error) {
	book := req.Msg

	if err := db.Query("UPDATE bookstore.books SET author = ?, pages = ?, publisher = ? WHERE title = ?", book.Author, book.Pages, book.Publisher, book.Title).Exec(); err != nil {
		log.Printf("Error updating book: %v", err)
		return connect.NewResponse[v1.AddStatus](&v1.AddStatus{
			Code:    401,
			Message: "Unable to update book in database",
		}), err
	}
	log.Printf("Book updated successfully: %v", book.Title)
	return connect.NewResponse[v1.AddStatus](&v1.AddStatus{
		Code:    200,
		Message: "Book updated successfully",
	}), nil

}
