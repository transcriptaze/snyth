package html

import (
	"embed"
)

//go:embed index.html css images fonts javascript favicon.ico
var HTML embed.FS
