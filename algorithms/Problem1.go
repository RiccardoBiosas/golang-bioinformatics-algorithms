package algorithms

import (
	"log"
)

func algo1(str string) map[string]uint64 {
	if len(str) > 1000 {
		log.Fatal("too long")
	}
	count := map[string]uint64{
		"A": 0,
		"C": 0,
		"G": 0,
		"T": 0,
	}
	for _, char := range str {
		currentKey := string(char)
		if _, ok := count[currentKey]; ok {
			count[currentKey] = count[currentKey] + 1
		}
	}
	return count
}
