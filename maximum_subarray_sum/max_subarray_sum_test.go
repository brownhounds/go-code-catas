package maximumsubarraysum

import "testing"

func TestMaxSubarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "mixed",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			name: "all_negative",
			nums: []int{-8, -3, -6, -2, -5, -4},
			want: -2,
		},
		{
			name: "all_positive",
			nums: []int{2, 3, 1, 5},
			want: 11,
		},
		{
			name: "single",
			nums: []int{7},
			want: 7,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := MaxSubarraySum(tc.nums)
			if got != tc.want {
				t.Fatalf("got %d, want %d for nums=%v", got, tc.want, tc.nums)
			}
		})
	}
}
