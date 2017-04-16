package goq

import (
	types "go/types"
)

// TypeMatcher tests whether type is matched or not.
type TypeMatcher interface {
	// Match returns true when given type is matched.
	Match(t types.Type) bool
}
