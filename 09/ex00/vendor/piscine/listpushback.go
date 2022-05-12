package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	var new NodeL
	new.Data = data
	new.Next = (*NodeL)(nil)
	if (*l).Head == (*NodeL)(nil) && (*l).Tail == (*NodeL)(nil) {
		(*l).Head = &new
		(*l).Tail = &new
		return
	}
	(*l).Tail.Next = &new
	(*l).Tail = &new
}
