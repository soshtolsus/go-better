package WorkerPool

import "sync"

type workerPool struct {
	max_size uint
	waiting_processes DLList
	is_stopping bool
	busy_semaphore *sync.WaitGroup
}

type WorkerPool interface {
	Go(f func())
	Do(f func())
	Stop()
}

//get a fresh instance of a WorkerPool
func NewWorkerPool(size uint) WorkerPool {
	return &workerPool{
		max_size: size,
		waiting_processes: NewDLList(),
		busy_semaphore: new(sync.WaitGroup),
	}
}

//queue the function and return
func (self *workerPool) Go(f func()) {
	if !self.is_stopping {
		go self.do_the_needful(f)
	}
}

//wait until there is an open slot, start the function, and return
func (self *workerPool) Do(f func()) {
	if !self.is_stopping {
		self.do_the_needful(f)
	}
}

//wait until all queued and executing processes are done
func (self *workerPool) Stop() {
	self.is_stopping = true
	self.busy_semaphore.Wait()
}

func (self *workerPool) do_the_needful(f func()) {
	//store the intent to execute before actually executing to ensure
	// all processes are executed before  returns
	self.busy_semaphore.Add(1)

	//prevent others from acquiring a place in the pool until it is
	// determined that I get a place
	self.waiting_processes.Lock()

	if self.waiting_processes.Length() < self.max_size {
		self.start_needful(f)
	}

	self.waiting_processes.Unlock()
}

func (self *workerPool) start_needful(f func()) {
	go func(){
		f()
		self.stop_needful()
	}()
}

func (self *workerPool) stop_needful() {
	self.waiting_processes.Lock()
	self.busy_semaphore.Done()
	{
		to_do := self.waiting_processes.RPop()

		if to_do != nil {
			self.start_needful(to_do.Value())
		}
	}
	self.waiting_processes.Unlock()
}
