package persistingdata

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// func ReadTxtFile reads any txt file converts it to a 2d matrix array
func ReadTxTFile(filename string) ([][]string, string) {

	//open txt file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error:", err)
	}

	log.Println("Opening txt file...")

	defer file.Close()

	//new Reader returns a new reader reading from the file
	newReader := csv.NewReader(file)

	//readAll reads all the records and returns an array of array
	records, err := newReader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	//output message
	msg := fmt.Sprintln("Read txt file successful")
	log.Println(msg)

	return records, msg
}

// // func WriteTxtFile writes any 2d matrix to disc as txt file
func WriteTxtFile(records [][]string) string {

	//output filename
	filename := "output_data.txt"

	//Create creates or truncates the named file. If the file already exists, it is truncated. If the file does not exist, it is created
	txtFile, err := os.Create(filename)
	if err != nil {
		log.Fatal("Failed to create file:", err)
	}

	//NewWriter returns a new Writer that writes to w.
	w := csv.NewWriter(txtFile)

	//
	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatal("Error writing record to disc:", err)
		}
	}

	//error handling
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	//flush and close file
	w.Flush()
	txtFile.Close()

	//output message
	msg := fmt.Sprintln("Write txt file to disc successful")

	log.Println(msg)

	return msg
}
