//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris

package resolver

import _ "unsafe"

//go:linkname defaultNS net.defaultNS
var defaultNS []string

func init() {
	defaultNS = []string{"114.114.114.114:53", "223.5.5.5:53"}
}
