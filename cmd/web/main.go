package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"os"

	"github.com/gin-contrib/sessions/cookie"

	_ "github.com/go-sql-driver/mysql"

	"masrurimz/snippetbox/pkg/models/mysql"
)

// Define an application struct to hold the application-wide dependencies for the
// web application. For now we'll only include fields for the two custom loggers, but
// we'll add more to it as the build progresses.
type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	store         *cookie.Store
	snippets      *mysql.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	// Define a new command-line flag with the name 'addr', a default value of ":4000"
	addr := flag.String("addr", ":4000", "HTTP network address")

	// Define a new command-line flag for the MySQL DSN string.
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")

	// Define new command-line flag with the name 'secret' as secretKey for session

	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")
	// Importantly, we use the flag.Parse() function to parse the command-line flag.
	flag.Parse()

	// Use log.New() to create a logger for writing information messages. This takes
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create error mesage logger using stderr and enable file name and line flag (Lshortfile)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initiate DB connecction and config from func openDB using dsn param
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	// Initialize new template cache
	templateCache, err := newTemplateChache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	// Initialize session
	store := cookie.NewStore([]byte(*secret))

	// Initialize a new instance of application containing the dependencies.
	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		store:         &store,
		snippets:      &mysql.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	srv := app.routes()

	infoLog.Printf("Starting server on :4000")

	err = srv.Run(*addr)
	errorLog.Fatal(err)
}

// The openDB() function wraps sql.Open() and returns a sql.DB connection pool
// for a given DSN.
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
