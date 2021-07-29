package maxflow

import (
	"fmt"
	"testing"
)


func TestBaseball(t *testing.T) {

	g,tar := buildGraph()
	ff := RunFordFolkerson((*g),0,tar)
	fmt.Println(ff.Value)

}