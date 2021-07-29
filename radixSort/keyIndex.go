package radixsort

import "fmt"

type KeyIndex struct {
	arr []interface{}
	index []int
	count []int
}


func newKeyIndex(arr []interface{}) *KeyIndex {
	k := &KeyIndex{
		arr: arr,
		index: make([]int,0),
		count: make([]int,len(arr)+1),
	}
	for i := 0; i < len(arr); i++ {
		k.index = append(k.index,i)
	}
	return k
}

func (k *KeyIndex) sort() []int {
	var N = len(k.arr)
	var R = len(k.count)-1
	var aux = make([]int,R)

	for i := 0; i < N; i++ {
		// fmt.Println(k.arr[i])
		k.count[k.index[i]+1]++
	}

	for i := 0; i < R; i++ {
		k.count[i+1] += k.count[i]
	}

	for i := 0; i < N; i++ {
		// fmt.Println(len(aux))
		fmt.Println(k.index[i]+1)
		fmt.Println(len(k.count))
		aux[k.count[k.index[i]]+1] = k.index[i]
	}

	for i := 0; i < N; i++ {
		k.arr[i] = aux[i]
	}

	return k.index
}
