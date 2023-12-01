//go:build windows || !cgo
// +build windows !cgo

package postgresql

import (
	nodes "github.com/wasilibs/go-pgquery"
)

var parseNodes = nodes.Parse