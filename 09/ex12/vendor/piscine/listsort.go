package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	for i, il := 0, l; il != nil; {
		for j, jl := i+1, il.Next; jl != nil; {
			if (i > j && il.Data < jl.Data) || (i < j && il.Data > jl. Data) {
				il.Data, jl.Data = jl.Data, il.Data
			}
			jl = jl.Next
			j++
		}
		il = il.Next
		i++
	}
	return l
}
