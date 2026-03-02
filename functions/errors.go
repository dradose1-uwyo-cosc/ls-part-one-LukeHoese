package functions

import (
	"fmt"
	"os"
)

// error format
func printAccessErr(target string, err error) {
	fmt.Fprintf(os.Stderr, "gols: cannot access '%s': %s\n", target, err.Error())
}
