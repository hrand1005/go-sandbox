package main

import (
    "fmt"
)

func main() {
    g := Graph{
        V: []map[int]int{
            // A
            {
                // B
                1: 4,
                // C
                2: 2,
            },
            // B
            {
                // C
                2: 3,
                // D
                3: 2,
                // E
                4: 3,
            },
            // C
           {
                // B
                1: 1,
                // D
                3: 4,
                // E
                4: 5,
            },
            // D
            {},
            // E
            {
                // D
                3: 1,
            },
        },
    }
    // shorteset path from A -> E
    weight := ShortestPath(0, 4, g)
    fmt.Printf("shortest path from A -> E: %v\n", weight)
}
