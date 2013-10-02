/*
This Source Code Form is subject to the terms of the Mozilla Public License, v.
2.0. If a copy of the MPL was not distributed with this file, You can obtain one
at http://mozilla.org/MPL/2.0/.

Copyright 2013 Carl Johnson IV
*/

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
