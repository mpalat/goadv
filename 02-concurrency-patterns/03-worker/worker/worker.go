package worker

type worker struct {
	size int
}

func New(sz int) *worker {
	return &worker{
		size: sz,
	}
}
