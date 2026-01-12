package code

import (
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

func TestGetPathSize(t *testing.T) {
	tests := []testGetSizeCase{
		{
			name: "empty directory",
			path: "empty_directory",
			want: "0B",
		},
		{
			name: "file",
			path: "test.json",
			want: "68B",
		},
		{
			name: "directory with nested folders",
			path: "amazing_directory",
			want: "68B",
		},
		{
			name:    "non-existent path",
			path:    "does_not_exist",
			wantErr: true,
		},
		{
			name:          "check hidden files",
			path:          "empty_directory",
			includeHidden: true,
			want:          "18B",
		},
		{
			name: "does not calculate hidden file",
			path: "empty_directory/.hidden_json",
			want: "0B",
		},
		{
			name:          "file in hidden dir",
			path:          ".hidden_dir/test.json",
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
			name: "does not calculate recursive dir",
			path: "recursive_dir",
			want: "0B",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			path := getTestDataPath(tc.path)

			got, err := GetPathSize(path, tc.includeHidden, tc.recursive, tc.humanize)

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
		{name: "Bytes", size: 1, want: "1.0B"},
		{name: "KyloBytes", size: KB, want: "1.0KB"},
		{name: "MegaBytes", size: MB, want: "1.0MB"},
		{name: "GigaBytes", size: GB, want: "1.0GB"},
		{name: "TeraBytes", size: TB, want: "1.0TB"},
		{name: "PetaBytes", size: PB, want: "1.0PB"},
		{name: "ExaBytes", size: EB, want: "1.0EB"},
	}

	for _, test := range tests {
		tc := test

		t.Run(tc.name, func(t *testing.T) {
			got := FormatSize(tc.size)
			if tc.want != got {
				t.Errorf("got %s, want %s", got, tc.want)
			}
		})
	}
}

func TestBuildOutput(t *testing.T) {
	tests := []testBuildSizeCase{
		{name: "humanize", humanize: true, want: "1.0MB"},
		{name: "not humanize", humanize: false, want: "1000000B"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := BuildOutput(1_000_000, "test", tc.humanize)

			if got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}
