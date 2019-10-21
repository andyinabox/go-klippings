package parser

import (
	"bufio"
	"bytes"
	"errors"
	. "github.com/andyinabox/go-klippings-api/pkg/types"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	Delimiter        = "\n=========="
	FirstLineRegExp  = `^(.+) \((.+)\)$`
	SecondLineRegExp = `^-\s(?:Your\s)?(\w+) (?:on page ([0-9-]*?) \| )?(?:Loc(?:ation|\.) ([0-9-]*?) +\| )?Added on (.*)$`
	TimeFormat       = "Monday, January 2, 2006 3:04:05 PM"
)

// first line regular expression
var l1re = regexp.MustCompile(FirstLineRegExp)

// second line regular expression
var l2re = regexp.MustCompile(SecondLineRegExp)

func Parse(r io.Reader) ([]Clipping, error) {
	var clippings []Clipping

	scanner := bufio.NewScanner(r)
	scanner.Split(scanClippings)

	for scanner.Scan() {
		c, err := parseChunk(scanner.Bytes())
		if err == nil {
			clippings = append(clippings, c)
		}
	}

	return clippings, nil
}

func parseChunk(b []byte) (Clipping, error) {
	var lines []string
	var clipping Clipping

	// create scanner for chunk bytes
	scanner := bufio.NewScanner(bytes.NewReader(b))
	scanner.Split(bufio.ScanLines)

	// scan through lines, ignoring empty ones
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(dropCR(line)) > 0 {
			lines = append(lines, string(dropCR(line)))
		}
	}

	// if there's less than 2 lines, we don't have enough
	if len(lines) < 2 {
		return clipping, errors.New("Not a complete clipping")
	}

	// create a new clipping struct
	clipping = Clipping{
		Raw: string(b),
	}

	// anything after the first two lines is part
	// of the clipping content
	if len(lines) > 2 {
		clipping.RawContent = strings.Join(lines[2:], " ")
	}

	// match data for first line
	l1 := l1re.FindStringSubmatch(lines[0])
	if l1 != nil {
		// title
		clipping.RawTitle = l1[1]
		// authors
		clipping.RawAuthors = l1[2]
	}

	// match data for second line
	l2 := l2re.FindStringSubmatch(lines[1])
	if l2 != nil {
		clipping.Type = l2[1]

		clipping.PageRangeStart, clipping.PageRangeEnd = parseRange(l2[2])
		clipping.LocationRangeStart, clipping.LocationRangeEnd = parseRange(l2[3])

		if t, err := time.Parse(TimeFormat, l2[4]); err == nil {
			clipping.Date = t
		}
	}

	return clipping, nil
}

func parseRange(r string) (num1 uint, num2 uint) {
	var n1 uint
	var n2 uint

	if r == "" {
		return n1, n2
	}

	a := strings.Split(r, "-")

	if len(a) > 0 {
		// first part of range
		if n, err := strconv.Atoi(a[0]); err == nil {
			n1 = uint(n)
		}
		// second part of range
		if len(a) > 1 {
			if n, err := strconv.Atoi(a[1]); err == nil {
				n2 = uint(n)
			}
		}
	}

	return n1, n2
}

func scanClippings(data []byte, atEOF bool) (advance int, token []byte, err error) {

	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.Index(data, []byte(Delimiter)); i >= 0 {
		// We have a full newline-terminated line.
		return i + len(Delimiter), data[0:i], nil

	}

	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}

	// Request more data.
	return 0, nil, nil

}

// dropCR drops a terminal \r from the data.

func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}
