package main

import (
	"crawler/internal/books"
	"crawler/internal/database"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	//ต่อ Data base
	dbConfig := database.Config{
		User:       "smalldoc",
		Password:   "example",
		Host:       "localhost",
		Port:       5432,
		Name:       "books_collection_th",
		DisableTLS: true,
	}

	db, err := database.Open(dbConfig)
	if err != nil {
		log.Fatal("connecting database fail", err)
	}
	bookPostgres := books.BookDB{
		Connection: db,
	}
	// อ่านค่าจาก csv file
	file, err := os.Open("book_data_3.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(file)
	columeList, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Can not read: ", err)
	}
	var errorInsert []string
	for index, data := range columeList {
		if index == 1 {
			continue
		}
		bookNew := books.Books{
			ISBN:       data[0],
			Name:       data[1],
			Writer:     data[2],
			Translator: data[3],
			Publisher:  data[4],
			PrintYear:  data[5],
		}
		fmt.Printf("insertbook :) ISBN: %s\n", data[0])
		// insert ลง Data base
		if err := bookPostgres.CreateBook(bookNew); err != nil {
			fmt.Printf("insertbook :( ISBN: %s error: %v\n", data[0], err)
			errorInsert = append(errorInsert, data[0])
		}
	}
	fmt.Println("error :", len(errorInsert))
	fmt.Println(errorInsert)
}
