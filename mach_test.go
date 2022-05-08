package mach

import (
	"testing"

	"github.com/go-darwin/mach/internal/machtesting"
)

func TestTaskSelfTrap(t *testing.T) {
	got := TaskSelfTrap()
	want := machtesting.TaskSelfTrap()

	if got != want {
		t.Fatalf("got %d but want %d", got, want)
	}
}

func TestThreadSelfTrap(t *testing.T) {
	got := ThreadSelfTrap()
	want := machtesting.MachThreadSelf()

	if got != want {
		t.Fatalf("got %d but want %d", got, want)
	}
}
