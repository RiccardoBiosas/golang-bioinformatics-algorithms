package algorithms

import (
	"strings"
)

func algo2(str string) string {
	RNA := strings.Replace(str, "T", "U", -1)
	return RNA
}
