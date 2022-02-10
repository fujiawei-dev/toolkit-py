//go:build !autogen
//+build !autogen

{{GOLANG_HEADER}}

// Package version is auto-generated at build-time
package {{GOLANG_PACKAGE}}

const LibraryImport = "library-import"

// Default build-time variable for library-import.
// These variables are overridden on build with build-time information.
var (
	BuildTime = LibraryImport
	GitCommit = LibraryImport
	Name      = LibraryImport
	Version   = LibraryImport
)
