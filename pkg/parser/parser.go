package parser

import (
	"bufio"
	"bytes"
	"errors"
	"hash/crc32"
	"io"
	"log"
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

type Data struct {
	Title           string
	TitleChecksum   uint32
	Authors         map[string]uint32
	Content         string
	ContentChecksum uint32
	LocationRange   [2]uint32
	PageRange       [2]uint32
	Type            string
	Date            time.Time
	Source          string
	SourceChecksum  uint32
}

// regular expressions
var l1re = regexp.MustCompile(FirstLineRegExp)
var l2re = regexp.MustCompile(SecondLineRegExp)

func Parse(r io.Reader) ([]Data, error) {
	var data []Data

	scanner := bufio.NewScanner(r)
	scanner.Split(ScanClippings)

	for scanner.Scan() {
		d, err := ParseChunk(scanner.Bytes())
		if err == nil {
			data = append(data, d)
		} else {
			log.Printf("Error parsing chunk, %v\n", err)
		}
	}

	return data, nil
}

func parseAuthors(b []byte) (map[string]uint32, error) {
	m := make(map[string]uint32)

	authors := bytes.Split(b, []byte(";"))

	if len(authors) < 1 {
		return m, errors.New("No authors found")
	}

	for _, a := range authors {
		a = bytes.TrimSpace(a)
		cs := crc32.ChecksumIEEE(a)
		s := string(a)
		m[s] = cs
	}

	return m, nil
}

func ParseChunk(b []byte) (Data, error) {
	var lines [][]byte

	var d = Data{
		Source:         string(b),
		SourceChecksum: crc32.ChecksumIEEE(b),
	}

	// create scanner for chunk bytes
	scanner := bufio.NewScanner(bytes.NewReader(b))
	scanner.Split(bufio.ScanLines)

	// scan through lines, ignoring empty ones
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}

	// if there's less than 2 lines, we don't have enough
	if len(lines) < 2 {
		return d, errors.New("Not a complete clipping")
	}

	// anything after the first two lines is part
	// of the clipping content
	if len(lines) > 2 {
		contentBytes := bytes.Join(lines[2:], []byte(" "))
		d.ContentChecksum = crc32.ChecksumIEEE(contentBytes)
		d.Content = string(contentBytes)
	}

	// match data for first line
	l1 := l1re.FindSubmatch(lines[0])
	if l1 != nil {
		// title
		d.TitleChecksum = crc32.ChecksumIEEE(l1[1])
		d.Title = string(l1[1])
		// authors
		authors, err := parseAuthors(l1[2])
		if err == nil {
			d.Authors = authors
		}
	}

	// match data for second line
	l2 := l2re.FindSubmatch(lines[1])
	if l2 != nil {
		d.Type = string(l2[1])

		d.PageRange = parseRange(l2[2])
		d.LocationRange = parseRange(l2[3])

		if t, err := time.Parse(TimeFormat, string(l2[4])); err == nil {
			d.Date = t
		}
	}

	return d, nil
}

func parseRange(b []byte) [2]uint32 {
	var a [2]uint32

	if b == nil {
		return a
	}

	// convert to array of strings
	s := strings.Split(string(b), "-")

	if len(s) > 0 {
		// first part of range
		if n, err := strconv.Atoi(s[0]); err == nil {
			a[0] = uint32(n)
		}
		// second part of range
		if len(s) > 1 {
			if n, err := strconv.Atoi(s[1]); err == nil {
				a[1] = uint32(n)
			}
		}
	}

	return a
}

func ScanClippings(data []byte, atEOF bool) (advance int, token []byte, err error) {

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
