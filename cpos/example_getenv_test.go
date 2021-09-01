package cpos_test

import (
	"fmt"
	"os"

	"github.com/chrispruitt/go-util/cpos"
)

func Example() {

	fmt.Printf("Environment var 'NAME' defaults to: %s\n", cpos.Getenv("NAME", "Bob"))

	os.Setenv("NAME", "Ted")

	fmt.Printf("Environment var 'NAME' has been set to: %s\n", cpos.Getenv("NAME", "Bob"))
}
