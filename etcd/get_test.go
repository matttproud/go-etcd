package etcd

import (
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	Set("foo", "bar", 100)

	// wait for commit
	time.Sleep(100 * time.Millisecond)

	results, err := Get("foo")

	if err != nil || results[0].Key != "/foo" || results[0].Value != "bar" {
		if err != nil {
			t.Fatal(err)
		}
		t.Fatalf("Get failed with %s %s %v", results[0].Key, results[0].Value, results[0].TTL)
	}

	results, err = GetFrom("foo", "0.0.0.0:4001")

	if err != nil || results[0].Key != "/foo" || results[0].Value != "bar" {
		if err != nil {
			t.Fatal(err)
		}
		t.Fatalf("Get failed with %s %s %v", results[0].Key, results[0].Value, results[0].TTL)
	}

	results, err = GetFrom("foo", "0.0.0.0:4009")

	if err == nil {
		t.Fatal("should not get from port 4009")
	}
}
