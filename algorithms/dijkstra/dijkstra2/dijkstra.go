package main

import (
    "fmt"
    "math"
)

type Graph struct {
    // V Represents the set of vertices in a graph.
    // for some index i, i represents the source vertex (u),
    // V[i] represents the set of outgoing weighted edges,
    // and for each edge in V[i] the key is the destination
    // vertex and the value is the weight
    V []map[int]int
}

type vertexWithWeight struct {
    v int
    pathWeight int
}

// Dijkstra's Shortest Path Algorithm
func ShortestPath(src int, dest int, g Graph) int {
    shortestPaths := make([]int, len(g.V))
    for i := range shortestPaths {
        shortestPaths[i] = math.MaxInt32
    }
    shortestPaths[src] = 0

    var v int
    queue := []vertexWithWeight{
        { src, 0 },
    }
    visited := make(map[int]bool)


    for len(queue) != 0 {
        v, queue = popNextVertex(queue)
        fmt.Printf("visiting %v\n", v)
        if visited[v] {
            continue
        }
        visited[v] = true

        edges := g.V[v]
        for e, w := range edges {
            shortestPaths[e] = min(shortestPaths[e], shortestPaths[v] + w)
            queue = append(queue, vertexWithWeight{
                v: e,
                pathWeight: shortestPaths[e],
            })
        }
    }

    // DEBUG:
    fmt.Printf("shortest paths:\n%+v\n", shortestPaths)

    return shortestPaths[dest]
}

func popNextVertex(q []vertexWithWeight) (int, []vertexWithWeight) {
    if len(q) == 1 {
        return q[0].v, []vertexWithWeight{}
    }
    minWeight := math.MaxInt32
    v := 0
    vIdx := 0
    for i, n := range q {
        if n.pathWeight < minWeight {
            minWeight = n.pathWeight
            v = n.v
            vIdx = i
        }
    }
    q = append(q[:vIdx], q[vIdx+1:]...)
    return v, q
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

