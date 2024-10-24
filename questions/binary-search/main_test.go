package main

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		arr []int
		i   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "-1",
			args: args{
				arr: []int{1, 2, 3},
				i:   -1,
			},
			want: -1,
		},
		{
			name: "1",
			args: args{
				arr: []int{1, 2, 3},
				i:   1,
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				arr: []int{1, 2, 3},
				i:   3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.arr, tt.args.i); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
