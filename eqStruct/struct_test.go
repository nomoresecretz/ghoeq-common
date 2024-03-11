package eqStruct

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStructs(t *testing.T) {
	// Find the paths of all input files in the data directory.
	paths, err := filepath.Glob(filepath.Join("testdata", "*", "*.input"))
	if err != nil {
		t.Fatal(err)
	}

	for _, path := range paths {
		_, filename := filepath.Split(path)
		testname := filename[:len(filename)-len(filepath.Ext(path))]

		// Each path turns into a test: the test name is the filename without the
		// extension.
		t.Run(testname, func(t *testing.T) {
			source, err := os.ReadFile(path)
			if err != nil {
				t.Fatal("error reading source file:", err)
			}

			typeString, _, _ := strings.Cut(testname, "-")
			eqTstr := fmt.Sprintf("EQT_%s", typeString)

			eqT, err := EQTypeString(eqTstr)
			if err != nil {
				t.Fatalf("failed to get eqType %s", eqTstr)
			}

			obj := reflect.New(eqT.TypeOf()).Interface()
			got, ok := obj.(EQStruct)
			if !ok {
				t.Fatalf("invalid type: %T", obj)
			}

			genericUnmarshal(t, got, source)

			// Each input file is expected to have a "golden output" file, with the
			// same path except the .input extension is replaced by .golden
			goldenfile := filepath.Join("testdata", testname+".golden")
			want, err := os.ReadFile(goldenfile)
			if err != nil {
				t.Fatal("error reading golden file:", err)
			}

			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("test diff: %s", diff)
			}
		})
	}
}

func genericUnmarshal[S EQStruct](t *testing.T, s S, b []byte) error {
	t.Helper()

	return s.Unmarshal(b)
}
