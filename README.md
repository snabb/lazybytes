lazybytes
=========

[![GoDoc](https://godoc.org/github.com/snabb/lazybytes?status.svg)](https://godoc.org/github.com/snabb/lazybytes)

The Go package lazybytes implements a bytes.Reader which is initialized
lazily on first access.

Documentation:

https://godoc.org/github.com/snabb/lazybytes

Simple example:
```
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
```

The Git repository is located at: https://github.com/snabb/lazybytes

License
-------

MIT
