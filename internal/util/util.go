package util

import (
	"log"
)

func Min(a, b interface{}) interface{} {
	switch a := a.(type) {
	case int64:
		b := b.(int64)
		if a < b {
			return a
		}
		return b
	case int:
		b := b.(int)
		if a < b {
			return a
		}
		return b
	default:
		log.Fatalf("Min: unsupported type %T", a)
		return nil
	}
}
