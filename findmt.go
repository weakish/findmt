package main

import (
	"fmt"
	"github.com/weakish/gosugar"
	"mime"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func matchesMimeType(f os.FileInfo, mimeTypePattern string) bool {
	ext := filepath.Ext(f.Name())
	mimeType := mime.TypeByExtension(ext)

	mimeTypePrefix := strings.TrimSuffix(mimeTypePattern, "*")
	if strings.HasPrefix(mimeType, mimeTypePrefix) {
		return true
	} else {
		return false
	}
}

func findmt(mimeTypePattern string) {
	_ = filepath.Walk(".", func(filePath string, f os.FileInfo, err error) error {
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to access %q: %v\n", filePath, err)
			return err
		} else if gosugar.IsUnixHiddenFile(f.Name()) {
			return gosugar.SkipDirOrFile(f)
		} else if f.Mode().IsRegular() {
			if matchesMimeType(f, mimeTypePattern) {
				parent := path.Dir(filePath)
				fmt.Println("\"" + parent + "\"")
				return filepath.SkipDir
			} else {
				return nil
			}
		} else if f.IsDir() {
			return nil
		} else {
			return nil
		}
	})
}

func main() {
	const usage = "findmt [pattern]\n\nExample: findmt 'image/*'\n"
	switch len(os.Args) {
	case 1:
		_, _ = os.Stderr.WriteString(usage)
		os.Exit(64)
	case 2:
		pattern := os.Args[1]
		if pattern == "-h" || pattern == "--help" || pattern == "help" {
			_, _ = os.Stderr.WriteString(usage)
		} else if strings.Contains(pattern, "/") {
			findmt(pattern)
		} else {
			_, _ = os.Stderr.WriteString("invalid pattern\n")
			os.Exit(64)
		}
	default:
		_, _ = os.Stderr.WriteString(usage)
		os.Exit(64)
	}
}
