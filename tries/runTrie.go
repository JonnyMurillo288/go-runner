package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	utils "github.com/Jonnymurillo288/GoUtils"
	trie "github.com/Jonnymurillo288/go-runner/tries"
)

func main() {
	var prefix, match string
	flag.StringVar(&prefix,"prefix","we","Prefix to find in the text")
	flag.StringVar(&match,"match",".i.e","What do you want to match in the text")
	flag.Parse()

	file, err := os.Open("./blue.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	st := trie.NewTrieSymbolTable(256)
	var t int
	for scanner.Scan() {
		line := utils.Strip(scanner.Text())
		for _, word := range strings.Split(line, " ") {
			t++
			fmt.Printf("Putting in %v value of %v\n",word,t)
			st.Put(word, t)
		}
	}
	fmt.Println("Keys:")
	for _,key := range st.Keys() {
		fmt.Println(key, st.Get(key))
	}
	fmt.Println("===================")
	fmt.Printf("Longest Prefix (%v):\n",prefix)
	fmt.Println(st.LongestPrefixOf(prefix))

	fmt.Println("===================")
	fmt.Printf("Keys With Prefix (%v):\n",prefix)
	fmt.Println(st.KeysWithPrefix(prefix))

	fmt.Println("===================")
	fmt.Printf("Keys that Match (%v):\n",match)
	fmt.Println(st.KeysThatMatch(match))


}