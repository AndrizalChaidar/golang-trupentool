package helpers

import (
	"log"
	"strconv"
)

func StringToUint(str string) uint {
	ui64, err := strconv.ParseUint(str, 0, 64)
	if err != nil {
		log.Fatal("Error converting string to uint")
	}
	return uint(ui64)
}
