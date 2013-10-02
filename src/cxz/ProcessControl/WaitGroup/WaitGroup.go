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
