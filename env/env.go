package env

import (
	"fmt"
	"os"
)

func GetOrPanic(name string) string {
	value := os.Getenv(name)
	if value == "" {
		panic(fmt.Sprintf("%s evn not defined", name))
	}
	return value
}
