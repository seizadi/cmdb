package runtime

import (
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

func PatternMatch(pattern runtime.Pattern, path string) bool {
	var components []string
	var idx, l int
	var c, verb string

	components = strings.Split(path[1:], "/")
	l = len(components)
	if idx = strings.LastIndex(components[l-1], ":"); idx > 0 {
		c = components[l-1]
		components[l-1], verb = c[:idx], c[idx+1:]
	}

	_, matchErr := pattern.Match(components, verb)
	return matchErr == nil
}

func JoinPath(path string, element string) string {
	if path == "" {
		return element
	}

	return path + "." + element
}
