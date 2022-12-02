package reader

import (
	"os"
)

func Read(filename string) string {
	res, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(res)
}
