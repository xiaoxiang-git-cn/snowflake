package snowflake

import (
	"fmt"
	"testing"
)



func TestSnowFlake_Uuid(t *testing.T) {
	Init(1, 0)
	tmpMap := make(map[int64]struct{})
	for i := 0 ; i <= 10000000; i ++ {
		uid := Uuid()
		if _, ok := tmpMap[uid]; ok {
			fmt.Println("TestSnowFlake_Uuid err")
			return
		}
		tmpMap[uid] = struct{}{}
	}
}
