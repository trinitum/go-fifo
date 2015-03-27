package fifo

import (
	"testing"
)

func checkLength(t *testing.T, expectedLen int, fifo *FIFO) {
	if fifo.Len() != expectedLen {
		t.Errorf("Expected length %d, but it is %d", expectedLen, fifo.Len())
	}
}

func checkShift(t *testing.T, expectedValue int, fifo *FIFO) {
	res := fifo.Shift()
	if res.(int) != expectedValue {
		t.Errorf("Expected shifted value to be %d, but got %d", expectedValue, res.(int))
	}
}

func TestInitialized(t *testing.T) {
	fifo := New(3)
	checkLength(t, 0, fifo)
	fifo.Push(1)
	fifo.Push(2)
	checkLength(t, 2, fifo)
	checkShift(t, 1, fifo)
	checkLength(t, 1, fifo)
	checkShift(t, 2, fifo)
	checkLength(t, 0, fifo)
}

func TestGrowing(t *testing.T) {
	fifo := New(3)
	fifo.Push(1)
	fifo.Push(2)
	fifo.Push(3)
	checkLength(t, 3, fifo)
	checkShift(t, 1, fifo)
	checkShift(t, 2, fifo)
	fifo.Push(4)
	t.Logf("after pushing 4: %v", fifo)
	if len(fifo.items) > 3 {
		t.Errorf("Pushing element caused storage to grow unexpectedly")
	}
	fifo.Push(5)
	t.Logf("after pushing 5: %v", fifo)
	fifo.Push(6)
	t.Logf("after pushing 6: %v", fifo)
	checkLength(t, 4, fifo)
	checkShift(t, 3, fifo)
	for i := 7; i <= 16; i++ {
		fifo.Push(i)
	}
	checkLength(t, 13, fifo)
	for i := 4; i <= 16; i++ {
		checkShift(t, i, fifo)
	}
	checkLength(t, 0, fifo)
}

func checkItem(t *testing.T, idx int, expect int, fifo *FIFO) {
	res := fifo.Item(idx)
	if res.(int) != expect {
		t.Errorf("Expected item %d to be %d, but got %d", idx, expect, res.(int))
	}
}

func TestUninitialized(t *testing.T) {
	fifo := new(FIFO)
	fifo.Push(9)
	fifo.Push(8)
	checkLength(t, 2, fifo)
	checkItem(t, 0, 9, fifo)
	checkItem(t, 1, 8, fifo)
}
