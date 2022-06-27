package designmode

import (
	"fmt"
	"testing"
)

func TestSimpleFactory(t *testing.T) {
	p := NewPerson("zhangsan", 18)
	fmt.Println(p)
	p.Greet()
}
