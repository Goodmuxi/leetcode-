package test

import (
	"reflect"
	"src/algorithm/topInterview/top1"
	"testing"
)

func TestTwoSumAll(t *testing.T) {
	type args struct {
		num    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []top1.Result
	}{
		// TODO: Add test cases.
		{
			name: "leetcode2",
			args: args{
				num:    []int{1, 2, 2, 7, 8, 11, 15},
				target: 9},
			want: []top1.Result{{2, 3}, {0, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := top1.TwoSumAll(tt.args.num, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwoSum(t *testing.T) {
	type args struct {
		num    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "leetcode1",
			args: args{
				num:    []int{2, 7, 11, 15},
				target: 9},
			want: []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := top1.TwoSum(tt.args.num, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
