package medium_type

import "strings"

func simplifyPath(path string) string {
	var res []string
	for _, char := range strings.Split(path, "/") {
		if char == "." {
			continue
		} else if char == ".." {
			if len(res) > 0 {
				res = res[:len(res)]
			}
		} else {
			res = append(res, char)
		}
	}

	return strings.Join(res, "/")
}
