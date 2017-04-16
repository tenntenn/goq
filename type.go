package goq

import "go/types"

func toType(v interface{}) types.Type {
	switch v := v.(type) {
	case types.Type:
		return v
	case types.Object:
		return v.Type()
	}
	return nil
}
