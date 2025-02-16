package main

type BucketNode struct {
	key  int
	val  int
	next *BucketNode
}

type MyMap struct {
	buckets []*BucketNode
	size    int
	used    int
}

func MyMapConstructor() *MyMap {
	var myMap = MyMap{
		buckets: make([]*BucketNode, 2),
		size:    2,
		used:    0,
	}
	return &myMap
}

func findKeyInChain(head *BucketNode, key int) *BucketNode {
	for head != nil && head.key != key {
		head = head.next
	}
	return head
}

func (mm *MyMap) Set(key, val int) {
	bucketIdx := key % mm.size

	ptr := findKeyInChain(mm.buckets[bucketIdx], key)
	if ptr != nil {
		ptr.val = val
	} else {
		mm.buckets[bucketIdx] = &BucketNode{
			key:  key,
			val:  val,
			next: mm.buckets[bucketIdx],
		}
	}

}

func (mm *MyMap) Get(key int) (int, bool) {
	bucketIdx := key % mm.size
	ptr := findKeyInChain(mm.buckets[bucketIdx], key)
	if ptr != nil {
		return ptr.val, true
	}
	return 0, false
}

func (mm *MyMap) Remove(key int) {
	bucketIdx := key % mm.size
	fwdPtr := findKeyInChain(mm.buckets[bucketIdx], key)
	if fwdPtr == nil {
		return
	}

	ptr := mm.buckets[bucketIdx]
	if ptr == fwdPtr {
		mm.buckets[bucketIdx] = fwdPtr.next
		fwdPtr.next = nil
		return
	}

	for ptr != nil && ptr.next != fwdPtr {
		ptr = ptr.next
	}
	ptr.next = fwdPtr.next
	fwdPtr.next = nil
}
