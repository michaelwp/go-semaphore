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
