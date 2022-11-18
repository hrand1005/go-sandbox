# Interface Comparisons

## interfaceComparisons

To run, uncomment `interfaceComparisons()` and run:
```
go run main.go
```

Result:
```
true
panic: runtime error: comparing uncomparable type []int
...
```

See The Go Programming Language 7.5 for details.

## nilComparisons

To run, uncomment `nilComparisons()` and run:
```
go run main.go
```

Result:
```
false
true
```

See The Go Programming Language 7.5.1 for details.

### References
_The Go Programming Language, Alan A.A. Donovan, Brian W. Kernighan_
