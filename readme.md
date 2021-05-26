# import

```shell
go get -u github.com/gorilla/mux
go get -u github.com/jmoiron/sqlx
go get -u github.com/lib/pq
```

# db
```shell
docker run --name postapidb --env POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres
```