package rpq

import (
	"math/rand"
	"os"
	"testing"
)

var rpq *RPQ

func TestMain(m *testing.M) {
	rpq = NewRPQ()
	exitVal := m.Run()
	os.Exit(exitVal)
}

func checkTop(t *testing.T, key KeyType, val ValueType) {
	if k, v := rpq.Top(); k != key || v != val {
		t.Fatalf("Top key: %d should be %d\nTop value: %d should be %d", k, key, v, val)
	}
}
func checkSize(t *testing.T, val ValueType) {
	if size := rpq.Size(); size != val {
		t.Fatalf("Incorrrect size: %d should be %d", size, val)
	}
}
func TestBasic(t *testing.T) {
	for i := 100; i != 0; i-- {
		rpq.Push(i, i)
	}
	checkSize(t, 100)

	for i := 0; i < 5; i++ {
		checkTop(t, i+1, i+1)
		rpq.Pop()
	}
	checkSize(t, 95)
	checkTop(t, 6, 6)
	for i := 5; i != 0; i-- {
		rpq.Push(i, i)
	}
	checkTop(t, 1, 1)
	for !rpq.Empty() {
		rpq.Pop()
	}
	checkSize(t, 0)
	for _, n := range [4]int{2, 4, 6, 8} {
		rpq.Push(n, n)
	}
	for _, n := range [4]int{1, 3, 5, 7} {
		rpq.Push(n, n)
	}

	for i := 1; !rpq.Empty(); i++ {
		checkTop(t, i, i)
		rpq.Pop()
	}
	rpq.Clear()
	checkSize(t, 0)
}

func TestDecrease(t *testing.T) {
	for i := 1; i <= 50; i++ {
		rpq.Push(i, i)
	}
	for i := 50; i != 0; i-- {
		rpq.Decrease(i, -51+i)
		checkTop(t, i, -51+i)
	}
	for !rpq.Empty() {
		rpq.Pop()
	}
	for i := 0; i < 1000; i++ {
		val := rand.Intn(1000)
		rpq.Push(val, val)
	}

	rpq.Clear()
}
