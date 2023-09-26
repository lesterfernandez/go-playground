#! /usr/bin/env bash

{ 
    echo "## Benchmarks for Radix Sort  
LsdRadixSort sorts 1 million 32-bit unsigned integers with a radix of 16 (4 bit chunks)  
MsdRadixSort and StringQuickSort sort 1 million random hex strings whose length is 1, 2, or 4  
"; 
    go test -bench . -benchmem ./radix | sed 's/$/  /';
} > README.md