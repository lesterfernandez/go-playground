package radix_test

import (
	"encoding/hex"
	"math/rand"
	"testing"

	"github.com/lesterfernandez/go-playground/radix-sort/radix"
)

func BenchmarkLsdRadixSort(b *testing.B) {
	sliceSize := 1_000_000
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := make([]uint, sliceSize)
		for i := 0; i < sliceSize; i++ {
			nums[i] = uint(rand.Int31())
		}
		b.StartTimer()
		radix.BinaryRadixSort(nums, 16)
	}
}

func BenchmarkMsdRadixSort(b *testing.B) {
	sliceSize := 1_000_000
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		strings := make([]string, sliceSize)
		for i := 0; i < sliceSize; i++ {
			b := make([]byte, rand.Intn(4)+1)
			s := hex.EncodeToString(b)
			strings[i] = s
		}
		b.StartTimer()
		radix.MsdRadixSort(strings)
	}
}

func BenchmarkStringQuickSort(b *testing.B) {
	sliceSize := 1_000_000
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		strings := make([]string, sliceSize)
		for i := 0; i < sliceSize; i++ {
			b := make([]byte, rand.Intn(4)+1)
			s := hex.EncodeToString(b)
			strings[i] = s
		}
		b.StartTimer()
		radix.StringQuickSort(strings)
	}
}
