## Benchmarks for Radix Sort  
LsdRadixSort sorts 1 million 32-bit unsigned integers with a radix of 16 (4 bit chunks)  
MsdRadixSort and StringQuickSort sort 1 million random hex strings whose length is 1, 2, or 4  

goos: darwin  
goarch: arm64  
pkg: github.com/lesterfernandez/go-playground/radix-sort/radix  
BenchmarkLsdRadixSort-8      	      91	  13162579 ns/op	 8004736 B/op	       9 allocs/op  
BenchmarkMsdRadixSort-8      	      24	  42864672 ns/op	16027904 B/op	      10 allocs/op  
BenchmarkStringQuickSort-8   	      87	  14282705 ns/op	       0 B/op	       0 allocs/op  
PASS  
ok  	github.com/lesterfernandez/go-playground/radix-sort/radix	11.651s  
