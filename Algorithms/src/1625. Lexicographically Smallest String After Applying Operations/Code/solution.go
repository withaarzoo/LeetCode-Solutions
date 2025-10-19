func findLexSmallestString(s string, a int, b int) string {
	n := len(s)
	seen := make(map[string]bool)
	q := list.New()
	seen[s] = true
	q.PushBack(s)
	ans := s

	for q.Len() > 0 {
		front := q.Remove(q.Front()).(string)
		if front < ans {
			ans = front
		}

		// add a to odd indices
		addBytes := []byte(front)
		for i := 1; i < n; i += 2 {
			d := int(addBytes[i]-'0')
			d = (d + a) % 10
			addBytes[i] = byte('0' + d)
		}
		addOp := string(addBytes)
		if !seen[addOp] {
			seen[addOp] = true
			q.PushBack(addOp)
		}

		// rotate right by b
		rotOp := front[n-b:] + front[:n-b]
		if !seen[rotOp] {
			seen[rotOp] = true
			q.PushBack(rotOp)
		}
	}

	return ans
}