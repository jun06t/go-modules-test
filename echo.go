package echo

import (
	"fmt"

	"github.com/jun06t/A"
	"github.com/jun06t/B"
)

func Call() {
	fmt.Println(A.Func())
	fmt.Println(B.Func())
}
