package tree

import (
	"testing"

	"github.com/gookit/goutil/dump"
)

func TestNew(t *testing.T) {
	dump.P(New(0, &[]int{1, 2, 3}))
}
