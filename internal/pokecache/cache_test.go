package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://exmaple,com?offset=20",
			val: []byte("mooooore"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)

			if !ok {
				t.Errorf("expected to find key: %s", c.key)
				return
			}

			if string(val) != string(c.val) {
				t.Errorf("expected to find value: %s", c.val)
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	const key = "https://example.com"

	cache := NewCache(baseTime)
	cache.Add(key, []byte("testdata"))

	_, ok := cache.Get(key)
	if !ok {
		t.Errorf("expected to find key: %s", key)
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get(key)
	if ok {
		t.Errorf("expected to not find key: %s", key)
		return
	}
}
