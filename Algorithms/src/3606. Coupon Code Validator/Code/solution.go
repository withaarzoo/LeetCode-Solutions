func validateCoupons(code []string, businessLine []string, isActive []bool) []string {

	priority := map[string]int{
		"electronics": 0,
		"grocery":     1,
		"pharmacy":    2,
		"restaurant":  3,
	}

	type Pair struct {
		p int
		c string
	}

	valid := []Pair{}

	for i := 0; i < len(code); i++ {

		if !isActive[i] {
			continue
		}

		p, ok := priority[businessLine[i]]
		if !ok {
			continue
		}

		if len(code[i]) == 0 {
			continue
		}

		validCode := true
		for _, ch := range code[i] {
			if !(ch >= 'a' && ch <= 'z' ||
				ch >= 'A' && ch <= 'Z' ||
				ch >= '0' && ch <= '9' ||
				ch == '_') {
				validCode = false
				break
			}
		}
		if !validCode {
			continue
		}

		valid = append(valid, Pair{p, code[i]})
	}

	sort.Slice(valid, func(i, j int) bool {
		if valid[i].p == valid[j].p {
			return valid[i].c < valid[j].c
		}
		return valid[i].p < valid[j].p
	})

	result := []string{}
	for _, v := range valid {
		result = append(result, v.c)
	}

	return result
}
