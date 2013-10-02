package WorkerPool

type element struct {
	next *element
	prev *element
	value func()
}

type Element interface {
	//Next() Element
	//Prev() Element
	Value() func() //ATTN: this is where generics would come into play
}

/*
func (self *element) Next() Element {
	return self.next
}
*/

/*
func (self *element) Prev() Element {
	return self.prev
}
*/

func (self *element) Value() func() {
	return self.value
}
