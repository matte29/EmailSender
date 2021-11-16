package csv

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

// TODO Allow for the option to read semi-colan instead of commas

//------------------------------------------------------------
//
//	filePath: Is a string that is the location of the csv file on the
//			 users computer
//
//
//	index: Is an integer that holds the column number for the email address.
//		TODO, In the future will be removed when readCsvFile also grabs data
//			 for the "data: interface" in SetBody(location, data).
//
//------------------------------------------------------------

func readCsvFile(filePath string, indexColumn int) map[int]string {

	csvFile, err := os.Open(filePath)

	if err != nil {
		log.Fatalln("Couldn't open the CSV file", err)
	}

	records := csv.NewReader(csvFile)

	var i int = 0

	emails := make(map[int]string)

	for {

		record, err := records.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		emails[i] = record[indexColumn]
		i++
	}

	return emails
}
