package goq

import "go/types"

func size(t *types.Basic) int {
	switch t.Kind() {
	case types.Int8, types.Uint8:
		return 8
	case types.Int16, types.Uint16:
		return 16
	case types.Int32, types.Uint32, types.Float32:
		return 32
	case types.Int64, types.Uint64, types.Float64, types.Complex64:
		return 64
	case types.Complex128:
		return 128
	}
	return 0
}
