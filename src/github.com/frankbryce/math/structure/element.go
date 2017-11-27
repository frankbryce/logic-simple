// Package structure contains data structures relevant for other packages
// in math to use, and also can be used by other developers that would like to
// extend github.com/frankbryce/math
package structure

import (
	"go/types"
)

const (
	// TODO(frankbryce): create more classes
	// TODO(frankbryce): is there a way to dynamically create classes instead?
	KNOWN_CLASSES = []string{"integer"}
)

// An element is an elementary object of mathematics, like a number, a shape,
// a point in space, or something more abstract.
type Element interface {
	GetClass() ElementClass
	String() string
	Value() interface{}
}

// Describes a set of elements that meet some given predicate
type ElementClass interface {
	IsMember(Element) bool
	Name() string
	Description() string
}

func GetIntegerClass() ElementClass {
	return &integerClass{}
}

type integerClass struct{}

func (ic *integerClass) IsMember(element Element) bool {
	switch t := element.Value().(type) {
	case int:
	case uint:
	case int32:
	case int64:
	case uint32:
	case uint64:
		return true
	default:
		return false
	}
}

func (ic *integerClass) Name() string {
	return "integer"
}

func (ic *integerClass) Description() string {
	return "The set of integers"
}

func (ic *integerClass) String() string {
	return ic.Description()
}
