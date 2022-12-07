package day06

import "testing"

func TestStartOfPacketMarker(t *testing.T) {
	type args struct {
		buffer string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				buffer: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			},
			want: 7,
		},
		{
			name: "test2",
			args: args{
				buffer: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			},
			want: 5,
		},
		{
			name: "test3",
			args: args{
				buffer: "nppdvjthqldpwncqszvftbrmjlhg",
			},
			want: 6,
		},
		{
			name: "test4",
			args: args{
				buffer: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			},
			want: 10,
		},
		{
			name: "test5",
			args: args{
				buffer: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StartOfPacketMarker(tt.args.buffer); got != tt.want {
				t.Errorf("StartOfPacketMarker() = %v, want %v", got, tt.want)
			}
		})
	}
}
