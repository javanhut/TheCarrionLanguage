package munin

import (
	"embed"
)

//go:embed *.crl
var MuninFs embed.FS
