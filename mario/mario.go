package mario

import (
	"errors"
	"strings"
)

func Mario(num int) (string, error) {
	var pyramid string
	if num < 1 || num > 8 {
		return "", errors.New("escoge un numero valido")
	} else {
		for i := 1; i <= num; i++ {
			var hash, spaces, line string
			spaces = strings.Repeat(" ", num-i)
			hash = strings.Repeat("#", i)
			line = spaces + hash + "\n"
			pyramid += line
		}
		return pyramid, nil
	}
}
