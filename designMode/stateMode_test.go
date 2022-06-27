package designmode

import (
	"reflect"
	"testing"
)

func TestStateMode(t *testing.T) {
	Run()
}

func TestNewDayContext(t *testing.T) {
	tests := []struct {
		name string
		want *DayContext
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDayContext(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDayContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDayContext_Today(t *testing.T) {
	tests := []struct {
		name string
		d    *DayContext
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.Today()
		})
	}
}

func TestDayContext_Next(t *testing.T) {
	tests := []struct {
		name string
		d    *DayContext
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.Next()
		})
	}
}

func TestMonday_Today(t *testing.T) {
	tests := []struct {
		name string
		m    *Monday
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Today()
		})
	}
}

func TestMonday_Next(t *testing.T) {
	type args struct {
		ctx *DayContext
	}
	tests := []struct {
		name string
		m    *Monday
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Next(tt.args.ctx)
		})
	}
}

func TestTuesday_Today(t *testing.T) {
	tests := []struct {
		name string
		tr   *Tuesday
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.Today()
		})
	}
}

func TestTuesday_Next(t *testing.T) {
	type args struct {
		ctx *DayContext
	}
	tests := []struct {
		name string
		tr   *Tuesday
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.Next(tt.args.ctx)
		})
	}
}

func TestWednesday_Today(t *testing.T) {
	tests := []struct {
		name string
		w    *Wednesday
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.Today()
		})
	}
}

func TestWednesday_Next(t *testing.T) {
	type args struct {
		ctx *DayContext
	}
	tests := []struct {
		name string
		w    *Wednesday
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.Next(tt.args.ctx)
		})
	}
}

func TestThursday_Today(t *testing.T) {
	tests := []struct {
		name string
		tr   *Thursday
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.Today()
		})
	}
}

func TestThursday_Next(t *testing.T) {
	type args struct {
		ctx *DayContext
	}
	tests := []struct {
		name string
		tr   *Thursday
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.Next(tt.args.ctx)
		})
	}
}

func TestFriday_Today(t *testing.T) {
	tests := []struct {
		name string
		f    *Friday
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f.Today()
		})
	}
}

func TestFriday_Next(t *testing.T) {
	type args struct {
		ctx *DayContext
	}
	tests := []struct {
		name string
		f    *Friday
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f.Next(tt.args.ctx)
		})
	}
}

func TestSaturday_Today(t *testing.T) {
	tests := []struct {
		name string
		s    *Saturday
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Today()
		})
	}
}

func TestSaturday_Next(t *testing.T) {
	type args struct {
		ctx *DayContext
	}
	tests := []struct {
		name string
		s    *Saturday
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Next(tt.args.ctx)
		})
	}
}

func TestSunday_Today(t *testing.T) {
	tests := []struct {
		name string
		s    *Sunday
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Today()
		})
	}
}

func TestSunday_Next(t *testing.T) {
	type args struct {
		ctx *DayContext
	}
	tests := []struct {
		name string
		s    *Sunday
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Next(tt.args.ctx)
		})
	}
}
