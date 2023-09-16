package main

import (
    "fmt"
)

func main() {
    vertices := []*Vertex{
        {ID: 1},
        {ID: 2},
        {ID: 3},
        {ID: 4},
    }

    edges := map[*Vertex][]VertexWithWeight{
        vertices[0]: {
            {
                V: vertices[1],
                W: 10,
            },
            {
                V: vertices[2],
                W: 8,
            },
        },
        vertices[1]: {
            {
                V: vertices[3],
                W: 1,
            },
        },
        vertices[2]: {
            {
                V: vertices[3],
                W: 4,
            },
        },
    }
    graph := Graph{
        Vertices: vertices,
        Edges: edges,
    }
    dist := Dijkstras(graph, vertices[0])
    fmt.Printf("min distance from 1 to 4: %v\n", dist[vertices[3]])
}
