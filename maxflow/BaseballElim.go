package maxflow

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Team struct {
	Name         string
	Wins, Losses int
	Remaining    int
	Against      map[string]int
	IsEliminated bool
}

func NewTeam(name string, w int, l int, rem int) Team {
	return Team {
		Name:name,
		Wins:w,
		Losses:l,
		Against: make(map[string]int),
	}
}

type BaseballElim struct {
	Teams map[string]Team
}

func (b *BaseballElim) wins(team string) int {
	return b.Teams[team].Wins
}

func (b *BaseballElim) losses(team string) int {
	return b.Teams[team].Losses
}

func (b *BaseballElim) remaining(team string) int {
	return b.Teams[team].Remaining
}

func (b *BaseballElim) against(team1 string, team2 string) int {
	return b.Teams[team1].Against[team2]
}

// parse the file to build teams
func buildTeams(file string) []Team {
	var against = make(map[string][]int)
	var res []Team
	f,err := os.Open(file)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := trimLine(strings.Split(scanner.Text()," "))
		if len(line) < 4 {
			continue
		}
		w,_ := strconv.Atoi(line[1])
		l,_ := strconv.Atoi(line[2])
		r,_ := strconv.Atoi(line[3])
		tm := NewTeam(line[0],w,l,r)
		res = append(res,tm)
		for _,a := range []string{line[4],line[5],line[6],line[7]} {
			v,_ := strconv.Atoi(a)
			against[tm.Name] = append(against[tm.Name],v)
		}
	}
	for _,team := range res {
		for t,games := range against {
			for _,g := range games {
				team.Against[t] = g
			}
		}
	}
	return res
}

// return the flow network and the int value of the target
func BuildGraph(file string,targetTeam int) (*FlowNetwork,*Team,int) {
	var perm []int
	b := &BaseballElim{
		Teams: make(map[string]Team),
	}
	tms := buildTeams(file)
	if targetTeam > len(tms) {
		log.Fatalln("Error: TargetTeam is greater than the length of teams")
		return nil,nil,-1
	}
	retTeam := tms[targetTeam]
	for i := 0; i < len(tms); i++ {
		perm = append(perm,i)
	} 
	fin := createCombo(perm) // combinations of matchups for all teams
	res := len(fin) + len(tms) + 1  // each matchup is a node and each team will be a node


	target := tms[targetTeam]
	for _,tm := range tms {
		b.Teams[tm.Name] = tm
	}
	network := NewFlowNetwork(res)
	nodes1 := 1 // nodes in the first layer
	nodes2 := len(fin)+1 // nodes in the second layer

	// creates the edges from source to games remaining for each opponent
	passed := []int{}
	for i := 0; i < len(tms); i++ {
		tmi := nodes2+i
		passed = append(passed,i)
		for j := 1; j < len(tms)+1; j++ {
			if targetTeam == j {
				continue
			}
			if nodes1 == nodes2 {
				continue
			}
			if inArr(passed,j) {
				continue
			}
			if j != i {
				tmj := nodes2+j
				// fmt.Println("Node:",nodes1,"\ni:",i,"j",j,"tmi:",tmi,"tmj",tmj)
				e := NewFlowEdge(0,nodes1,float64(tms[i].Against[tms[j-1].Name]))
				// fmt.Printf("connecting %v --> %v\n",0,nodes1)
				network.AddEdge(e)
				ed := NewFlowEdge(nodes1,tmi,math.Inf(1))
				// fmt.Printf("Connecting %v --> %v\n",nodes1,tmi)
				edd := NewFlowEdge(nodes1,tmj,math.Inf(1))
				// fmt.Printf("Connecting %v --> %v\n",nodes1,tmj)
				network.AddEdge(ed)
				network.AddEdge(edd)
				nodes1++
			}
		}
		// fmt.Println("Remaining wins for node:",nodes2+i, "-->",res,":",float64(target.Wins + target.Remaining - tms[i].Wins))
		f := NewFlowEdge(nodes2+i,res,float64(target.Wins + target.Remaining - tms[i].Wins))
		network.AddEdge(f)
		// fmt.Printf("Connecting %v --> %v\n",nodes2+i,res)

		// nodes2++
	}
	return network,&retTeam,res
}

func inArr(arr []int, a int) bool {
	for _,p := range arr {
		if p == a {
			return true
		}
	}
	return false
}

func createCombo(arr []int) [][]int {
	var res [][]int
	for _,a := range arr {
		for _, b := range arr {
			if a != b && !reverse(a,b,res) {
				res = append(res,[]int{a,b})
			}
		}
	}
	return res
}

func reverse(a int, b int, arr [][]int) bool {
	for _,p := range arr {
		i := 0
		for _,m := range p {
			if m == a || m == b {
				i++
				if i == 2 {
					return true
				}
			}
		}
	}
	return false
}

func trimLine(line []string) []string {
	var res []string
	for i := 0; i < len(line); i++ {
		if line[i] != "" {
			res = append(res,line[i])
		}
	}
	return res
}