package cache

import "testing"

func TestCache(t *testing.T) {
	cache := NewCache()

	cache.Put("foo", "bar")
	a, found := cache.Get("foo")

	if !found {
		t.Error("expected to find foo")
	}
	if a != "bar" {
		t.Errorf("expected foo to be bar, got %s", a)
	}

	b, found := cache.Get("bar")
	if found {
		t.Error("expected not to find bar")
	}
	if b != nil {
		t.Errorf("expected bar to be nil, got %s", b)
	}
}
