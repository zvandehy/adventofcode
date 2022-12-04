package day03

import "testing"

func TestPriority(t *testing.T) {
	type args struct {
		item string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a",
			args: args{
				item: "a",
			},
			want: 1,
		},
		{
			name: "A",
			args: args{
				item: "A",
			},
			want: 27,
		},
		{
			name: "z",
			args: args{
				item: "z",
			},
			want: 26,
		},
		{
			name: "Z",
			args: args{
				item: "Z",
			},
			want: 52,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Priority(tt.args.item); got != tt.want {
				t.Errorf("Priority(%s) = %v, want %v", tt.args.item, got, tt.want)
			}
			if got := PriorityFromCharCode(int(tt.args.item[0])); got != tt.want {
				t.Errorf("PriorityFromCharCode(%d) = %v, want %v", int(tt.args.item[0]), got, tt.want)
			}
		})
	}
}
