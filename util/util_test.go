package util

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestCalculateTax(t *testing.T) {
	t_1_want, _ := decimal.New(100, 0).Float64()
	t_1_want1, _ := decimal.New(1100, 0).Float64()

	t_2_want, _ := decimal.New(50, 0).Float64()
	t_2_want1, _ := decimal.New(2050, 0).Float64()

	t_3_want, _ := decimal.New(0, 0).Float64()
	t_3_want1, _ := decimal.New(100, 0).Float64()

	t_4_want, _ := decimal.NewFromFloat(0.01).Float64()
	t_4_want1, _ := decimal.NewFromFloat(101.01).Float64()

	type args struct {
		taxCode int64
		price   int64
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 float64
	}{
		{"1", args{1, 1000}, t_1_want, t_1_want1},
		{"2", args{2, 2000}, t_2_want, t_2_want1},
		{"3", args{3, 100}, t_3_want, t_3_want1},
		{"4", args{3, 101}, t_4_want, t_4_want1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CalculateTax(tt.args.taxCode, tt.args.price)
			if got != tt.want {
				t.Errorf("CalculateTax() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CalculateTax() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
