# Custom Flags

It's possible to create custom flags in Go by defining methods on
an object such that it satisfies interfaces defined in Go's `flag` package.

See 7.4 in the Go Book. 

# Run

Build and run with rect flag:

```
go build
./custom_flags --rect 5x10
```

Prints the result:

```
Rectangle parsed:

Length: 5, Width: 10

**********
**********
**********
**********
**********
```

### References
_7.4, The Go Programming Language, Alan A.A. Donovan, Brian W. Kernighan_
