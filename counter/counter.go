package counter

import (
	"strings"
	"strconv"
)

type IpAdressCounter struct {
	count int64
	ipStore [4][256]bool
}

func NewIpAdressCounter() IpAdressCounter {
	return IpAdressCounter{
		count:0,
		ipStore: [4][256]bool{},
	}
}

func (c * IpAdressCounter) Increment() {
	c.count++
}

func (c *IpAdressCounter) NumberUnique() int64 {
	return c.count
}

func (c *IpAdressCounter) CheckAndAdd (addr string) (bool, error) {
	present := true
	addrSplit := strings.Split(addr, ".")

	for i := 0; i < 4; i++ {
		part, err := strconv.Atoi(addrSplit[i])
		if err != nil {
			return false, err
		}
		present = present && c.ipStore[i][part]
		c.ipStore[i][part] = true
	}
	
	return present, nil
}