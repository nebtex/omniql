package hash

import "github.com/spaolacci/murmur3"
const hashAddress = uint16(1)

func hashOf(s string) (h uint16) {

	h32 := murmur3.New32WithSeed(16)
	h32.Write([]byte(s))

	return uint16(65535 & h32.Sum32())

}

