package designmode

import (
	"fmt"
)

func OldCode() {
	fmt.Println("I'm old COde")
}

func NewCode() {
	fmt.Println("I'm new Code")
}

func Run() {
	_OldCode := OldCode

	OldCode := func() {
		_OldCode()
		NewCode()
	}
	OldCode()
}
