package code

import (
	"path/filepath"
	"testing"
)

type testCase struct {
	name    string
	path    string
	want    int64
	wantErr bool
}

const baseRoute = "./testdata/"

func TestGetSize(t *testing.T) {
	tests := []testCase{
		{
			name: "empty directory",
			path: "empty_directory",
			want: 0,
		},
		{
			name: "file",
			path: "test.json",
			want: 68,
		},
		{
			name: "directory with nested folders",
			path: "amazing_directory",
			want: 68,
		},
		{
			name:    "non-existent path",
			path:    "does_not_exist",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			path := filepath.Join(baseRoute, tc.path)
			got, err := GetSize(path)

			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}

				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
