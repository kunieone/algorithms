package sort

import (
	"reflect"
	"testing"
)

func TestHSort(t *testing.T) {
	type args struct {
		array []int
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
			if got := HSort(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
