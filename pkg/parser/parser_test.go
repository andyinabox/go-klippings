package parser

import (
	"hash/crc32"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestSanity(t *testing.T) {
	t.Log("You appear to be sane")
}

func TestRegExp(t *testing.T) {
	l1 := "Debt: The First 5,000 Years (Graeber, David)"
	l2 := "- Your Highlight Location 181-183 | Added on Monday, January 7, 2013 5:09:10 PM"

	line1 := l1re.FindStringSubmatch(l1)
	line2 := l2re.FindStringSubmatch(l2)

	if line1 == nil {
		t.Fatal("Line 1 regex returned nil")
	}

	if len(line1) != 3 {
		t.Fatalf("Expected 3 results for line 1, got %v: %v", len(line1), line1)
	}

	if line2 == nil {
		t.Fatal("Line 1 regex returned nil")
	}

	if len(line2) != 5 {
		t.Fatalf("Expected 4 results for line 2, got %v: %v", len(line2), line2)
	}

	expectedTitle := "Debt: The First 5,000 Years"
	title := line1[1]
	if title != expectedTitle {
		t.Fatalf("Expected '%v', got '%v'", expectedTitle, title)
	}

	expectedAuthor := "Graeber, David"
	author := line1[2]
	if author != expectedAuthor {
		t.Fatalf("Expected '%v', got '%v'", expectedAuthor, author)
	}

	expectedType := "Highlight"
	cType := line2[1]
	if cType != expectedType {
		t.Fatalf("Expected '%v', got '%v'", expectedType, cType)
	}

	expectedLocationRange := "181-183"
	locationRange := line2[3]
	if locationRange != expectedLocationRange {
		t.Fatalf("Expected '%v', got '%v'", expectedLocationRange, locationRange)
	}

	expectedDate := "Monday, January 7, 2013 5:09:10 PM"
	date := line2[4]
	if date != expectedDate {
		t.Fatalf("Expected '%v', got '%v'", expectedDate, date)
	}

}

func TestParser(t *testing.T) {
	f, err := os.Open("../../test/data/my_clippings.txt")
	if err != nil {
		t.Fatalf("Error opening clippings file: %v", err)
	}

	defer f.Close()

	data, err := Parse(f)
	if err != nil {
		t.Fatalf("Error parsing: %v", err)
	}

	if len(data) != 23 {
		t.Fatalf("Expected 23 results, got %v", len(data))
	} else {
		t.Logf("%v results were returned", len(data))
	}

	c := data[0]

	expectedSource := `Debt: The First 5,000 Years (Graeber, David)
- Your Highlight Location 181-183 | Added on Monday, January 7, 2013 5:09:10 PM

As the great classicist Moses Finley often liked to say, in the ancient world, all revolutionary movements had a single program: “Cancel the debts and redistribute the land.”5`
	if !reflect.DeepEqual(c.Source, expectedSource) {
		t.Fatalf("Expected '%s', got '%s'", expectedSource, c.Source)
	}

	expectedChecksum := crc32.ChecksumIEEE([]byte(expectedSource))
	if c.Checksum != expectedChecksum {
		t.Fatalf("Expected '%v', got '%v'", expectedChecksum, c.Checksum)
	}

	expectedTitle := "Debt: The First 5,000 Years"
	if c.Title != expectedTitle {
		t.Fatalf("Expected '%v', got '%v", expectedTitle, c.Title)
	}

	expectedAuthors := "Graeber, David"
	if c.Authors != expectedAuthors {
		t.Fatalf("Expected '%v', got '%v'", expectedAuthors, c.Authors)
	}

	expectedContent := "As the great classicist Moses Finley often liked to say, in the ancient world, all revolutionary movements had a single program: “Cancel the debts and redistribute the land.”5"
	if c.Content != expectedContent {
		t.Fatalf("Expected '%v', got '%v'", expectedContent, c.Content)
	}

	expectedType := "Highlight"
	if c.Type != expectedType {
		t.Fatalf("Expected '%v', got '%v'", expectedType, c.Type)
	}

	expectedLocationRange := [2]uint32{181, 183}
	if !reflect.DeepEqual(c.LocationRange, expectedLocationRange) {
		t.Fatalf("Expected '%v', got '%v'", expectedLocationRange, c.LocationRange)
	}

	expectedDate, err := time.Parse(TimeFormat, "Monday, January 7, 2013 5:09:10 PM")
	if err != nil {
		t.Fatalf("Error parsing date: %v'", err)
	}
	if c.Date != expectedDate {
		t.Fatalf("Expected '%v', got '%v'", expectedDate, c.Date)
	}
}
