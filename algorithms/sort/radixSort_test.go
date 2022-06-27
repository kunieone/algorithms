package sort

import (
	"fmt"
	"testing"
)

func TestRSort(t *testing.T) {
	fmt.Println(RSort([]int{12, 345, 67, 1, 9999, 2}, 2))
}

func Test_GetDigit(t *testing.T) {
	type args struct {
		num  int
		dig  int
		base int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "测试取一个数字的某个进制的某位数",
			args: args{
				num:  3764,
				dig:  0,
				base: 10,
			},
			want: 4,
		},
		{name: "测试取一个数字的某个进制的某位数",
			args: args{
				num:  3764,
				dig:  1,
				base: 10,
			},
			want: 6,
		},
		{name: "测试取一个数字的某个进制的某位数",
			args: args{
				num:  3764,
				dig:  2,
				base: 10,
			},
			want: 7,
		},
		{name: "测试取一个数字的某个进制的某位数",
			args: args{
				num:  3764,
				dig:  3,
				base: 10,
			},
			want: 3,
		},
		{name: "对于八进制来说",
			args: args{
				num:  10,
				dig:  0,
				base: 16,
			},
			want: 0xa,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDigit(tt.args.num, tt.args.dig, tt.args.base); got != tt.want {
				t.Errorf("getDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
