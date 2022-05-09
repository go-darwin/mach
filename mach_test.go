package mach_test

import (
	"testing"

	"github.com/go-darwin/mach"
	"github.com/go-darwin/mach/internal/machtesting"
)

func TestTaskSelfTrap(t *testing.T) {
	t.Parallel()

	got := mach.TaskSelfTrap()
	want := machtesting.TaskSelfTrap()

	if got != want {
		t.Fatalf("got %d but want %d", got, want)
	}
}

func TestThreadSelfTrap(t *testing.T) {
	t.Parallel()

	got := mach.ThreadSelfTrap()
	want := machtesting.MachThreadSelf()

	if got != want {
		t.Fatalf("got %d but want %d", got, want)
	}
}

func TestHostSelfTrap(t *testing.T) {
	t.Parallel()

	got := mach.HostSelfTrap()
	want := machtesting.MachHostSelf()

	if got != want {
		t.Fatalf("got %d but want %d", got, want)
	}
}
