package game

import "testing"

func Test_longestStraight(t *testing.T) {
	type args struct {
		dice []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				dice: []int{1, 2, 3, 4, 5},
			},
			want: 5,
		},
		{
			name: "2",
			args: args{
				dice: []int{1, 2, 4, 5, 6},
			},
			want: 3,
		},
		{
			name: "3",
			args: args{
				dice: []int{1, 6, 2, 5, 4},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestStraight(tt.args.dice); got != tt.want {
				t.Errorf("longestStraight() = %v, want %v", got, tt.want)
			}
		})
	}
}
