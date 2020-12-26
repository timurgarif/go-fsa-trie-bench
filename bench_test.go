package bench

import (
	"testing"
)

func BenchmarkBuildMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildMap()
	}
}

func BenchmarkLookupMap(b *testing.B) {
	m := BuildMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range Randkeys {
			_ = m[string(k)]
		}
	}
}

func BenchmarkBuildFstVellum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildFstVellum()
	}
}

func BenchmarkLookupFstVellum(b *testing.B) {
	fst := BuildFstVellum()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range Randkeys {
			_, _, err := fst.Get(k)
			if err != nil {
				panic(err)
			}
		}
	}
}

func BenchmarkBuildPatricia(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildPatricia()
	}
}

func BenchmarkLookupPatricia(b *testing.B) {
	trie := BuildPatricia()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range Randkeys {
			_ = trie.Get(k)
		}
	}
}

func BenchmarkBuildSmhanovDawgNoVal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildSmhanovDawgNoVal()
	}
}

func BenchmarkLookupSmhanovDawgNoVal(b *testing.B) {
	dawg := BuildSmhanovDawgNoVal()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range Randkeys {
			_ = dawg.IndexOf(string(k))
		}
	}
}

func BenchmarkBuildHashicorpIradix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildHashicorpIradix()
	}
}

func BenchmarkLookupHashicorpIradix(b *testing.B) {
	rt := BuildHashicorpIradix()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range Randkeys {
			_, _ = rt.Get(k)
		}
	}
}

func BenchmarkBuildArmonRadixTrie(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildArmonRadixTrie()
	}
}

func BenchmarkLookupArmonRadixTrie(b *testing.B) {
	rt := BuildArmonRadixTrie()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range Randkeys {
			_, _ = rt.Get(string(k))
		}
	}
}
