package pokecache

import (
	"testing"
	"time"
)

func TestAddAndGet(t *testing.T) {
	interval := 5 * time.Second
	cache := NewCache(interval)

	cases := []struct {
		key   string
		value []byte
	}{
		{
			"https://martingallauner.com", []byte("testdata"),
		},
		{
			"https://martingallauner.com", []byte("testdata"),
		},
	}

	for _, c := range cases {
		cache.Add(c.key, c.value)
		value, ok := cache.Get(c.key)

		if !ok {
			t.Errorf("Key expected but not found")
			return
		}
		if string(value) != string(c.value) {
			t.Errorf("Value expected but not found")
		}
	}

}

func TestReapLoop(t *testing.T) {
	interval := 1 * time.Second
	cache := NewCache(interval)

	key := "https://learntowalk.dev"
	value := []byte("testdate")

	cache.Add(key, value)

	time.Sleep(2 * time.Second)

	_, exists := cache.Get(key)
	if exists {
		t.Errorf("Entry must not exist in the cache anymore.")
		return
	}
}
