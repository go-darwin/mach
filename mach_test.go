package mach

import (
	"testing"

	"github.com/go-darwin/mach/internal/machtesting"
)

func TestTaskSelfTrap(t *testing.T) {
	t.Parallel()

	got := TaskSelfTrap()
	want := machtesting.TaskSelfTrap()

	if got != want {
		t.Fatalf("got %d but want %d", got, want)
	}
}

func TestThreadSelfTrap(t *testing.T) {
	t.Parallel()

	got := ThreadSelfTrap()
	want := machtesting.MachThreadSelf()

	if got != want {
		t.Fatalf("got %d but want %d", got, want)
	}
}

func TestHostSelf(t *testing.T) {
	t.Parallel()

	got := HostSelfTrap()
	want := machtesting.MachHostSelf()

	if got != want {
		t.Fatalf("got %d but want %d", got, want)
	}
}
