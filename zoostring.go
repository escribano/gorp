package gorp

import (
	"strings"
	
	"github.com/escribano/camelcase"
)

func CamelToSnake(camel string) string {
	splitted := camelcase.Split(camel)
	var snake string
	for m, p := range splitted {
		if m > 0 {
			snake += "_"
		}
		snake += strings.ToLower(p)
	}
	return snake
}

