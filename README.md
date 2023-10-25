# [WIP] bookstore




```
buf curl \
  --schema api \
  --data '{"title": "The Alchemist","author": "Paulo Coelho", "pages": 197,"publisher": "HarperCollins"}' \
  http://localhost:8080/book.v1.InventoryService/AddBooks
```

```
buf curl \
  --schema api \
  --data '{"title": "Kafka on shore","author": "Haruki Murakami", "pages": 505,"publisher": "HarperCollins"}' \
  http://localhost:8080/book.v1.InventoryService/AddBooks
```

```
buf curl \
  --schema api \
  --data '{"title": "The Alchemist"}' \
  http://localhost:8080/book.v1.InventoryService/GetBooks
```

```
create keyspace bookstore with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
create table bookstore.book(id UUID, title text, author text, pages int, publisher text, PRIMARY KEY(id));
```

```
INSERT INTO bookstore.book (id, title, author, pages, publisher) VALUES (b74599f5-7332-11ee-9be9-88a4c226c1e1, 'The Alchemist', 'Paulo Coelho', 197, 'HarperCollins');
```

```
create index on bookstore.book(title);
```