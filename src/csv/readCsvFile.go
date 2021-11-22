package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// TODO Allow for the option to read semi-colan instead of commas

//--------------------------------------------------
//
//	filePath: Is a string that is the location of the csv file on the
//			 users computer
//
//
//	index: Is an integer that holds the column number for the email address.
//		TODO, In the future will be removed when readCsvFile also grabs
//			data for the "data: interface" in SetBody(location, data).
//
//--------------------------------------------------
func ReadCsvFile(filePath string) map[int]string {

	csvFile, err := os.Open(filePath)

	if err != nil {
		log.Fatalln("Couldn't open the CSV file", err)
	}

	records := csv.NewReader(csvFile)

	emails := make(map[int]string)

	//  Reads the first record and sends it to FindIndexOfEmail to get the
	//  index in the record.
	record, err := records.Read()

	if err == io.EOF {
		return emails
	}

	if err != nil {
		log.Fatal(err)
	}

	var indexColumn int = 0

	indexColumn = FindIndexOfEmail(record)

	//  Returns empty emails if no email address was found in any column of
	//  first record
	if indexColumn == -1 {
		fmt.Println("Error there was no email address found CSV file")

		return emails
	}

	//--------------------------------------------------

	emails[0] = strings.TrimSpace(record[indexColumn])

	var i int = 1

	for {

		record, err := records.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		emails[i] = strings.TrimSpace(record[indexColumn])
		i++
	}

	return emails
}
