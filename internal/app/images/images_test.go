package images

import (
	"path/filepath"
	"testing"
)

func TestNumToImage(t *testing.T) {
	testCases := []struct {
		name    string
		num     int
		wantErr bool
	}{
		{
			name:    "good num",
			num:     1,
			wantErr: false,
		},
		{
			name:    "bad num",
			num:     -1,
			wantErr: true,
		},
	}

	p, _ := filepath.Abs("../../../images/")

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NumToImage(p, tc.num)
			if tc.wantErr != (err != nil) {
				t.Errorf("Unexpected error %v", err)
			}
		})
	}
}
