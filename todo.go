package main

import (
    "github.com/labstack/echo"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "go-echo-vue/handlers"
    // "net/http"
    // "github.com/labstack/engine/standard"

)

func main() {
    db := initDB("storage.db")
    migrate(db)

    // Create a new instance of Echo
    e := echo.New()


    e.File("/", "public/index.html")
    e.GET("/tasks", handlers.GetTasks(db))
    e.PUT("/tasks/:id", handlers.EditTasks(db))
    e.PUT("/tasks", handlers.PutTask(db))
    e.DELETE("/tasks/:id", handlers.DeleteTask(db))

    // Start as a web server
    // e.Run(standard.New(":8000"))
    e.Logger.Fatal(e.Start(":8000"))
}

func initDB(filepath string) *sql.DB {
    db, err := sql.Open("sqlite3", filepath)

    // Here we check for any db errors then exit
    if err != nil {
        panic(err)
    }

    // If we don't get any errors but somehow still don't get a db connection
    // we exit as well
    if db == nil {
        panic("db nil")
    }
    return db
}

func migrate(db *sql.DB) {
    sql := `
    CREATE TABLE IF NOT EXISTS tasks(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name VARCHAR NOT NULL
    );
    `

    _, err := db.Exec(sql)
    // Exit if something goes wrong with our SQL statement above
    if err != nil {
        panic(err)
    }
}