package main

import "embed"

const (
	UnknownID = -1
)

var (
	EbitengineLogoTextureID = UnknownID
	DefaultFontID           = UnknownID
)

//go:embed all:assets
var FS embed.FS
