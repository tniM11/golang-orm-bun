package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"

	"github.com/example/todo/internal/domain"
	httpapi "github.com/example/todo/internal/http"
	"github.com/example/todo/internal/infra/sqlite"
	"github.com/example/todo/internal/service"
)

func main() {
	sqldb, err := sql.Open(sqliteshim.ShimName, "file:todos.db?cache=shared&_foreign_keys=on")
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer sqldb.Close()

	db := bun.NewDB(sqldb, sqlitedialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	if _, err := db.NewCreateTable().Model((*domain.Todo)(nil)).IfNotExists().Exec(nil); err != nil {
		log.Fatalf("create table: %v", err)
	}

	repo := sqlite.NewTodoRepository(db)
	svc := service.NewTodoService(repo)
	handler := httpapi.NewHandler(svc)

	mux := http.NewServeMux()
	handler.Register(mux)

	log.Println("listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
