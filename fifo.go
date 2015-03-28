// Package fifo implements FIFO buffer. You can push items into buffer,
// shift items from it, or you can read arbitrary items in the buffer.
// Buffer automatically grows when required.
package fifo

type FIFO struct {
	items  []interface{}
	start  int
	length int
}

// New creates a new buffer and preallocates storage of the specified size.
func New(size int) *FIFO {
	return &FIFO{
		items: make([]interface{}, size),
	}
}

// Push adds a new item to the end of the buffer.
func (buf *FIFO) Push(item interface{}) {
	bufSize := len(buf.items)

	// if buffer is full, increase its size
	if buf.length == bufSize {
		var newSize int
		if bufSize > 0 {
			newSize = bufSize * 2
		} else {
			newSize = 1024
		}
		newItems := make([]interface{}, newSize)
		j := 0
		stop := min(buf.start+buf.length, bufSize)
		for i := buf.start; i < stop; i, j = i+1, j+1 {
			newItems[j] = buf.items[i]
		}
		if buf.start+buf.length > bufSize {
			for i := 0; i < buf.start+buf.length-bufSize; i, j = i+1, j+1 {
				newItems[j] = buf.items[i]
			}
		}
		buf.items = newItems
		buf.start = 0
		bufSize = newSize
	}

	next := (buf.start + buf.length) % bufSize
	buf.items[next] = item
	buf.length++
}

// Shift removes the first item from the buffer.
// It may return nil if the buffer is empty.
func (buf *FIFO) Shift() interface{} {
	if buf.length == 0 {
		return nil
	}
	item := buf.items[buf.start]
	buf.start = (buf.start + 1) % len(buf.items)
	buf.length--
	return item
}

// Item returns specified item from the buffer.
// If the index is out of range it returns nil.
func (buf *FIFO) Item(idx int) interface{} {
	if idx >= buf.length {
		return nil
	}
	return buf.items[(buf.start+idx)%len(buf.items)]
}

// Len returns the number of items in the buffer.
func (buf *FIFO) Len() int {
	return buf.length
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
