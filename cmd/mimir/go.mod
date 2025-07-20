module github.com/javanhut/TheCarrionLanguage/cmd/mimir

go 1.23.0

toolchain go1.24.5

replace github.com/javanhut/TheCarrionLanguage => ../..

require github.com/peterh/liner v1.2.2

require (
	github.com/mattn/go-runewidth v0.0.3 // indirect
	golang.org/x/sys v0.33.0 // indirect
)
