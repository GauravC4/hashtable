package main

type BucketNode struct {
	key  int
	val  int
	next *BucketNode
}

type MyMap struct {
	buckets []*BucketNode
	used    int
}

func MyMapConstructor() *MyMap {
	var myMap = MyMap{
		buckets: make([]*BucketNode, 2),
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
	bucketIdx := key % len(mm.buckets)

	ptr := findKeyInChain(mm.buckets[bucketIdx], key)
	if ptr != nil {
		ptr.val = val
	} else {
		mm.buckets[bucketIdx] = &BucketNode{
			key:  key,
			val:  val,
			next: mm.buckets[bucketIdx],
		}
		mm.used++
		mm.Resize()
	}

}

func (mm *MyMap) Get(key int) (int, bool) {
	bucketIdx := key % len(mm.buckets)
	ptr := findKeyInChain(mm.buckets[bucketIdx], key)
	if ptr != nil {
		return ptr.val, true
	}
	return 0, false
}

func (mm *MyMap) Remove(key int) {
	bucketIdx := key % len(mm.buckets)
	fwdPtr := findKeyInChain(mm.buckets[bucketIdx], key)
	if fwdPtr == nil {
		return
	}

	ptr := mm.buckets[bucketIdx]
	if ptr == fwdPtr {
		mm.buckets[bucketIdx] = fwdPtr.next
		fwdPtr.next = nil
		mm.used--
		mm.Resize()
		return
	}

	for ptr != nil && ptr.next != fwdPtr {
		ptr = ptr.next
	}
	ptr.next = fwdPtr.next
	fwdPtr.next = nil
	mm.used--
	mm.Resize()
}

func (mm *MyMap) Size() int {
	return len(mm.buckets)
}

func (mm *MyMap) TransferBuckets(newLen int) {
	ogBuckets := mm.buckets
	mm.buckets = make([]*BucketNode, newLen)
	for _, ptr := range ogBuckets {
		for ptr != nil {
			bucketIdx := ptr.key % len(mm.buckets)
			ptr2 := ptr.next

			// transfer node to new bucket
			ptr.next = mm.buckets[bucketIdx]
			mm.buckets[bucketIdx] = ptr

			ptr = ptr2
		}
	}
}

func (mm *MyMap) Resize() {
	if mm.used < 2 {
		return
	}

	size := len(mm.buckets)

	//fmt.Printf("Resize func size : %d used : %d\n", size, mm.used)

	alpha := float32(mm.used) / float32(size)
	if alpha > 0.6 {
		//grow
		mm.TransferBuckets(size * 2)
	} else if alpha < 0.125 {
		//shrink
		mm.TransferBuckets(size / 2)
	}
}
