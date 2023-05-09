// Package testutil reads in data of various types from standard input, files
// and URLs.
package testutil

import (
	"bufio"
	"compress/gzip"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// ErrEmpty indicates that an empty argument has been provided.
var ErrEmpty = errors.New("argument is empty")

// In provides a convenient interface for processing input streams from various
// sources.
type In struct {
	reader     io.Reader
	scanner    *bufio.Scanner
	hasScanned bool
	hasNext    bool
}

func NewInWithError(uri string) (*In, error) {
	r, err := newReader(uri)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(r)

	return &In{reader: r, scanner: scanner}, nil
}

// Factory method for new In instance. Wrapper for NewInReadLines
func NewIn(uri string) *In {
	return NewInReadLines(uri)
}

// NewInReadWords returns a new In instance with Scanner SplitFunc set to
// ScanWords
func NewInReadWords(uri string) *In {
	r, err := newReader(uri)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(r)

	scanner.Split(bufio.ScanWords)

	return &In{reader: r, scanner: scanner}
}

// NewInReadLines function returns a new In instance.
func NewInReadLines(uri string) *In {
	r, err := newReader(uri)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(r)

	return &In{reader: r, scanner: scanner}
}

func newReader(uri string) (io.Reader, error) {
	if uri == "" {
		return nil, ErrEmpty
	}

	// first try to read file from local file system
	f, err := os.Open(uri)
	if err == nil {
		// is it a gzip?
		if strings.HasSuffix(uri, ".gz") {
			gz, err := gzip.NewReader(f)
			if err != nil {
				panic(err)
			}
			return gz, nil
		} else {
			return f, nil
		}
	} else {
		// URL from web
		resp, err := http.Get(uri)
		if err != nil {
			return nil, err
		}
		return resp.Body, nil
	}
}

// ReadString method reads the next token from this input stream and returns it
// as a string.
func (in *In) ReadString() string {
	in.next()
	return in.scanner.Text()
}

// ReadLine method reads the next token from this input stream and returns it as a
// string.
func (in *In) ReadLine() string {
	in.next()
	return in.scanner.Text()
}

// ReadInt reads the next token from this input stream, parses it as an int, and
// returns the int.
func (in *In) ReadInt() int {
	i, _ := strconv.Atoi(in.ReadString())
	return i
}

// ReadFloat reads the next token from this input stream, parses it as a float64,
// and returns the float64.
func (in *In) ReadFloat() float64 {
	f, _ := strconv.ParseFloat(in.ReadString(), 64)
	return f
}

// IsEmpty method returns true if input stream is empty (except possibly whitespace).
// Use this to know whether the next call to ReadString(), ReadInt(), etc. will
// succeed.
func (in *In) IsEmpty() bool {
	return !in.HasNext()
}

// HasNext method returns true if this input stream has more input (including whitespace).
func (in *In) HasNext() bool {
	if !in.hasScanned {
		in.hasNext = in.scanner.Scan()
		in.hasScanned = true
	}
	return in.hasNext
}

// next method calls in.scanner.Scan() and returns true if a token has been scanned
// it returns false otherwise. After scanning, the token is available through
// in.scanner.Text()
func (in *In) next() bool {
	// check the hasScanned flag to see if there has been a previous scan that
	// returned false, if the flag is set to true then return the hasNext
	// flag which will be false.
	if in.hasScanned {
		in.hasScanned = false
		return in.hasNext
	}

	return in.scanner.Scan()
}

// ReadAll reads and returns the remainder of this input stream, as a string.
func (in *In) ReadAll() string {
	data, err := io.ReadAll(in.reader)
	if err != nil {
		panic(err)
	}

	return string(data)
}

// ReadAllLines method reads all remaining lines from this input stream and returns
// them as an array of strings.
func (in *In) ReadAllLines() []string {
	ss := make([]string, 0)

	for in.HasNext() {
		ss = append(ss, in.ReadString())
	}

	return ss
}

// ReadAllStrings method reads all remaining tokens from this input stream and
// returns them as an array of strings
func (in *In) ReadAllStrings() []string {
	str := in.ReadAll()

	return strings.Fields(str)
}

// ReadAllInts method reads all remaining tokens from this input stream, parses
// them as ints and returns them as an array of ints.
func (in *In) ReadAllInts() []int {
	strSlice := in.ReadAllStrings()
	n := len(strSlice)
	intSlice := make([]int, n)

	for i := range strSlice {
		intSlice[i], _ = strconv.Atoi(strSlice[i])
	}
	return intSlice
}
