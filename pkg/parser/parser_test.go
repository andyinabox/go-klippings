package parser

import (
	"os"
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

	clippings, err := Parse(f)
	if err != nil {
		t.Fatalf("Error parsing: %v", err)
	}

	if len(clippings) < 1 {
		t.Fatalf("No clippings were returned")
	} else {
		t.Logf("%v clippings were returned", len(clippings))
	}

	c := clippings[0]

	expectedTitle := "Debt: The First 5,000 Years"
	if c.RawTitle != expectedTitle {
		t.Fatalf("Expected '%v', got '%v", expectedTitle, c.RawTitle)
	}

	expectedAuthors := "Graeber, David"
	if c.RawAuthors != expectedAuthors {
		t.Fatalf("Expected '%v', got '%v'", expectedAuthors, c.RawAuthors)
	}

	expectedContent := "As the great classicist Moses Finley often liked to say, in the ancient world, all revolutionary movements had a single program: “Cancel the debts and redistribute the land.”5"
	if c.RawContent != expectedContent {
		t.Fatalf("Expected '%v', got '%v'", expectedContent, c.RawContent)
	}

	expectedType := "Highlight"
	if c.Type != expectedType {
		t.Fatalf("Expected '%v', got '%v'", expectedType, c.Type)
	}

	if c.LocationRangeStart != 181 || c.LocationRangeEnd != 183 {
		t.Fatalf("Expected '%v-%v', got '%v-%v'", 181, 183, c.LocationRangeStart, c.LocationRangeEnd)
	}

	expectedDate, err := time.Parse(TimeFormat, "Monday, January 7, 2013 5:09:10 PM")
	if err != nil {
		t.Fatalf("Error parsing date: %v", err)
	}
	if c.Date != expectedDate {
		t.Fatalf("Expected '%v', got '%v'", expectedDate, c.Date)
	}
}
