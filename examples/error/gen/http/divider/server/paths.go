// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// HTTP request path constructors for the divider service.
//
// Command:
// $ goa gen goa.design/goa/examples/error/design

package server

import (
	"fmt"
)

// IntegerDivideDividerPath returns the URL path to the divider service integer_divide HTTP endpoint.
func IntegerDivideDividerPath(a int, b int) string {
	return fmt.Sprintf("/idiv/%v/%v", a, b)
}

// DivideDividerPath returns the URL path to the divider service divide HTTP endpoint.
func DivideDividerPath(a float64, b float64) string {
	return fmt.Sprintf("/div/%v/%v", a, b)
}
