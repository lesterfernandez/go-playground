package radix

import (
	"fmt"
	"math"
)

func BinaryRadixSort(nums []uint, radix uint16) []uint {
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
			loc := num >> shift & mask
			aux[count[loc]] = num
			count[loc]++
		}

		copy(nums, aux)
	}

	return nums
}

func MsdRadixSort(arr []string) []string {
	fmt.Println(arr)
	return arr
}
