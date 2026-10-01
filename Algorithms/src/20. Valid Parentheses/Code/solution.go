func isValid(s string) bool {

    st := []rune{}

    for _, c := range s {

        if c == '(' || c == '{' || c == '[' {

            st = append(st, c)

        } else {

            if len(st) == 0 {

                return false

            }

            top := st[len(st)-1]

            st = st[:len(st)-1]

            if (c == ')' && top != '(') || (c == '}' && top != '{') || (c == ']' && top != '[') {

                return false

            }

        }

    }

    return len(st) == 0

}