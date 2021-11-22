package csv

import (
	"testing"
)

// Already Have a error chekc for if the file doesn't exist or doesn't open

func TestReadCsvFile_Successful_Last_Value(t *testing.T) {
	got := ReadCsvFile("tests/test.csv")
	want := make(map[int]string)

	want[0] = "test@test.com"

	if got[0] != want[0] {
		t.Errorf("Got: %v, Wanted: %v", got, want)
	}
}

func TestReadCsvFile_Successful_First_Value(t *testing.T) {
	got := ReadCsvFile("tests/test2.csv")

	want := make(map[int]string)

	want[0] = "test@test.com"

	if got[0] != want[0] {
		t.Errorf("Got: %v, Wanted: %v", got, want)
	}
}

func TestReadCsvFile_Successful_Not_First_Or_Last_Value(t *testing.T) {
	got := ReadCsvFile("tests/test3.csv")

	want := make(map[int]string)

	want[0] = "test@test.com"

	if got[0] != want[0] {
		t.Errorf("Got: %v, Wanted: %v", got, want)
	}
}
