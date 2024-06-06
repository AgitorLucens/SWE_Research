package db

import (
	"backend/utils/types"
	utils "backend/utils/util"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	_ "github.com/lib/pq"
)

type ApiDatabase struct {
	connStr string
	DB      *sql.DB
}

func NewApiDatabase(connStr string) *ApiDatabase {
	return &ApiDatabase{
		connStr: connStr,
	}
}

func (apidatabase *ApiDatabase) Run() (*sql.DB, error) {
	db, err := sql.Open("postgres", apidatabase.connStr)
	utils.CheckError(err)
	apidatabase.DB = db
	createRecordTable(db)
	createMillionRecord(db)
	createIndexTable(db)
	createFuncFilter(db)
	log.Println("Conexión exitosa a la base de datos PostgreSQL")
	return db, err
}

func insertRecordsInBatch(db *sql.DB, totalRecords, batchSize int) {
	for i := 0; i < totalRecords; i += batchSize {
		remaining := totalRecords - i
		batchSize := calculateBatchSize(batchSize, remaining)
		valuesStr := buildValuesString(batchSize)
		query := buildQuery(valuesStr)
		executeQuery(db, query)
	}
}

func calculateBatchSize(batchSize, remaining int) int {
	if remaining < batchSize {
		return remaining
	}
	return batchSize
}

func buildValuesString(batchSize int) string {
	var values strings.Builder
	values.WriteString("")
	for j := 0; j < batchSize; j++ {
		name, price, description, img, quant, topic, created := utils.RandomData()
		values.WriteString(fmt.Sprintf("('%s', '%s', %.2f, '%s', %d, '%s', '%s'),", name, img, price, description, quant, topic, created))
	}
	valuesStr := values.String()
	return valuesStr[:len(valuesStr)-1]
}


func executeQuery(db *sql.DB, query string) {
	_, err := db.Exec(query)
	utils.CheckError(err)
}

func createRecordTable(db *sql.DB) {
	data, _ := os.ReadFile("internal/server/db/changelog/001-create-table.sql")
	query := string(data)
	_, err := db.Exec(query)
	utils.CheckError(err)
}
func createIndexTable(db *sql.DB) {
	data, _ := os.ReadFile("internal/server/db/changelog/006-create-table.sql")
	query := string(data)
	_, err := db.Exec(query)
	utils.CheckError(err)
}

func ExtractRecord(db *sql.DB, values []string) ([]types.Record, error) {
	rows, err := getRecords(db, values)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	var records []types.Record
	for rows.Next() {
		var record types.Record
		err := rows.Scan(&record.Name, &record.Descr, &record.Image, &record.Price, &record.Topic, &record.Created, &record.Published, &record.Quant)
		utils.CheckError(err)
		records = append(records, record)
	}
	err2 := rows.Err()
	utils.CheckError(err2)
	return records, err2
}

func getRecords(db *sql.DB, params []string) (*sql.Rows, error) {

	query := GetQuery(params)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return rows, err
}

func buildQuery(valuesStr string) string {
	data, err := os.ReadFile("internal/server/db/changelog/002-insert-record.sql")
    utils.CheckError(err)
	return fmt.Sprintf("%s %s", string(data), valuesStr)
}

func createFuncFilter(db *sql.DB) {
	data, err := os.ReadFile("internal/server/db/changelog/004-filter-function.sql")
	query := string(data)
	_, err = db.Exec(query)
	utils.CheckError(err)
}
//Genera query BD
func GetQuery(params []string) string {
	paths := []string{"name", "descr", "quant", "price", "topic", "created", "published", "page", "limit"}
	list := make([]interface{}, 9)
	for i := 0; i < len(paths); i++ {
		if params[i] != "" {
			if paths[i]== "quant" || paths[i]== "price" || paths[i]== "page" || paths[i]== "limit" {
				list[i] = params[i]
			} else {
				list[i] = "'" + params[i] + "'"
			}
		} else {
			list[i] = "NULL"
		}
	}
	return fmt.Sprintf("SELECT * FROM filter_records (%s, %s, %s, %s, %s, %s, %s, %s, %s)", list[0], list[1], list[2], list[3], list[4], list[5], list[6], list[7], list[8])
}

func createMillionRecord(db *sql.DB) {
	var numWorkers = 10
	var recordsPerWorker = 100000
	var waitGroup sync.WaitGroup
	waitGroup.Add(numWorkers)
	batchSize := 10000
	for i := 0; i < numWorkers; i++ {
		go routineFunc(&waitGroup, db , recordsPerWorker, batchSize)
	}
	waitGroup.Wait()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func routineFunc(w *sync.WaitGroup, db *sql.DB, recordsPerWorker int, batchSize int){
	defer w.Done()
	insertRecordsInBatch(db, recordsPerWorker, batchSize)
}
