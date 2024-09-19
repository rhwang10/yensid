package cache

import "fmt"

type Node struct {
	children map[string]*Node
	value    interface{}
}

type Cache struct {
	root *Node
}

func NewCache() *Cache {
	c := &Cache{}
	c.root = &Node{
		children: map[string]*Node{},
	}
	return c
}

func (c *Cache) Put(key string, value interface{}) {
	root := c.root

	// Find the shortest prefix that exists
	postfix, node := c.insertAt(key, root)

	// If the node is nil, there exists no node, so we create one at the root
	if node == nil {
		root.children[key] = &Node{
			value: value,
		}
	} else {
		if node.children[postfix] == nil {
			node.children[postfix] = &Node{
				value: value,
			}
		} else {
			node.children[postfix].value = value
		}
	}
}

func (c *Cache) insertAt(key string, root *Node) (string, *Node) {
searchLoop:
	for key != "" {
		if root.children == nil {
			break
		}

		for i := 0; i < len(key); i++ {
			p := key[:i]
			if n, ok := root.children[p]; ok {
				root = n
				key = key[i:]
			} else {
				break searchLoop
			}
		}
	}

	return key, root
}

func (c *Cache) Get(key string) (interface{}, bool) {
	n := c.search(key, c.root)
	if n == nil {
		return nil, false
	}

	return n.value, true
}

func (c *Cache) search(key string, root *Node) *Node {
	if len(key) == 0 && root.value != nil {
		return root
	}

	candidates := map[string][]*Node{}
	for i := 0; i <= len(key); i++ {
		prefix := key[:i]

		if root.children[prefix] == nil {
			continue
		}

		postfix := key[i:]
		candidates[postfix] = append(candidates[postfix], root.children[prefix])
	}

	if len(candidates) == 0 {
		return nil
	}

	for postfix, candidates := range candidates {
		for _, candidate := range candidates {
			res := c.search(postfix, candidate)
			if res != nil {
				return res
			}
		}
	}

	return nil
}

func cache() {
	cache := NewCache()
	cache.Put("a", "hello world")
	cache.Put("ab", "hi again")
	cache.Put("abcd", "hi one more time")
	cache.Put("abcd", "hihi")

	toFind := []string{"a", "ab", "abcd", "abc", "abcd"}
	for _, key := range toFind {
		value, found := cache.Get(key)
		if found {
			fmt.Println("Found value:", value)
		} else {
			fmt.Println("Value not found")
		}
	}
}
