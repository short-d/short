package cache

import (
	"strconv"
	"testing"

	"github.com/short-d/app/mdtest"
)

func TestCache_Get(t *testing.T) {
	testCases := []struct {
		name        string
		cache       Cache
		key         interface{}
		expectedRes []interface{}
		expectedSeq string
	}{
		{
			name:        "success",
			cache:       buildCache(10, 10),
			key:         "5",
			expectedRes: []interface{}{"5"},
			expectedSeq: "5->0->1->2->3->4->6->7->8->9",
		},
		{
			name:        "not found",
			cache:       buildCache(11, 11),
			key:         "12",
			expectedRes: []interface{}{},
			expectedSeq: "0->1->2->3->4->5->6->7->8->9->10",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.cache.Get(testCase.key)

			mdtest.Equal(t, testCase.expectedRes, res)
			mdtest.Equal(t, testCase.expectedSeq, nodeToString(testCase.cache.Head))
		})
	}
}

func TestCache_Set(t *testing.T) {
	testCases := []struct {
		name        string
		cache       Cache
		key         interface{}
		value       []interface{}
		expectedVal []interface{}
		expectedSeq string
	}{
		{
			name:        "cache does not reach capacity and key does not exist",
			cache:       buildCache(10, 8),
			key:         "8",
			value:       []interface{}{"k"},
			expectedVal: []interface{}{"k"},
			expectedSeq: "8->0->1->2->3->4->5->6->7",
		},
		{
			name:        "cache does not reach capacity and key exists",
			cache:       buildCache(10, 8),
			key:         "3",
			value:       []interface{}{"d"},
			expectedVal: []interface{}{"d"},
			expectedSeq: "3->0->1->2->4->5->6->7",
		},
		{
			name:        "cache does reach capacity and key does not exist",
			cache:       buildCache(11, 11),
			key:         "11",
			value:       []interface{}{"11"},
			expectedVal: []interface{}{"11"},
			expectedSeq: "11->0->1->2->3->4->5->6->7->8->9",
		},
		{
			name:        "cache does not reach capacity and key does not exists",
			cache:       buildCache(10, 8),
			key:         "22",
			value:       []interface{}{"22"},
			expectedVal: []interface{}{"22"},
			expectedSeq: "22->0->1->2->3->4->5->6->7",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.cache.Set(testCase.key, testCase.value)

			mdtest.Equal(t, testCase.expectedVal, testCase.cache.Mapping[testCase.key].Val)
			mdtest.Equal(t, testCase.expectedSeq, nodeToString(testCase.cache.Head))
		})
	}
}

func TestCache_Add(t *testing.T) {
	testCases := []struct {
		name        string
		cache       Cache
		node        *Node
		expectedRes string
	}{
		{
			name:  "empty linkedlist",
			cache: NewCache(10),
			node: &Node{
				Key: "1",
			},
			expectedRes: "1",
		},
		{
			name: "non-empty linkedlist",
			cache: func() Cache {
				cache := NewCache(10)
				node1 := Node{
					Key: "1",
				}
				node2 := Node{
					Key: "2",
				}
				node1.Next = &node2
				node2.Prev = &node1
				cache.Head = &node1
				cache.Tail = &node2
				return cache
			}(),
			node: &Node{
				Key: "3",
			},
			expectedRes: "3->1->2",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.cache.Add(testCase.node)

			mdtest.Equal(t, testCase.expectedRes, nodeToString(testCase.cache.Head))
		})
	}
}

func TestCache_Remove(t *testing.T) {
	cache1 := buildCache(1, 1)
	cache2 := buildCache(3, 3)
	cache3 := buildCache(3, 3)
	cache4 := buildCache(3, 3)

	testCases := []struct {
		name        string
		cache       Cache
		node        *Node
		expectedRes string
	}{
		{
			name:        "single node",
			cache:       cache1,
			node:        cache1.Head,
			expectedRes: "nil",
		},
		{
			name:        "multiple nodes: remove Head",
			cache:       cache2,
			node:        cache2.Head,
			expectedRes: "1->2",
		},
		{
			name:        "multiple nodes: remove Tail",
			cache:       cache3,
			node:        cache3.Tail,
			expectedRes: "0->1",
		},
		{
			name:        "multiple nodes: remove middle",
			cache:       cache4,
			node:        cache4.Head.Next,
			expectedRes: "0->2",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.cache.Remove(testCase.node)

			mdtest.Equal(t, testCase.expectedRes, nodeToString(testCase.cache.Head))
		})
	}
}

func TestNewCache(t *testing.T) {
	cache := NewCache(1)
	mdtest.NotEqual(t, nil, cache)
	mdtest.Equal(t, 1, cache.Capacity)
}

// nodeToString converts node key to its string representation
func nodeToString(node *Node) string {
	if node == nil {
		return "nil"
	}

	var nodeStr string
	for node != nil {
		nodeStr = nodeStr + node.Key.(string) + "->"
		node = node.Next
	}
	return nodeStr[:len(nodeStr)-2]
}

// buildCache returns a Cache with a doubly linked list of key 0->1->2->3->...
func buildCache(capacity int, size int) Cache {
	cache := NewCache(capacity)
	nodeList := make([]Node, size)
	for i := 0; i < len(nodeList); i++ {
		key := strconv.Itoa(i)
		val := []interface{}{strconv.Itoa(i)}
		nodeList[i] = Node{
			Key: key,
			Val: val,
		}
		cache.Mapping[key] = &nodeList[i]
	}
	for i := 1; i < len(nodeList)-1; i++ {
		nodeList[i].Prev = &nodeList[i-1]
		nodeList[i].Next = &nodeList[i+1]
	}
	if size > 1 {
		nodeList[0].Next = &nodeList[1]
		nodeList[len(nodeList)-1].Prev = &nodeList[len(nodeList)-2]
	}

	cache.Head = &nodeList[0]
	cache.Tail = &nodeList[len(nodeList)-1]
	return cache
}
