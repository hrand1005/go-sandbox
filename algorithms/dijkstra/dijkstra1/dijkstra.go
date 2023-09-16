package main

import (
    "container/heap"
    "math"
)

type MinHeap []VertexWithWeight

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].W < h[j].W }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(VertexWithWeight))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Vertex struct {
    ID int
    Next *Vertex
}

type VertexWithWeight struct {
    V *Vertex
    W int
}

type Graph struct {
    Vertices []*Vertex
    Edges map[*Vertex][]VertexWithWeight
}

func Dijkstras(g Graph, src *Vertex) map[*Vertex]int {
    finished := make(map[*Vertex]bool)
    dist := make(map[*Vertex]int)
    for _, v := range g.Vertices {
        dist[v] = math.MaxInt32
    }

    queue := &MinHeap{}
    heap.Push(queue, VertexWithWeight{
        V: src,
        W: 0,
    })
    dist[src] = 0
    for len(*queue) != 0 {
        u := heap.Pop(queue).(VertexWithWeight)

        finished[u.V] = true
        for _, n := range g.Edges[u.V] {
            if finished[n.V] {
                continue
            }
            pathWeight := u.W + n.W
            if pathWeight < dist[n.V] {
                dist[n.V] = pathWeight
                n.W = pathWeight
                n.V.Next = u.V
                heap.Push(queue, n)
            }
        }
    }
    return dist
}
