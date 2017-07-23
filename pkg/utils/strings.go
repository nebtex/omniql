package utils

import (
	"strings"
	"regexp"
)


var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake  = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func CommentGolang(s string) (ns string) {
	lines := strings.Split(s, "\n")

	for i := 0; i < len(lines); i++ {
		if i == 0 {
			ns = ns + "//" + lines[i]
		} else {
			ns = ns + "\n//" + lines[i]
		}
	}
	return
}

func GetVersion(s string) (v string) {
	r := strings.Split(s, ".")
	if len(r) > 0 {
		v = r[len(r)-1]
	}
	return

}

func GolangGetPackageName(app string) (v string) {
	r := strings.Split(app, ".")
	if len(r) > 2 {
		v = strings.ToLower(r[len(r)-2]) + strings.ToLower(r[len(r)-1])
	}
	return

}
