package pokecache

import (
    "testing"
    "time"
    "bytes"
)

func TestAddGet(t *testing.T) {
    cache := NewCache(time.Minute)

    cases := []struct {
        key string
        val []byte
    }{
        {
            key: "key1",
            val: []byte("val1"),
        },
        {
            key: "key2",
            val: []byte("val2"),
        },
    }

    // Test adding and getting entries
    for _, c := range cases {
        cache.Add(c.key, c.val)
        val, ok := cache.Get(c.key)
        if !ok {
            t.Errorf("expected to find key %s in cache", c.key)
            continue
        }
        if !bytes.Equal(val, c.val) {
            t.Errorf("expected to find value %s in cache, got %s", c.val, val)
        }
    }

    // Test getting non-existent entry
    _, ok := cache.Get("non-existent")
    if ok {
        t.Error("expected to not find non-existent key in cache")
    }
}

func TestReap(t *testing.T) {
    interval := time.Millisecond * 100
    cache := NewCache(interval)

    // Add an entry
    cache.Add("key1", []byte("val1"))
    cache.Add("key2", []byte("val2"))
    cache.Add("key3", []byte("val3"))

    // Wait for slightly longer than the interval
    time.Sleep(interval * 2)

    // Verify entry was reaped
    for _, key := range []string{"key1", "key2", "key3"} {
	    _, ok := cache.Get(key)
	    if ok {
		t.Errorf("expected %s to be reaped from cache", key)
	    }
    }
}

func TestConcurrency(t *testing.T) {
    cache := NewCache(time.Minute)
    done := make(chan bool)

    // Start multiple goroutines to add/get entries
    for i := 0; i < 10; i++ {
        go func() {
            cache.Add("key", []byte("val"))
            cache.Get("key")
            done <- true
        }()
    }

    // Wait for all goroutines to complete
    for i := 0; i < 10; i++ {
        <-done
    }
    // If we get here without deadlock or panic, the test passes
}
