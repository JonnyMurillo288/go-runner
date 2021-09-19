package compression

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"

	utils "github.com/Jonnymurillo288/GoUtils"
)

type Node struct {
	Char rune
	Freq int
	Left, Right *Node
}

func NewNode(char rune, freq int, left *Node, right *Node) *Node {
	return &Node{
		Char: char,
		Freq: freq,
		Left: left,
		Right: right,
	}
}

func (n *Node) isLeaf() bool {
	return n.Left == nil && n.Right == nil
}

func (n *Node) compareTo(that *Node) int {
	return n.Freq - that.Freq
}

func readBool(in *bufio.Reader) (byte, bool) {
	c,err := in.ReadByte()
	return c,err == io.EOF 
}

func Compress() {
	compress := "Hello World!"
	// p := make([]byte,len(compress)-1)
	freq := make([]int,257)
	for _,l := range compress {
		freq[l]++
	}
	root := BuildTrie(freq)

	st := make([]string,257)
	st = buildCode(st,root,"")

	writeTrie(root)

	fmt.Println("Len of uncompressed message:",len(compress))

	for i := 0; i < len(compress); i++ {
		code := st[compress[i]]
		for j := 0; j < len(code); j++ {
			if code[j] == '0' {
				writeFile(0)
			} else if code[j] == '1' {
				writeFile(1)
			}
		}
	}

}

func buildCode(st []string, x *Node, s string) []string {
	if !x.isLeaf() {
		buildCode(st,x.Left,s+"0")
		buildCode(st,x.Right,s+"1")
	} else {
		st[x.Char] = s
	}
	return st
}

func Expand() {
	expansion := "1110001010110111001111101100010001101"
	p := make([]byte,8)
	buf := bytes.NewBufferString(expansion)
	in := bufio.NewReader(buf)
	in2 := in
	N,err := io.ReadFull(in,p)
	if err != nil {
		fmt.Println(err)
	}
	root := readTrie(in)
	fmt.Println("N:",N,"root:",root.Char)
	for i := 0; i < N; i++ {
		x := root
		if !x.isLeaf() {
			fmt.Println("Not a leaf")
			_,em := readBool(in2)
			if !em{
				x = x.Left
			} else {
				x = x.Right
			}
		}
		writeFile(x.Char)
	}
}


func writeTrie(x *Node) {
	if x.isLeaf() {
		fmt.Println(string(x.Char))
		return 
	}
	writeTrie(x.Left)
	writeTrie(x.Right)

}

func readTrie(in *bufio.Reader) *Node {
	c,leaf := readBool(in)
	if leaf { // if not EOF
		return NewNode(rune(c),0,nil,nil)
	}
	x := readTrie(in)
	y := readTrie(in)
	return NewNode('0',0,x,y)
}

// build a trie given the frequencies with the Huffman algo
// take two smallest freq and make a trie
// add that Node and freq to the queue and repeat until we have one trie
func BuildTrie(freq []int) *Node{
	R := 256
	pq := utils.NewIndexMinPQ(R*2)
	for i := 0; i < R; i++ {
		if freq[i] > 0 {
			pq.Insert(i,NewNode(rune(i),freq[i],nil,nil),float64(freq[i]))
		}
	}
	for !pq.IsEmpty() {
		R++
		_,xx := pq.DelMin()
		_,yy := pq.DelMin()
		switch xx.(type){
		case nil:
			break
		}
		switch yy.(type){
		case nil:
			return xx.(*Node)
		}
		x := xx.(*Node)
		y := yy.(*Node)
		newFreq := x.Freq+y.Freq
		parent := NewNode('0',newFreq,x,y)
		pq.Insert(R,parent,float64(newFreq))
	}
	_,res := pq.DelMin()
	return res.(*Node)
}

func readNextBytes(file *os.File, lim int) []byte {
	by := make([]byte,lim)
	_,err := file.Read(by)
	if err != nil {
		log.Fatal(err)
	}
	return by
}

func writeFile(item interface{}) {
	fmt.Println("Writing file for:",item)
	txtFile,err := os.OpenFile("./expand.txt",os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	file,err := os.OpenFile("./test.bin",os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	defer txtFile.Close()
	for i := 0; i < 10; i++ {
		var buf bytes.Buffer
		binary.Write(&buf, binary.BigEndian, item)
		//b :=bin\_buf.Bytes()
		//l := len(b)
		//fmt.Println(l)
		writeNextBytes(file, buf.Bytes())
	}
	txtFile.WriteString(fmt.Sprint(item))
}

func writeNextBytes(file *os.File, bytes []byte) {
	_, err := file.Write(bytes)

	if err != nil {
		log.Fatal(err)
	}

}