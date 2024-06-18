package staticfiles

import "embed"

//go:embed files
var embeddedFiles embed.FS
