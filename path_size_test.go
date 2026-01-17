package code

import (
	"math"
	"path/filepath"
	"testing"
)

type testGetSizeCase struct {
	name          string
	path          string
	includeHidden bool
	recursive     bool
	humanize      bool
	want          string
	wantErr       bool
}

type testFormatSizeCase struct {
	name string
	size int64
	want string
}

type testBuildSizeCase struct {
	name     string
	humanize bool
	want     string
}

func getTestDataPath(file string) string {
	return filepath.Join("testdata", file)
}

func getPow(x float64, n float64) int64 {
	return int64(math.Pow(x, n))
}

func TestGetPathSize(t *testing.T) {
	tests := []testGetSizeCase{
		{
			name: "empty directory",
			path: "shadow_dir",
			want: "0B",
		},
		{
			name: "small file",
			path: "test.json",
			want: "68B",
		},
		{
			name:     "directory with nested folders and humanize",
			path:     "amazing_dir",
			humanize: true,
			want:     "5.4KB",
		},
		{
			name:    "non-existent path",
			path:    "does_not_exist",
			wantErr: true,
		},
		{
			name:          "check hidden files",
			path:          "shadow_dir",
			includeHidden: true,
			want:          "18B",
		},
		{
			name: "does not calculate hidden file",
			path: "shadow_dir/.hidden_json",
			want: "0B",
		},
		{
			name:          "file in hidden dir",
			path:          "shadow_dir/.hidden_dir/test.json",
			includeHidden: true,
			want:          "72B",
		},
		{
			name: "does not calculate file in hidden dir",
			path: ".hidden_dir/test.json",
			want: "0B",
		},
		{
			name:      "recursive",
			path:      "recursive_dir",
			recursive: true,
			want:      "102B",
		},
		{
			name: "non-recursive directory ignores nested files",
			path: "recursive_dir",
			want: "0B",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			path := getTestDataPath(tc.path)

			got, err := GetPathSize(path, tc.recursive, tc.humanize, tc.includeHidden)

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
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestFormatSize(t *testing.T) {
	tests := []testFormatSizeCase{
		{name: "Bytes", size: 1, want: "1B"},
		{name: "KyloBytes", size: 1024, want: "1.0KB"},
		{name: "MegaBytes", size: getPow(1024, 2), want: "1.0MB"},
		{name: "GigaBytes", size: getPow(1024, 3), want: "1.0GB"},
		{name: "TeraBytes", size: getPow(1024, 4), want: "1.0TB"},
		{name: "PetaBytes", size: getPow(1024, 5), want: "1.0PB"},
		{name: "ExaBytes", size: getPow(1024, 6), want: "1.0EB"},
	}

	for _, test := range tests {
		tc := test

		t.Run(tc.name, func(t *testing.T) {
			got := formatSize(tc.size)
			if tc.want != got {
				t.Errorf("got %s, want %s", got, tc.want)
			}
		})
	}
}

func TestBuildOutput(t *testing.T) {
	tests := []testBuildSizeCase{
		{name: "humanize", humanize: true, want: "976.6KB"},
		{name: "not humanize", humanize: false, want: "1000000B"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := buildOutput(1_000_000, tc.humanize)

			if got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}
