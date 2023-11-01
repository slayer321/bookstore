# 📚 Bookstore

This is simple project which uses [gRPC](https://grpc.io/) along with [buf.build](https://buf.build/) to create a bookstore backend where you can store all the books that you have read. It uses [cassandra](https://cassandra.apache.org/_/index.html) as it Database.


## 🛠️ Pre-requisite
- docker
- [buf cli](https://buf.build/docs/installation)
- cassandra db
  
If you already have cassandra db running on your system run below command to create the keyspace and it's schema.
```
create keyspace bookstore with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
create table bookstore.books(id UUID, title text, author text, pages int, publisher text, PRIMARY KEY(title));
```

We also have a make command which will setup all this for you. You just need install docker and buf cli. And run below command

_Note: Clone the repo_
```
make db
```

## 🚋 Usage
_Note: Run it from root of the project_

Open the first terminal and run below commands.

```
$ go run cmd/main.go
```

Open second terminal and run below commands.

- Add Book

```
$ buf curl \
  --schema api \
  --data '{"title": "The Alchemist","author": "Paulo Coelho", "pages": 197,"publisher": "HarperCollins"}' \
  http://localhost:8080/book.v1.InventoryService/AddBooks
```

```
$ buf curl \
  --schema api \
  --data '{"title": "Kafka on shore","author": "Haruki Murakami", "pages": 505,"publisher": "HarperCollins"}' \
  http://localhost:8080/book.v1.InventoryService/AddBooks
```

- Get Book
```
$ buf curl \
  --schema api \
  --data '{"title": "The Alchemist"}' \
  http://localhost:8080/book.v1.InventoryService/GetBooks
```

- Delete Book
```
$ buf curl \
  --schema api \
  --data '{"title": "The Alchemist"}' \
  http://localhost:8080/book.v1.InventoryService/DeleteBooks
```
- Update Book

```
$ buf curl \
  --schema api \
  --data '{"title": "Kafka on shore","author": "Haruki Murakami2", "pages": 5051,"publisher": "HarperCollinss"}' \
  http://localhost:8080/book.v1.InventoryService/UpdateBooks
```


## 🙌 Credits

[edouard-lopez](https://github.com/edouard-lopez/progress-bar.sh) : Progess Bar

## 🍁 Contributing

If you want to add new features or find a bug raise an issue. 
