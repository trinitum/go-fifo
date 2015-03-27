package fifo

type FIFO struct {
    items    []interface{}
    start    int
    length   int
}

func New(size int) *FIFO {
    return &FIFO{
        items: make([]interface{}, size),
    }
}

func (buf *FIFO) Push(item interface{}) {
    // TODO
}

func (buf *FIFO) Shift() interface{} {
    // TODO
}

func (buf *FIFO) Item(idx int) interface{} {
    // TODO
}
