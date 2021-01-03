package rpq

import (
	"os"
	"testing"
)

var rpq *RPQ

func TestMain(m *testing.M) {
	rpq = NewRPQ()
	exitVal := m.Run()
	os.Exit(exitVal)
}

func checkTop(t *testing.T, val int) {
	if rpq.Top() != val {
		t.Fatalf("Top value: %d should be %d", rpq.Top(), val)
	}
}
func checkSize(t *testing.T, val int) {
	if size := rpq.Size(); size != val {
		t.Fatalf("Incorrrect size: %d should be %d", size, val)
	}
}
func TestBasic(t *testing.T) {
	for i := 100; i != 0; i-- {
		rpq.Push(i)
	}
	checkSize(t, 100)

	for i := 0; i < 5; i++ {
		checkTop(t, i+1)
		rpq.Pop()
	}
	checkSize(t, 95)
	checkTop(t, 6)
	for i := 5; i != 0; i-- {
		rpq.Push(i)
	}
	checkTop(t, 1)
	for !rpq.Empty() {
		rpq.Pop()
	}
	checkSize(t, 0)
	for _, n := range [4]int{2, 4, 6, 8} {
		rpq.Push(n)
	}
	for _, n := range [4]int{1, 3, 5, 7} {
		rpq.Push(n)
	}

	for i := 1; !rpq.Empty(); i++ {
		checkTop(t, i)
		rpq.Pop()
	}
	rpq.Clear()
	checkSize(t, 0)
}
