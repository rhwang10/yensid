# Yensid
A cache that uses a trie-based structure to store multiple values with overlapping key ranges under a single key. New keys are inserted at the shortest possible overlapping interval, even if an even shorter possible overlap exists.

Ex.
```go
c := NewCache()

c.Put("abcd", "hi")
c.Put("abd", "coffee")
c.Get("abcd") # hi
c.Get("abd") # coffee

c.Put("acd", "hi again")
c.Get("acd") # hi again
```

We could index all of these keys under a single key `a`, but because `abcd` and `abd` were created before `acd`, we have allocated one node for `ab` and one node for `acd`.