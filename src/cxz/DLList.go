package WorkerPool

import "sync"

type dlList struct {
	first *element
	last *element
	current *element
	length uint
	sync.Mutex
}

type DLList interface {
	First() Element
	Last() Element
	Next() Element
	Prev() Element
	Current() Element
	Length() uint
	LPush(func()) //ATTN: insert generic here
	RPush(func()) //ATTN: insert generic here
	LPop() Element
	RPop() Element
	LPeek() Element
	RPeek() Element
	Lock()
	Unlock()
}

func NewDLList() DLList {
	return new(dlList) //NOTE: leaves everything nil or zero
}

func (self *dlList) First() Element {
	self.current = self.first
	return self.current
}

func (self *dlList) Last() Element {
	self.current = self.last
	return self.current
}

func (self *dlList) Next() Element {
	if self.current == nil {
		return nil
	}
	self.current = self.current.next
	return self.current
}

func (self *dlList) Prev() Element {
	if self.current == nil {
		return nil
	}
	self.current = self.current.prev
	return self.current
}

func (self *dlList) Current() Element {
	return self.current
}

func (self *dlList) Length() uint {
	return self.length
}

func (self *dlList) LPush(f func()) {
	self.Lock()

	if self.current == nil {
		self.addFirstElement(f)
	} else {
		self.first.prev = &element{value: f, next: self.first}
		self.first = self.first.prev
	}

	self.Unlock()
}

func (self *dlList) RPush(f func()) {
	self.Lock()

	if self.current == nil {
		self.addFirstElement(f)
	} else {
		self.last.next = &element{value: f, prev: self.last}
		self.last = self.last.next
	}

	self.Unlock()
}

func (self *dlList) LPop() Element {
	self.Lock()

	to_return := self.first
	if self.current == self.first {
		self.current = self.first.next
	}
	self.first = self.first.next

	self.Unlock()

	return to_return
}

func (self *dlList) RPop() Element {
	self.Lock()

	to_return := self.last
	if self.current == self.last {
		self.current = self.last.prev
	}
	self.last = self.last.prev

	self.Unlock()

	return to_return
}

func (self *dlList) LPeek() Element {
	return self.first
}

func (self *dlList) RPeek() Element {
	return self.last
}

func (self *dlList) addFirstElement(f func()) {
	self.current = &element{value: f}
	self.first = self.current
	self.last = self.current
}
