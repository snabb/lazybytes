package lazybytes_test

import (
	"fmt"
	"github.com/snabb/lazybytes"
	"os"
)

func Example() {
	f := func() []byte {
		fmt.Println("initializing")
		return []byte("hello world\n")
	}
	lr := lazybytes.NewReader(f)
	fmt.Println("not yet initialized")
	fmt.Println("len =", lr.Len())
	lr.WriteTo(os.Stdout)
	// Output:
	// not yet initialized
	// initializing
	// len = 12
	// hello world
}
