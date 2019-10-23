package counter

type Counter struct {
	value int
}

func New(v int) Counter {
	return Counter{v}
}

func (counter *Counter) Inc() {
	counter.value += 1
}

func (counter *Counter) Dec() {
	counter.value -= 1
}

func (counter *Counter) Get() int {
	return counter.value
}
