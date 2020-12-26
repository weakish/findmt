package main

import (
	"encoding/json"
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

	patterns := strings.Split(mimeTypePattern, ",")
	var includes []string
	var excludes []string
	for _, pattern := range patterns {
		pattern = strings.TrimSuffix(pattern, "*")
		if strings.HasPrefix(pattern, "-") {
			excludes = append(excludes, strings.TrimPrefix(pattern, "-"))
		} else {
			includes = append(includes, pattern)
		}
	}

	for _, exclude := range excludes {
		if strings.HasPrefix(mimeType, exclude) {
			return false
		}
	}

	for _, include := range includes {
		if strings.HasPrefix(mimeType, include) {
			return true
		}
	}

	return false
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
				jsonBytes, err := json.Marshal(parent)
				if err == nil {
					fmt.Println(string(jsonBytes))
				} else {
					_, _ = fmt.Fprintf(os.Stderr, "failed to encode `%v` as json", err)
				}

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
	const usage = "findmt [pattern]\n\nExample: findmt 'image/*,-image/vnd.djvu'\n"
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
