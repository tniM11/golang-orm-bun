# Todo list with Bun

This is a simple todo list project demonstrating a hexagonal architecture in Go.
It uses the [Bun](https://bun.uptrace.dev/) ORM with SQLite and the
[Gin](https://gin-gonic.com/) web framework.

Run the server:

```bash
go run ./cmd/todo
```

The HTTP API exposes CRUD endpoints under `/todos`.
