package greetings

import (
	"fmt"

	"rsc.io/quote"
)

func Greetings(name string) string {
	message := fmt.Sprintf("Hello, my name is %s", name)
	println(quote.Glass())
	return message
}
