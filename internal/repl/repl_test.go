package repl

import (
	"testing"

	"kvstore/internal/store"
)

func TestWriteRead(t *testing.T) {
	s := store.New()

	Execute(s, "write 1 hello")

	out, _ := Execute(s, "read 1")
	if out != "hello" {
		t.Fatalf("expected hello, got %s", out)
	}
}

func TestDelete(t *testing.T) {
	s := store.New()

	Execute(s, "write 1 hello")
	Execute(s, "delete 1")

	out, _ := Execute(s, "read 1")
	if out == "hello" {
		t.Fatal("key should be deleted")
	}
}

func TestTransactionAbort(t *testing.T) {
	s := store.New()

	Execute(s, "write 1 hello")
	Execute(s, "start")
	Execute(s, "write 1 world")
	Execute(s, "abort")

	out, _ := Execute(s, "read 1")
	if out != "hello" {
		t.Fatalf("expected hello after abort, got %s", out)
	}
}

func TestTransactionCommit(t *testing.T) {
	s := store.New()

	Execute(s, "start")
	Execute(s, "write 1 world")
	Execute(s, "commit")

	out, _ := Execute(s, "read 1")
	if out != "world" {
		t.Fatalf("expected world, got %s", out)
	}
}

