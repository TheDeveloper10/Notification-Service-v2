package util

import "strings"

type TemplateString string

func (ts *TemplateString) Fill(values map[string]string) string {
	// TODO: optimize
	str := string(*ts)

	for key, value := range values {
		fullKey := "%" + key + "%"

		str = strings.ReplaceAll(str, fullKey, value)
	}

	return str
}
