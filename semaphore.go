/*
	author: Michael Putong @2024
	-----------------------------------------------------------------
	This is a library that is implementing semaphore pattern
	that can be used to manage number of concurrent processes.
	-----------------------------------------------------------------
	This code is free to use, modify and distribute, although
	the author is not responsible for any damage happened in its use.
	-----------------------------------------------------------------
	visit the code repository in github.com/michaelwp/go-semaphore
*/

package go_semaphore

type Semaphore interface {
	Acquire()
	Release()
}

type semaphore struct {
	semC chan struct{}
}

func SemaphoreNew(maxConcurrency int) Semaphore {
	return &semaphore{
		semC: make(chan struct{}, maxConcurrency),
	}
}

func (s *semaphore) Acquire() {
	s.semC <- struct{}{}
}

func (s *semaphore) Release() {
	<-s.semC
}
