package option

func Float32(f float32) *float32 {
	return &f
}

func Float64(f float64) *float64 {
	return &f
}

func Int(n int) *int {
	return &n
}

func String(s string) *string {
	return &s
}

func Bool(b bool) *bool {
	return &b
}

func Has(opt interface{}, v interface{}) bool {
	switch opt := opt.(type) {
	case *int:
		if v, ok := v.(int); ok && (opt == nil || *opt == v) {
			return true
		}
	case *int32:
		if v, ok := v.(int32); ok && (opt == nil || *opt == v) {
			return true
		}
	case *int64:
		if v, ok := v.(int64); ok && (opt == nil || *opt == v) {
			return true
		}
	case *float32:
		if v, ok := v.(float32); ok && (opt == nil || *opt == v) {
			return true
		}
	case *float64:
		if v, ok := v.(float64); ok && (opt == nil || *opt == v) {
			return true
		}
	case *bool:
		if v, ok := v.(bool); ok && (opt == nil || *opt == v) {
			return true
		}
	case *string:
		if v, ok := v.(string); ok && (opt == nil || *opt == v) {
			return true
		}
	case *Parttern:
		if v, ok := v.(string); ok && (opt == nil || opt.Match(v)) {
			return true
		}
	}

	return false
}
