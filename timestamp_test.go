package anonymize

import (
	"testing"
	"time"
)

func TestRoundTimestamp(t *testing.T) {
	tests := []struct {
		name string
		in   time.Time
		want time.Time
	}{
		{
			name: "zero time",
			in:   time.Time{},
			want: time.Time{},
		},
		{
			name: "with seconds",
			in:   time.Date(2026, 4, 10, 14, 23, 45, 0, time.UTC),
			want: time.Date(2026, 4, 10, 14, 23, 0, 0, time.UTC),
		},
		{
			name: "with nanoseconds",
			in:   time.Date(2026, 4, 10, 14, 23, 45, 123456789, time.UTC),
			want: time.Date(2026, 4, 10, 14, 23, 0, 0, time.UTC),
		},
		{
			name: "already on minute boundary",
			in:   time.Date(2026, 4, 10, 14, 23, 0, 0, time.UTC),
			want: time.Date(2026, 4, 10, 14, 23, 0, 0, time.UTC),
		},
		{
			name: "last second before next minute",
			in:   time.Date(2026, 4, 10, 14, 23, 59, 999999999, time.UTC),
			want: time.Date(2026, 4, 10, 14, 23, 0, 0, time.UTC),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := RoundTimestamp(tc.in)
			if !got.Equal(tc.want) {
				t.Errorf("RoundTimestamp(%v) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}
