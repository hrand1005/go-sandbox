package main

import (
  "fmt"
  "sort"
  "math/rand"
)

var adjacencyMatrix = [][]int{
  {0, 1, 0, 0},
  {1, 0, 1, 1},
  {0, 1, 0, 1},
  {0, 1, 1, 0},
}


func main() {
  printGraphInfo(adjacencyMatrix)
}

func printGraphInfo(graph [][]int) {
  topRow, matrix, sep, vertices := "   ", "", "-", "" 
  edges := map[string]int{}

  for i, v := range graph {
    topRow += fmt.Sprintf("%v ", i)
    matrix += fmt.Sprintf("%v %+v\n", i, v)
    sep += sep
    vertices += fmt.Sprintf("(%v) ", i)
    
    for j := range v {
      if graph[i][j] == 1 {
        e := []int{i, j}
        sort.Ints(e)
        key := fmt.Sprintf("(%v, %v)", e[0], e[1])
        edges[key]++
      }
    }
  }

  fmt.Printf("Commencing DFS Demo...\n\n%v\n\n", sep)
  fmt.Println("Here is a graph,\nrepresented by an\nadjacency matrix:\n")
  fmt.Printf("%v\n%v\n", topRow, matrix)
  fmt.Printf("There are\n%v vertices:\n\n%s\n\n", len(graph), vertices)
  fmt.Printf("There are\n%v edges:\n\n", len(edges))

  for k, _ := range edges {
    fmt.Printf("%v ", k)
  }

  fmt.Println("\n")
  fmt.Printf("Let's use DFS\nto traverse\neach vertex\nexactly once.\n\n")

  start := rand.Intn(len(graph))

  fmt.Printf("Using random\nstart vertex:\n\n(%v)\n\n", start)
  fmt.Println("Traversing...")
}
