package util

type Strings []string

func (strs *Strings) Has(str string) bool {
	for _, s := range *strs {
		if s == str {
			return true
		}
	}
	return false
}

func (strs *Strings) Join(op string) string {
	res := ""
	length := len(*strs)
	if length == 0 {
		return ""
	} else if length == 1 {
		return (*strs)[0]
	}

	for i := 0; i < length-1; i++ {
		res += (*strs)[i] + op
	}
	res += (*strs)[length-1]

	return res
}
