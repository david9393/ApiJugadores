package bd

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "players"
	
)

//var PostgresCN = ConectarBD()
var (
	once sync.Once
	db   *sql.DB
)

func ConectarBD() {
	once.Do(func() {
		var err error
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=		%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
		db, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Fatalf("error al conectarse")
		}
		err = db.Ping()
		if err != nil {
			log.Fatalf("no ping")
		}
		log.Printf("Conexion exitosaa")

	})
}

//pool retorna una unica instancia de db
func Pool() *sql.DB {
	return db
}
