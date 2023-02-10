package testreq

import (
	"testing"

	"github.com/pmezard/go-difflib/difflib"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// Assert if provided pb message matches expected message
func AssertPbEqual(t *testing.T, expected proto.Message, actual proto.Message) {
	expectedStr := protojson.Format(expected)
	actualStr := protojson.Format(actual)
	if expectedStr != actualStr {
		diff, err := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
			A:        difflib.SplitLines(expectedStr),
			B:        difflib.SplitLines(actualStr),
			FromFile: "Expected",
			FromDate: "",
			ToFile:   "Actual",
			ToDate:   "",
			Context:  13,
		})
		assert.NoError(t, err)

		t.Errorf("\nProtobuff Missmatch: \n"+
			"expected: %s\n"+
			"actual  : %s\n\n"+
			"\n\nDiff:\n%s", expected, actual, diff)
		t.FailNow()
	}
}
