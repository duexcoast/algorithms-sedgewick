package xsum_test

import (
	"log"
	"os"
	"testing"

	. "github.com/duexcoast/algorithms-sedgewick/fundamentals/xsum"
	"github.com/duexcoast/algorithms-sedgewick/testutil"
)

func readInts(path string) []int {
	in := testutil.NewInReadWords(path)
	return in.ReadAllInts()
}

var (
	ints1k = readInts("testdata/1kints.txt")
	ints2k = readInts("testdata/2kints.txt.gz")
	ints4k = readInts("testdata/4kints.txt.gz")
)

func TestMain(m *testing.M) {
	log.Println("before...")

	r := m.Run()

	log.Println("after...")

	os.Exit(r)
}

func TestThreeSumCount(t *testing.T) {
	// t.Parallel()

	want := []int{70, 528, 4039}

	got := []int{
		ThreeSumCount(ints1k),
		ThreeSumCount(ints2k),
		ThreeSumCount(ints4k),
	}

	for i, v := range want {
		if got[i] != v {
			t.Errorf("got: %v; want: %v", got, want)
		}
	}
}

func TestThreeSumCountFast(t *testing.T) {
	// t.Parallel()

	want := []int{70, 528, 4039}

	got := []int{
		ThreeSumCountFast(ints1k),
		ThreeSumCountFast(ints2k),
		ThreeSumCountFast(ints4k),
	}

	for i, v := range want {
		if got[i] != v {
			t.Errorf("got: %v; want: %v", got, want)
		}
	}
}
