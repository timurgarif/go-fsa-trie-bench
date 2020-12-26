package bench

import (
	"bufio"
	"bytes"
	"math/rand"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"

	"github.com/armon/go-radix"
	"github.com/couchbase/vellum"
	iradix "github.com/hashicorp/go-immutable-radix"
	smhanov_dawg "github.com/smhanov/dawg"
	"github.com/tchap/go-patricia/patricia"
)

type enLemPos struct {
	s  string
	ps []byte
}

// Capacity for the en lem pos sample, limit to it if the raw data set is larger
const (
	elpCap    = 82000
	RandCount = 100
)

var (
	Randkeys [][]byte
	elp      []enLemPos = make([]enLemPos, elpCap, elpCap)
)

func init() {
	PrepareEnLemPosSample()
	Randkeys = getRandKeys(RandCount)
}

// Encodes up to the first 8 bytes of given slice to uint64
// binary package not suitable: it requires exactly 8 bytes slice on input
func byteSliceToUint64(s []byte) (u uint64) {
	for i, v := range s {
		if i > 7 {
			break
		}
		u |= uint64(v) << ((7 - i) * 8)
	}

	return
}

func PrepareEnLemPosSample() {
	split := regexp.MustCompile(`\s+`)
	scanner := bufio.NewScanner(strings.NewReader(sampleEnLemPos))
	for j := 0; scanner.Scan(); {
		s := split.Split(strings.TrimSpace(scanner.Text()), -1)

		// If data row is OK
		if len(s) > 1 {
			ps := make([]byte, len(s[1:]))

			for i, v := range s[1:] {
				iv, _ := strconv.ParseInt(v, 10, 8)
				ps[i] = byte(iv)
			}

			elp[j] = enLemPos{s[0], ps}
			j++

			if j >= elpCap {
				// set  elp len to actual number of items set
				elp = elp[:j]
				break
			}
		}
	}

	// Sort elp
	sort.Slice(elp, func(i, j int) bool {
		return elp[i].s < elp[j].s
	})

	sampleEnLemPos = ""
	runtime.GC()
}

func getRandKeys(count int) [][]byte {
	keys := make([][]byte, count, count)
	elpLen := len(elp)

	for i := 0; i < count; i++ {
		r := rand.Intn(elpLen)
		keys[i] = []byte(elp[r].s)
	}

	return keys
}

func BuildFstVellum() *vellum.FST {
	var buf bytes.Buffer
	builder, err := vellum.New(&buf, nil)
	if err != nil {
		panic(err)
	}

	for _, v := range elp {
		err := builder.Insert([]byte(v.s), byteSliceToUint64(v.ps))
		if err != nil {
			panic(err)
		}
	}
	builder.Close()

	fst, err := vellum.Load(buf.Bytes())
	if err != nil {
		panic(err)
	}

	return fst
}

func BuildPatricia() *patricia.Trie {
	trie := patricia.NewTrie()

	for _, v := range elp {
		trie.Insert([]byte(v.s), byteSliceToUint64(v.ps))
	}

	return trie
}

func BuildMap() map[string]uint64 {
	m := make(map[string]uint64)

	for _, v := range elp {
		m[v.s] = byteSliceToUint64(v.ps)
	}

	return m
}

func BuildSmhanovDawgNoVal() smhanov_dawg.Finder {
	builder := smhanov_dawg.New()

	for _, v := range elp {
		builder.Add(v.s)
	}

	return builder.Finish()
}

func BuildHashicorpIradix() *iradix.Tree {
	r := iradix.New()

	for _, v := range elp {
		r, _, _ = r.Insert([]byte(v.s), byteSliceToUint64(v.ps))
	}

	return r
}

func BuildArmonRadixTrie() *radix.Tree {
	r := radix.New()

	for _, v := range elp {
		_, _ = r.Insert(v.s, byteSliceToUint64(v.ps))
	}

	return r
}
