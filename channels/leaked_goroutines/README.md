# Leaked Goroutines

Leaked goroutines are goroutines that don't terminate after they're no longer 
needed. The example in `main.go` uses an unbuffered channel to perform a query.
Three goroutines are spawned, but only the first result returned is used in the
main goroutine. Since the two "slower" goroutines continue executing after the
main goroutine has already received on the channel, they are considered "leaked"
goroutines.

Run the program:
```
go run main.go
```

Observe from the output that even though three goroutines are spawned, the 
exit statement is printed only once:
```
exiting...
result
```
