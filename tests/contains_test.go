package main

import (
	"testing"
)

func Test_testNoChans(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "noChans500",
			args: args{
				num: 500,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testNoChans(tt.args.num); got != tt.want {
				t.Errorf("testNoChans() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_testChans(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "noChans500",
			args: args{
				num: 500,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testChans(tt.args.num); got != tt.want {
				t.Errorf("testChans() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_NoChans(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testNoChans(4000000)
	}
}

func Benchmark_Chans(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testChans(4000000)
	}
}
