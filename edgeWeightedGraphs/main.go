package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Jonnymurillo288/go-runner/edgeWeightedGraphs/graphs"
)

func newGraph() graphs.EdgeWeightGraph {
    g  = graphs.NewGraph()
    file, err := os.Open("./graph.txt")
    scanner := bufio.NewScanner(file)
    if err != nil {
        panic(err)
    }

    for scanner.Scan() {
        points := strings.Split(scanner.Text()," ")
        p1,err := strconv.Atoi(points[0])
        p2,err := strconv.Atoi(points[1])
		w,err := strconv.ParseFloat(points[3],64)
        if err != nil {
            panic(err)
        }
        e := graphs.Edge{
			V: p1,
			W: p2,
		    Weight: w,
		}
		g.AddEdge(e)
    }
    return g
}


func main() {
	g := newGraph()
	fmt.Println(g.Kruskal())
	
}