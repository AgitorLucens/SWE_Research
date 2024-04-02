package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Record struct {
	Name  string `json:"name"`
	Price string `json:"price"`
	Descr string `json:"desc"`
}
type APIDB struct {
	connStr string
	DB      *sql.DB
}

func newAPIDB(connStr string) *APIDB {
	return &APIDB{
		connStr: connStr,
	}
}

func (d *APIDB) initDB() error {

	db, err := sql.Open("postgres", d.connStr)
	if err != nil {
		log.Fatal(err)
	}
	d.DB = db
	createRecordTable(db)
	//insertRecordTable(db)
	//defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Conexión exitosa a la base de datos PostgreSQL")
	return err
}

//it works
func extractRecord(db *sql.DB) (*Record, error) {
	query := "select name, price, descr from  record"

	var data Record

	err := db.QueryRow(query).Scan(&data.Name, &data.Price, &data.Descr)

	if err != nil {
		log.Fatal(err)
		return nil,err
	}
	
	return &data, err
}

// it works
func createRecordTable(db *sql.DB) {

	query := "CREATE TABLE IF NOT EXISTS record (id SERIAL PRIMARY KEY,name VARCHAR(100) NOT NULL,img BYTEA, price NUMERIC(6,2) NOT NULL,descr VARCHAR(200), created timestamp DEFAULT NOW())"

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

}

func insertRecordTable(db *sql.DB) int {
	query := `INSERT INTO record (name, price, descr, created)
VALUES 
    ('Producto 1',222,'Productadadsd', NOW() - INTERVAL '5 days')`
	_, err := db.Exec(query)
	var pk int
	if err != nil {
		log.Fatal(err)
	}
	return pk
}
