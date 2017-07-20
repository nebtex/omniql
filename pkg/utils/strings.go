package utils

import "strings"

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
