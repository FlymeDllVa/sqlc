// Code generated by sqlc. DO NOT EDIT.

package querytest

import (
	"net"
)

type Foo struct {
	Bar  bool
	Inet net.IP
	Cidr net.IP
}