package types

import "go/types"

type Type interface {
	Check(t types.Type) bool
}
