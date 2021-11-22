package csv

import "strings"

//--------------------------------------------------
//
//  Purpose: To find the index value from the record that contains the email address
//
//  record []string: This is the first record recived from the CSV file.
//
//--------------------------------------------------
func FindIndexOfEmail(record []string) int {
	var index int = -1

	var i int = 0

	for i < len(record) {
		var x bool = false

		x = strings.Contains(record[i], "@")

		if x {
			index = i
			break
		} else {
			i++
		}

	}

	return index
}
