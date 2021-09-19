package trie

// hash table for searching for an item
type StringSymbolTable struct {
	Table []int
	Values []string
}

func NewStringSymbolTable() *StringSymbolTable {
	return &StringSymbolTable{
		Table: []int{},
		Values: []string{},
	}
}

func (st *StringSymbolTable) put(key int, value string) {
	st.Table = append(st.Table,key)
	st.Values = append(st.Values, value)
}

func (st *StringSymbolTable) get(key int) string {
	return st.Values[key]
}

func (st *StringSymbolTable) delete(key int) {
	v := st.Table[:key]
	del := st.Table[key]
	st.Table = st.Table[key:]
	st.Table = append(st.Table,v...)
	q := st.Values[:del]
	st.Values = st.Values[del:]
	st.Values = append(st.Values, q...)
}