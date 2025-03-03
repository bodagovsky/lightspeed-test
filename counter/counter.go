package counter

import (
	"strconv"
	"strings"
)

const PartitionSize uint32 = 64

// StorageSize is counted by dividing 4294967296 (the maximum number of unique encoded ip addresses) by 64 (the size of each partition)
const StorageSize uint32 = 67_108_864

// bitmap is used to efficiently store encoded ip addresses
// each address is an index into one of 67_108_864 partitions for 0 - absent 1 - present
type bitmap struct {
	// storage is an array of uint64 integers each of which is a partition
	// as there is no integer large enough to keep 4294967296 elements in it
	storage [StorageSize]uint64

	// cardinality field stores the number of unique elements
	cardinality uint64
}

// Add inserts the ip address into storage and increments
// the cardinality if the value was not present before
func (b *bitmap) Add(val uint32) {
	// div is used to access the partition and mod is the bit we need to set
	div, mod := val/PartitionSize, val%PartitionSize
	if b.setbit(div, mod) {
		b.cardinality++
	}
}

// setbit sets n-th bit on i-th partition to 1
func (b *bitmap) setbit(i, n uint32) bool {
	isNew := b.storage[i]>>n&0b1 == 0
	b.storage[i] |= 1 << n
	return isNew
}

func (b *bitmap) GetCardinality() uint64 { return b.cardinality }

type IpAdressCounter struct {
	bitMap bitmap
}

func NewIpAdressCounter() IpAdressCounter {
	return IpAdressCounter{
		bitMap: bitmap{},
	}
}

func (counter *IpAdressCounter) Process(addr string) uint64 {
	addrParts := strings.Split(addr, ".")
	var ipAddr [4]uint8
	for i := 0; i < 4; i++ {
		part, err := strconv.Atoi(addrParts[i])
		if err != nil {
			panic(err)
		}
		ipAddr[i] = uint8(part)
	}
	encodedAddr := Encode(ipAddr)
	counter.bitMap.Add(encodedAddr)
	return counter.bitMap.GetCardinality()
}

// Encode represents each ipV4 address as uint32 integer
// as we know that it takes only 4 bytes to store each ip address
// Example:
//
//		                255.255.0.1
//	                   /     |  |  \
//	                  /      |  |   \
//	         0b11111111    ... 0b0   0b00000001
//
// which are then concatenated into single integer
func Encode(addr [4]uint8) uint32 {
	var result uint32
	for _, val := range addr {
		result = result << 8
		result = result | uint32(val)
	}
	return result
}
