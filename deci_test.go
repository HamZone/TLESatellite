package satellite

import (
	"testing"
)

func TestDecimalAdd2(t *testing.T) {
	type args struct {
		d1 float64
		d2 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "d1 + d2",
			args: args{
				d1: 1.1115,
				d2: 1.1115,
			},
			want: 2.223,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecimalAdd2(tt.args.d1, tt.args.d2); got != tt.want {
				t.Errorf("DecimalAdd2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecimalMul3(t *testing.T) {
	type args struct {
		d1 float64
		d2 float64
		d3 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "d1 * d2 * d3",
			args: args{
				d1: 1.1,
				d2: 1.2,
				d3: 1.3,
			},
			want: 1.716,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecimalMul3(tt.args.d1, tt.args.d2, tt.args.d3); got != tt.want {
				t.Errorf("DecimalMul3() = %v, want %v", got, tt.want)
			}
		})
	}
}
