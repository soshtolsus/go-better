package WorkerPool

import "sync"

//TODO: replace available_goroutines and adder_lock with linked list of functions

type workerPool struct {
	available_goroutines chan interface{}
	is_waiting_to_stop bool
	waiting_lock *sync.WaitGroup
	adder_lock *sync.Mutex
}

type WorkerPool interface {
	Go(f func())
	Do(f func())
	Wait()
}

//get a fresh instance of a WorkerPool
func New(size uint) WorkerPool {
	to_return := new(workerPool)
	to_return.available_goroutines = make(chan interface{}, size)
	for i:=uint(0); i<size; i++ {
		to_return.available_goroutines <- nil
	}
	to_return.waiting_lock = new(sync.WaitGroup)
	to_return.adder_lock = new(sync.Mutex)
	return to_return
}

//queue the function and return
func (self *workerPool) Go(f func()) {
	if self.is_waiting_to_stop {
		panic("trying to Go when Wait has been called")
	}
	go self.do_the_needful(f)
}

//wait until there is an open slot, start the function, and return
func (self *workerPool) Do(f func()) {
	if self.is_waiting_to_stop {
		panic("trying to Do when Wait has been called")
	}
	self.do_the_needful(f)
}

//wait until all queued and executing processes are done
func (self *workerPool) Wait() {
	self.waiting_lock.Wait()
	close(self.available_goroutines)
}

func (self *workerPool) do_the_needful(f func()) {
	//store the intent to execute before actually executing to ensure
	// all processes are executed before Wait returns
	self.waiting_lock.Add(1)

	//prevent others from blocking on channel read, since all goroutines
	// that are blocked on reading a channel will simultaneously become
	// unblocked once an item is added to the channel
	//NOTE: not scalable - if A gets the lock and B and C block on getting
	// the lock, when A releases the lock, B and C will race, which defeats
	// the intent of a FIFO
	self.adder_lock.Lock()

	//wait for a slot to open up
	<-self.available_goroutines

	go func(){
		f()
		self.available_goroutines <- nil
		self.waiting_lock.Done()
	}()
	self.adder_lock.Unlock()
}
