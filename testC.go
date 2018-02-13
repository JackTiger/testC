package testC

import (
	"github.com/JackTiger/testB"
	"fmt"
)

func OutPutC() {
	fmt.Println("OutPutC...")
	fmt.Println("Call OutPutB...")
	testB.OutPutB()
}
