package easy_type

import "strings"

func numUniqueEmails(emails []string) int {
	var uniqeEmails []string
	var count int

	for _, val := range emails {
		res := strings.Split(val, "@")
		email := getValidLocalName(res[0]) + "@" + res[1]
		if !isEmailExists(email, uniqeEmails) {
			uniqeEmails = append(uniqeEmails, email)
			count += 1
		}
	}

	return count
}

func getValidLocalName(name string) string {
	var res string
	for _, val := range name {
		switch string(val) {
		case ".":
			continue
		case "+":
			return res
		default:
			res = res + string(val)
		}
	}
	return res
}

func isEmailExists(email string, uniqueEmails []string) bool {
	for _, val := range uniqueEmails {
		if val == email {
			return true
		}
	}
	return false
}
