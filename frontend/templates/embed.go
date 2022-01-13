package templates

import (
	"embed"
)

//go:embed public/*
var Frontend embed.FS
