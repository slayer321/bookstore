package server

import (
	"context"
	"fmt"
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

// func NewBookStoreServiceServer(db gocql.Session) *bookStoreServiceServer {
// 	return &bookStoreServiceServer{
// 		db: db,
// 	}
// }

func CreateServer(d *gocql.Session) {

	mux := http.NewServeMux()
	path, handler := bookv1connect.NewInventoryServiceHandler(&bookStoreServiceServer{})
	mux.Handle(path, handler)
	db = d
	//server := NewBookStoreServiceServer(db)
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
	log.Println("Just before inserting book into database")
	if err := db.Query("INSERT INTO book (id, title, author, pages, publisher) VALUES (?, ?, ?, ?, ?)", gocql.TimeUUID(), book.Title, book.Author, book.Pages, book.Publisher).Exec(); err != nil {
		log.Println("Error while inserting book into database", err)
		return connect.NewResponse[v1.AddStatus](&v1.AddStatus{
			Code:    401,
			Message: "Unable to add book to database",
		}), err
	}

	return connect.NewResponse[v1.AddStatus](&v1.AddStatus{
		Code:    200,
		Message: "Book added successfully",
	}), nil
}

func (b *bookStoreServiceServer) GetBooks(ctx context.Context, req *connect.Request[v1.BookTitle]) (*connect.Response[v1.Book], error) {
	//books := getSampleBooks()
	bookTitle := req.Msg.Title

	var id string
	var title, author, publisher string
	var pages int32

	if err := db.Query("SELECT id, title , author, pages, publisher FROM book WHERE title = ?", bookTitle).Scan(&id, &title, &author, &pages, &publisher); err != nil {
		log.Fatal(err)
	}
	book := &v1.Book{
		Id:        id,
		Title:     title,
		Author:    author,
		Pages:     pages,
		Publisher: publisher,
	}

	fmt.Println("Book:", book)
	return connect.NewResponse[v1.Book](book), nil
}

// func getSampleBooks() []*v1.Book {
// 	return []*v1.Book{
// 		{
// 			Id:        1,
// 			Title:     "The Alchemist",
// 			Author:    "Paulo Coelho",
// 			Pages:     197,
// 			Publisher: "HarperCollins",
// 		},
// 		{
// 			Id:        2,
// 			Title:     "Kafka on shore",
// 			Author:    "Haruki Murakami",
// 			Pages:     505,
// 			Publisher: "HarperCollins",
// 		},
// 	}
// }
