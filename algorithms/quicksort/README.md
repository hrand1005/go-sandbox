# Quicksort

Go implementation of the quicksort algorithm. Uses generics to sort 'sortable' elements
which are manually defined in `quicksort.go`.

To run:

```
go build
./quicksort
```

Result:

```
Sorting slice with elements of type int:

Before:
        [2 4 5 1 0 -1 10]

After:
        [-1 0 1 2 4 5 10]

------------

Sorting slice with elements of type string:

Before:
        [b hi herb zebra mom apple]

After:
        [apple b herb hi mom zebra]

------------
```


