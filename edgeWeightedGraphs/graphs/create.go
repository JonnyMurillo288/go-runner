package graphs

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func CreateEWD() *EdgeWeightDigraph {
    var g *EdgeWeightDigraph
	var e Edge
    file, err := os.Open("./graph.txt")
    scanner := bufio.NewScanner(file)
    if err != nil {
        panic(err)
    }

    for scanner.Scan() {
		points := strings.Split(scanner.Text()," ")
        p1,err := strconv.Atoi(points[0])
        p2,err := strconv.Atoi(points[1])
		w, err := strconv.ParseFloat(points[2],64)
        if err != nil {
			panic(err)
        }
		e = Edge{
			V: p1,
			W: p2,
			Weight: w,
		}
        g.AddEdge(e)
    }
    return g
}

func CreateEWG() *EdgeWeightGraph {
    var g = &EdgeWeightGraph{
        Edges: make([]Edge,0),
        Adj: make(map[int][]Edge),
    }
	var e Edge
    file, err := os.Open("./graph.txt")
    scanner := bufio.NewScanner(file)
    if err != nil {
        panic(err)
    }

    for scanner.Scan() {
		points := strings.Split(scanner.Text()," ")
        p1,err := strconv.Atoi(points[0])
        p2,err := strconv.Atoi(points[1])
		w, err := strconv.ParseFloat(points[2],64)
        if err != nil {
			panic(err)
        }
		e = Edge{
			V: p1,
			W: p2,
			Weight: w,
		}
        g.AddEdge(e)
    }
    return g
}
// Param: t = type of graph to create
// EWD: EdgeWeightDigraph
// EWG: EdgeWeightGraph
func CreateGraph(t string) interface{} {
    if t == "EWD" {
        return CreateEWD()
    } else if t == "EWG" {
        return CreateEWG() 
    } 
    return nil
}