package WorkerPool

import "sync"

type dlList struct {
	first *element
	last *element
	current *element
	length int
	mod_locker sync.Mutex
}

type DLList interface {
	First() Element
	Last() Element
	Next() Element
	Prev() Element
	Current() Element
	Length() int
	LPush(func()) //ATTN: insert generic here
	RPush(func()) //ATTN: insert generic here
	LPop() Element
	RPop() Element
	LPeek() Element
	RPeek() Element
}

func NewRLList() DLList {
	return new(dlList) //NOTE: leaves everything nil or zero
}

func (self *rlList) First() Element {
	self.current = self.first
	return self.current
}

func (self *rlList) Last() Element {
	self.current = self.last
	return self.current
}

func (self *rlList) Next() Element {
	if self.current == nil {
		return nil
	}
	self.current = self.current.Next()
	return self.current
}

func (self *rlList) Prev() Element {
	if self.current == nil {
		return nil
	}
	self.current = self.current.Prev()
	return self.current
}

func (self *rlList) Current() Element {
	return self.current
}

func (self *rlList) Length() int {
	return self.length
}

func (self *rlList) LPush(f func()) {
	self.mod_locker.Lock()

	if self.current == nil {
		self.addFirstElement(f)
	} else {
		self.first.prev = &element{value: f, next: self.first}
		self.first = self.first.prev
	}

	self.mod_locker.Unlock()
}

func (self *rlList) RPush(f func()) {
	self.mod_locker.Lock()

	if self.current == nil {
		self.addFirstElement(f)
	} else {
		self.last.next = &element{value: f, prev: self.last}
		self.last = self.last.next
	}

	self.mod_locker.Unlock()
}

func (self *rlList) LPop() Element {
	self.mod_locker.Lock()

	to_return := self.first
	if self.current == self.first {
		self.current = self.first.next
	}
	self.first = self.first.next
	return to_return

	self.mod_locker.Unlock()
}

func (self *rlList) RPop() Element {
	self.mod_locker.Lock()

	to_return := self.last
	if self.current == self.last {
		self.current = self.last.prev
	}
	self.last = self.last.prev
	return to_return

	self.mod_locker.Unlock()
}

func (self *rlList) LPeek() Element {
	return self.first
}

func (self *rlList) RPeek() Element {
	return self.last5
}

func (self *rlList) addFirstElement(f func()) {
	self.current = &element{value: f}
	self.first = self.current
	self.last = self.current
}
