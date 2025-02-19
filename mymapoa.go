package main

import (
	"log"
)

type BucketPair struct {
	key      int
	val      int
	isActive bool
}

type MyMapOA struct {
	buckets []*BucketPair
	used    int
	active  int
}

// uses open addressing linear probing for conflict resolution
func MyMapOAConstructor() *MyMapOA {
	var myMap = MyMapOA{
		buckets: make([]*BucketPair, 4),
		used:    0,
		active:  0,
	}
	return &myMap
}

func isBucketEmpty(bucketPairPtr *BucketPair) bool {
	return bucketPairPtr == nil || !bucketPairPtr.isActive
}

func (mm *MyMapOA) Set(key, val int) {
	bucketIdx := key % len(mm.buckets)

	i := bucketIdx
	for {
		if mm.buckets[i] == nil {
			mm.buckets[i] = &BucketPair{
				key:      key,
				val:      val,
				isActive: true,
			}
			mm.used += 1
			mm.active += 1
			mm.Resize()
			break
		} else if mm.buckets[i].isActive && mm.buckets[i].key == key {
			mm.buckets[i].val = val
			break
		}
		i = (i + 1) % len(mm.buckets)
		if i == bucketIdx {
			// this should have been an return error but i dont want to make changes to interface and chaining as well
			// this is anyways a corner case never to be encourtered if growth and shrink happens correctly
			log.Fatal("Map full ! no empty bucket found for insertion")
		}
	}

}

func (mm *MyMapOA) Get(key int) (int, bool) {
	bucketIdx := key % len(mm.buckets)
	i := bucketIdx
	for {
		if !isBucketEmpty(mm.buckets[i]) && mm.buckets[i].key == key {
			return mm.buckets[i].val, true
		} else if mm.buckets[i] == nil {
			return 0, false
		}
		i = (i + 1) % len(mm.buckets)
		if i == bucketIdx {
			break
		}
	}
	return 0, false
}

func (mm *MyMapOA) Remove(key int) {
	bucketIdx := key % len(mm.buckets)
	i := bucketIdx
	for {
		if !isBucketEmpty(mm.buckets[i]) && mm.buckets[i].key == key {
			mm.buckets[i].isActive = false
			mm.active -= 1
			mm.Resize()
			break
		}
		i = (i + 1) % len(mm.buckets)
		if i == bucketIdx {
			break
		}
	}
}

func (mm *MyMapOA) Size() int {
	return len(mm.buckets)
}

func (mm *MyMapOA) TransferBuckets(newLen int) {
	ogBuckets := mm.buckets
	mm.buckets = make([]*BucketPair, newLen)
	for _, pair := range ogBuckets {
		if pair != nil && pair.isActive {
			bucketIdx := pair.key % len(mm.buckets)
			i := bucketIdx
			for {
				if mm.buckets[i] == nil {
					mm.buckets[i] = pair
					break
				}
				i = (i + 1) % len(mm.buckets)
				if i == bucketIdx {
					break
				}
			}
		}
	}
	mm.used = mm.active
}

func (mm *MyMapOA) Resize() {
	if mm.used < 2 {
		return
	}

	size := len(mm.buckets)

	//fmt.Printf("Resize func size : %d, used : %d, active : %d\n", size, mm.used, mm.active)

	alpha := float32(mm.used) / float32(size)
	if alpha > 0.6 {
		//grow
		mm.TransferBuckets(size * 2)
	} else if alpha < 0.125 {
		//shrink
		mm.TransferBuckets(size / 2)
	}

	beta := float32(mm.active) / float32(mm.used)
	if beta < 0.5 {
		//if lot of tomb stones, create new table with same size
		mm.TransferBuckets(size)
	}
}
