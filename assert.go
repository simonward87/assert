// package assert contains a selection of test helper functions, used for
// making type-safe assertions that will generate a descriptive test
// error if not satisfied

package assert

import (
	"reflect"
	"strings"
	"testing"
)

func Equal[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got: %v; want %v", got, want)
	}
}

// NOTE: While got and want are both of type any, using a generic like
// this requires them to both be of the SAME type to satisfy the
// constraint
func DeepEqual[T any](t *testing.T, got, want T) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v; want %v", got, want)
	}
}

func StringContains(t *testing.T, s, substr string) {
	t.Helper()

	if !strings.Contains(s, substr) {
		t.Errorf(
			"got: %q; want to contain: %q",
			s,
			substr,
		)
	}
}

func NilError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("got: %v; want: nil", err)
	}
}
