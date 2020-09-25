package datastructures

// GenericList is a generic list datastructure
type GenericList struct {
	Head *node
}

type listData struct {
	key   string
	value interface{}
}

type node struct {
	Next *node
	Data listData
}

// IsEmpty returns true if GenericList is empty
func (ll *GenericList) IsEmpty() bool {
	return ll.Head == nil
}

// Size returns the number of items in the list
func (ll *GenericList) Size() int {

	curNode := ll.Head
	count := 0

	for curNode != nil {
		count++
		curNode = curNode.Next
	}

	return count
}

// Append adds an key/value to the end of the list
func (ll *GenericList) Append(key string, value interface{}) bool {
	nextNode := &node{
		Next: ll.Head,
		Data: listData{key, value},
	}

	ll.Head = nextNode

	return true
}

// Get returns the value of key if present
func (ll *GenericList) Get(key string) interface{} {
	cur := ll.Head

	for cur != nil {
		if cur.Data.key == key {
			return cur.Data.value
		}
		cur = cur.Next
	}

	return nil
}

// Remove a key/value pair by key
func (ll *GenericList) Remove(key string) bool {
	cur := ll.Head

	if cur == nil {
		return false
	}

	if cur.Data.key == key {
		ll.Head = cur.Next
		return true
	}

	for cur.Next != nil {
		if cur.Next.Data.key == key {
			cur.Next = cur.Next.Next
			return true
		}
		cur = cur.Next
	}

	return false
}
