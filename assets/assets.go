package assets

import (
	"embed"
)

const (
	DirMedia  = "media"
	DirStatic = "static"
)

//go:embed static/*
var StaticFS embed.FS
