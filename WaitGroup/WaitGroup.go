/*
This Source Code Form is subject to the terms of the Mozilla Public License, v.
2.0. If a copy of the MPL was not distributed with this file, You can obtain one
at http://mozilla.org/MPL/2.0/.

Copyright 2013 Carl Johnson IV
*/

package WaitGroup

import "sync"

type waitGroup struct {
	is_stopping bool
	busy_semaphore *sync.WaitGroup
}

type WaitGroup interface {
	Go(f func())
	Stop()
}

//get a fresh instance of a WaitGroup
func NewWaitGroup() WaitGroup {
	return &waitGroup{
		busy_semaphore: new(sync.WaitGroup),
	}
}

//queue the function and return
func (self *waitGroup) Go(f func()) {
	if !self.is_stopping {
		go func(){
			self.busy_semaphore.Add(1)
			f()
			self.busy_semaphore.Done()
		}()
	}
}

//wait until all queued and executing processes are done
func (self *waitGroup) Stop() {
	self.is_stopping = true
	self.busy_semaphore.Wait()
}
