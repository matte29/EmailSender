package csv

import (
	"testing"
)

func TestFindIndexOfEmail_Successful(t *testing.T) {
	x := []string{"testingATtest.com",
		"testing@Test.com",
		"testingATtest.com",
		"testingATtest.com"}

	got := FindIndexOfEmail(x)

	want := 1

	if got != want {
		t.Errorf("Got: %d, Wanted: %d", got, want)
	}
}

func TestFindIndexOfEmail_Fail(t *testing.T) {
	x := []string{"testingATtest.com", "testingAtTest.com", "testingATtest.com"}

	got := FindIndexOfEmail(x)

	want := -1

	if got != want {
		t.Errorf("Got: %d, Wanted: %d", got, want)
	}
}

func TestFindIndexOfEmail_With_Empty_Input(t *testing.T) {
	x := []string{}

	got := FindIndexOfEmail(x)

	want := -1

	if got != want {
		t.Errorf("Got: %d, Wanted: %d", got, want)
	}
}

func TestFindIndexOfEmail_Successful_With_One_Input(t *testing.T) {
	x := []string{"testing@test.org"}

	got := FindIndexOfEmail(x)

	want := 0

	if got != want {
		t.Errorf("Got: %d, Wanted: %d", got, want)
	}
}

func TestFindIndexOfEmail_Successful_With_Index_As_Last(t *testing.T) {
	x := []string{"testingATtest.com",
		"Question?",
		"testingATtest.com",
		"testingATtest.com",
		"LOUDNOISES&",
		"notaThing!",
		"test",
		"52",
		"testing@test.com",
	}

	got := FindIndexOfEmail(x)

	want := 8

	if got != want {
		t.Errorf("Got: %d, Wanted: %d", got, want)
	}
}
