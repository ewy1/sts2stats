package main

import (
	"github.com/spf13/pflag"
)

var Profile = pflag.StringP("profile", "p", ".", "slay the spire 2 profile folder")
