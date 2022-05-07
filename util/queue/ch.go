package queue

// ChanQueue 顺序执行
type ChanQueue struct {
	taskChan chan func()
}

func NewChanQueue() *ChanQueue {
	return &ChanQueue{
		taskChan: make(chan func(), 16),
	}
}

func (q *ChanQueue) Start() {
	go func() {
		for {
			(<-q.taskChan)()
		}
	}()
}

func (q *ChanQueue) Exec(f func()) {
	q.taskChan <- f
}
