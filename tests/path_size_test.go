package tests

import (
	"code"
	"path/filepath"
	"runtime"
	"testing"
)

type testGetSizeCase struct {
	name          string
	path          string
	includeHidden bool
	want          int64
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
	_, currentFile, _, ok := runtime.Caller(0)

	if !ok {
		panic("can't get current file path")
	}

	projectRoot := filepath.Dir(filepath.Dir(currentFile))

	return filepath.Join(projectRoot, "testdata", file)
}

func TestGetSize(t *testing.T) {
	tests := []testGetSizeCase{
		{
			name:          "empty directory",
			path:          "empty_directory",
			includeHidden: false,
			want:          0,
		},
		{
			name:          "file",
			path:          "test.json",
			includeHidden: false,
			want:          68,
		},
		{
			name:          "directory with nested folders",
			path:          "amazing_directory",
			includeHidden: false,
			want:          68,
		},
		{
			name:          "non-existent path",
			path:          "does_not_exist",
			includeHidden: false,
			wantErr:       true,
		},
		{
			name:          "check hidden files",
			path:          "empty_directory",
			includeHidden: true,
			want:          18,
		},
		{
			name:          "does not calculate hidden file",
			path:          "empty_directory/.hidden_json",
			includeHidden: false,
			want:          0,
		},
		{
			name:          "file in hidden dir",
			path:          ".hidden_dir/test.json",
			includeHidden: true,
			want:          72,
		},
		{
			name:          "does not calculate file in hidden dir",
			path:          ".hidden_dir/test.json",
			includeHidden: false,
			want:          0,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			path := getTestDataPath(tc.path)
			got, err := code.GetSize(path, tc.includeHidden)

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

func TestFormatSize(t *testing.T) {
	tests := []testFormatSizeCase{
		{name: "Bytes", size: 1, want: "1.0B"},
		{name: "KyloBytes", size: code.KB, want: "1.0KB"},
		{name: "MegaBytes", size: code.MB, want: "1.0MB"},
		{name: "GigaBytes", size: code.GB, want: "1.0GB"},
		{name: "TeraBytes", size: code.TB, want: "1.0TB"},
		{name: "PetaBytes", size: code.PB, want: "1.0PB"},
		{name: "ExaBytes", size: code.EB, want: "1.0EB"},
	}

	for _, test := range tests {
		tc := test

		t.Run(tc.name, func(t *testing.T) {
			got := code.FormatSize(tc.size)
			if tc.want != got {
				t.Errorf("got %s, want %s", got, tc.want)
			}
		})
	}
}

func TestBuildSize(t *testing.T) {
	var testSize int64 = 1_000_000
	testPath := "test"
	testHumanize := []testBuildSizeCase{
		{name: "humanize", humanize: true, want: "1.0MB\ttest\n"},
		{name: "not humanize", humanize: false, want: "1000000B\ttest\n"},
	}

	for _, test := range testHumanize {
		tc := test

		t.Run(tc.name, func(t *testing.T) {
			got := code.BuildOutput(testSize, testPath, tc.humanize)

			if got != tc.want {
				t.Errorf("got %s, want %s", got, tc.want)
			}
		})
	}
}
