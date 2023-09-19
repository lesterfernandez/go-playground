package radix

import (
	"math"
)

func BinaryRadixSort(nums []uint, radix uint16) {
	if radix < 2 || radix&(radix-1) != 0 {
		panic("Radix must be a base of two for binaryRadixSort")
	}

	mask := uint(radix - 1)
	shiftIncrement := int(math.Log2(float64(radix)))
	aux := make([]uint, len(nums))

	for shift := 0; shift < 32; shift += shiftIncrement {
		count := make([]uint, radix+1)

		for _, num := range nums {
			count[num>>shift&mask+1]++
		}

		for i := 1; i <= int(radix); i++ {
			count[i] += count[i-1]
		}

		for _, num := range nums {
			pos := num >> shift & mask
			aux[count[pos]] = num
			count[pos]++
		}

		copy(nums, aux)
	}
}

func MsdRadixSort(arr []string) {
	msdRadixSort(arr, make([]string, len(arr)), 0, len(arr)-1, 0, 256)
}

func msdRadixSort(arr []string, aux []string, lo int, hi int, d int, radix int) {
	if hi <= lo {
		return
	}

	count := make([]int, radix+2)
	for i := lo; i <= hi; i++ {
		count[msdRadixChar(arr[i], d)+2]++
	}
	for i := 1; i < radix+1; i++ {
		count[i] += count[i-1]
	}
	for i := lo; i <= hi; i++ {
		pos := msdRadixChar(arr[i], d) + 1
		aux[count[pos]] = arr[i]
		count[pos]++
	}
	for i := lo; i <= hi; i++ {
		arr[i] = aux[i-lo]
	}

	for i := 1; i < radix+1; i++ {
		msdRadixSort(arr, aux, lo+count[i-1], lo+count[i]-1, d+1, radix)
	}
}

func msdRadixChar(s string, d int) int {
	if len(s) <= d {
		return -1
	}
	return int(s[d])
}
