package sort

import (
	"reflect"
	"testing"
)

func TestMSort(t *testing.T) {
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
			if got := MSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
