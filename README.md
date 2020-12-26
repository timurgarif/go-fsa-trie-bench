# go-fsa-trie-bench
Radix trie and FSA build and lookup benchmarks on English lemma list (~80k entries).

Currently eveluated implementations:
- "github.com/armon/go-radix" (a radix trie)
- "github.com/couchbase/vellum" (FST)
- "github.com/hashicorp/go-immutable-radix" (an immutable radix trie)
- "github.com/smhanov/dawg" (DAWG)
- "github.com/tchap/go-patricia/patricia" (a Patricia trie)

Also standard Go map is evaluated (it supports both build and lookup) just for insights.

## Notes
For the DAWG the sampling values (uint64 for other structs) are not used. So the memory footprint could be a bit large for it. 

The values cannot be added via the explicit API of the DAWG package, but it could be encoded as a suffx of the corresponding key. But this particular detail is out of my interest so far for the benchmark.

## Benchmarks
```
BenchmarkBuildMap-8                           67          16995949 ns/op         7763578 B/op       2440 allocs/op
BenchmarkLookupMap-8                      466863              2319 ns/op               0 B/op          0 allocs/op
BenchmarkBuildFstVellum-8                     13          79263268 ns/op        21808056 B/op     426844 allocs/op
BenchmarkLookupFstVellum-8                 19183             61066 ns/op           16000 B/op        100 allocs/op
BenchmarkBuildPatricia-8                      25          47764423 ns/op        24757191 B/op     522027 allocs/op
BenchmarkLookupPatricia-8                  59445             20539 ns/op               0 B/op          0 allocs/op
BenchmarkBuildSmhanovDawgNoVal-8               3         337258668 ns/op        68881608 B/op    1689707 allocs/op
BenchmarkLookupSmhanovDawgNoVal-8           5244            233610 ns/op            1248 B/op        200 allocs/op
BenchmarkBuildHashicorpIradix-8                3         357529713 ns/op        231748285 B/op   3633714 allocs/op
BenchmarkLookupHashicorpIradix-8           34830             34602 ns/op               0 B/op          0 allocs/op
BenchmarkBuildArmonRadixTrie-8                24          46933478 ns/op        15257619 B/op     472478 allocs/op
BenchmarkLookupArmonRadixTrie-8            34176             33263 ns/op               0 B/op          0 allocs/op
```

## Memory stats
https://golang.org/pkg/runtime/#MemStats

Two stats are shown for each implementation. 

Actual memory footprint ~= `Alloc after structure build (populated)` - `Alloc before build`


```
MEMSTAT STD MAP
Mem statbefore build
Alloc = 7283 Kb TotalAlloc = 32361 Kb   Sys = 72529 Kb
Mem stat after build
Alloc = 10838 Kb        TotalAlloc = 39936 Kb   Sys = 72785 Kb
-----------------
MEMSTAT VELLUMFST
Mem stat before build
Alloc = 7281 Kb TotalAlloc = 32287 Kb   Sys = 72273 Kb
Mem stat after build
Alloc = 8307 Kb TotalAlloc = 53589 Kb   Sys = 72529 Kb
-----------------
MEMSTAT PATRICIA
Mem statbefore build
Alloc = 7282 Kb TotalAlloc = 32249 Kb   Sys = 72017 Kb
Mem stat after build
Alloc = 27894 Kb        TotalAlloc = 56429 Kb   Sys = 72977 Kb
-----------------
MEMSTAT SMHANOV_DAWG
Mem statbefore build
Alloc = 7280 Kb TotalAlloc = 32324 Kb   Sys = 72017 Kb
Mem stat after build
Alloc = 7802 Kb TotalAlloc = 99597 Kb   Sys = 73233 Kb
-----------------
MEMSTAT HASHICORP_IRADIX
Mem stat before build
Alloc = 7279 Kb TotalAlloc = 32251 Kb   Sys = 72785 Kb
Mem stat after build
Alloc = 40044 Kb        TotalAlloc = 258570 Kb  Sys = 142132 Kb
-----------------
MEMSTAT ARMON_RADIXTRIE
Mem stat before build
Alloc = 7283 Kb TotalAlloc = 32362 Kb   Sys = 72017 Kb
Mem stat after build
Alloc = 17510 Kb        TotalAlloc = 47264 Kb   Sys = 72849 Kb
```

