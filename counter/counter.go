package counter

import (
	"strconv"
	"strings"
)

const StorageSize uint32 = 67_108_864

type bitmap struct {
	storage     [StorageSize]uint64
	cardinality uint64
}

func (b *bitmap) Add(val uint32) {
	div, mod := val/64, val%64
	if b.setbit(div, mod) {
		b.cardinality++
	}
}

func (b *bitmap) setbit(i, shift uint32) bool {
	isNew := b.storage[i]>>shift&0b1 == 0
	b.storage[i] |= 1 << shift
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

func Encode(addr [4]uint8) uint32 {
	var result uint32
	for _, val := range addr {
		result = result << 8
		result = result | uint32(val)
	}
	return result
}
