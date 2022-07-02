package sort

import (
	"reflect"
	"testing"
)

func TestISort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ISort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ISort() = %v, want %v", got, tt.want)
			}
		})
	}
}
