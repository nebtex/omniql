package hash

import (
	"testing"
	"fmt"
	"strconv"
)

func Test_hash(t *testing.T) {
	collision := map[uint16]string{}
	fmt.Println(hashOf("99"), hashOf("42"))

	for i := 0; i < 120000; i++ {
		r := hashOf(strconv.Itoa(i))
		id, ok := collision[r]

		if ok {
			fmt.Println(strconv.Itoa(i) ,id,  "222")

			break
		}
		collision[r] = strconv.Itoa(i)

	}
}
