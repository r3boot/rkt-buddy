package system

import "regexp"

type SystemData struct {
	Hostname  string
	Address   string
	Interface string
}

var (
	reAddress   = regexp.MustCompile("inet ([0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3})/[0-9]{1,2} scope global ")
	reInterface = regexp.MustCompile("^([a-z]{2,3}[0-9]).*Link encap:Ethernet")
)
