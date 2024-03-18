package eqStruct

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/prototext"
)

func TestStructs(t *testing.T) {
	// Find the paths of all input files in the data directory.
	paths, err := filepath.Glob(filepath.Join("testdata", "**", "*.input"))
	if err != nil {
		t.Fatal(err)
	}
	mo := prototext.MarshalOptions{
		Multiline: true,
		Indent:    "",
	}

	for _, path := range paths {
		folder, filename := filepath.Split(path)
		testname := filename[:len(filename)-len(filepath.Ext(filename))]

		// Each path turns into a test: the test name is the filename without the
		// extension.
		t.Run(testname, func(t *testing.T) {
			source, err := os.ReadFile(path)
			if err != nil {
				t.Fatal("error reading source file:", err)
			}

			out := make([]byte, len(source))
			l, err := base64.StdEncoding.Decode(out, source)
			if err != nil {
				t.Fatalf("failed to decode input: %s", err)
			}
			out = out[:l]

			typeString, _, _ := strings.Cut(testname, "-")
			eqTstr := fmt.Sprintf("EQT_%s", typeString)

			eqT, err := EQTypeString(eqTstr)
			if err != nil {
				t.Fatalf("failed to get eqType %s", eqTstr)
			}

			obj := reflect.New(eqT.TypeOf()).Interface()
			eqS, ok := obj.(EQStruct)
			if !ok {
				t.Fatalf("invalid type: %T", obj)
			}

			genericUnmarshal(t, eqS, out)

			got := mo.Format(eqS.ProtoMess())

			// Each input file is expected to have a "golden output" file, with the
			// same path except the .input extension is replaced by .golden
			goldenfile := filepath.Join(folder, testname+".golden")
			want, err := os.ReadFile(goldenfile)
			if err != nil {
				t.Fatal("error reading golden file:", err)
			}

			wantStr := string(want[:])
			wantStr = strings.Replace(wantStr, "\r", "", -1)
			if diff := cmp.Diff(wantStr, got); diff != "" {
				t.Errorf("test diff: %s", diff)
			}
		})
	}
}

func genericUnmarshal[S EQStruct](t *testing.T, s S, b []byte) (int, error) {
	t.Helper()

	return s.Unmarshal(b)
}
