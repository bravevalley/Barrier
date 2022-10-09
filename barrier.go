package barrier

import "sync"

// Defining the barrier struct
type Barrier struct {
	// Total number of threads
	total int

	// Number of barriers consumed
	count int

	// Protection for the data above so no two threads
	// can read and write into it simultaneously - race condition
	mu *sync.Mutex

	// Logic for pausing all  threads
	condition *sync.Cond
}

// A function to create a barrier using OOP composition and 
// encapsulation
func NewBarrier(size int) *Barrier {

	locks := &sync.Mutex{}
	cond := sync.NewCond(locks)

	//  Return a new barrier can be worked on or with
	return &Barrier{
		total: size,
		count: size ,
		mu: locks,
		condition: cond,
	}

}


// Create a function for the barrier the details the functionality
// of a barrier.

// It works by
func (barrier *Barrier) Wait() {

	// 1. Lock access to the elements of the struct that would be 
	// updated
	barrier.mu.Lock()

	// This Wait function is called from within multiple threads
	// reduce the value fo the count by 1 for each thread that calls 
	// this barrier.Wait method
	barrier.count--

	// Test if we have exhausted the total number of threads predefined for
	// the barrier
	if barrier.count == 0 {

		// if we have exhausted then we reset the value to the original 
		// value of the number of threads so we can use it again if needed
		barrier.count = barrier.total

		// Broadcast to all the thread waiting for the depletion of the 
		// count variable so that continue execution
		barrier.condition.Broadcast()
	} else {

		// Now if there are still some threads that are yet to reach or call
		// this wait function, this is where we wait for them
		barrier.condition.Wait()
	}

	// Unlock the memory so other threads can write to it
	barrier.mu.Unlock()

}