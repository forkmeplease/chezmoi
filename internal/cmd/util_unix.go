//go:build unix

package cmd

import (
	"io/fs"
	"syscall"

	"github.com/twpayne/chezmoi/internal/chezmoi"
)

const defaultEditor = "vi"

var defaultInterpreters = make(map[string]chezmoi.Interpreter)

func fileInfoUID(info fs.FileInfo) int {
	return int(info.Sys().(*syscall.Stat_t).Uid) //nolint:forcetypeassert
}

func windowsVersion() (map[string]any, error) {
	return nil, nil
}
