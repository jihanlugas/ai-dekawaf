package helper

import "strings"

// "Hello World Golang" returns "hello-world-golang"
func Slugify(value string) string {
	return strings.ReplaceAll(strings.ToLower(value), " ", "-")
}

// "/Common/test" returns "~Common~test"
func FormatF5Path(path string) string {
	return strings.ReplaceAll(path, "/", "~")
}
