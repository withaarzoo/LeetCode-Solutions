func parseBoolExpr(expression string) bool {
    stack := []rune{}

    for _, c := range expression {
        if c == ')' {
            subExpr := []rune{}
            for stack[len(stack)-1] != '(' {
                subExpr = append(subExpr, stack[len(stack)-1])
                stack = stack[:len(stack)-1]
            }
            stack = stack[:len(stack)-1] // Remove '('
            
            op := stack[len(stack)-1]
            stack = stack[:len(stack)-1] // Remove operator

            if op == '!' {
                if subExpr[0] == 't' {
                    stack = append(stack, 'f')
                } else {
                    stack = append(stack, 't')
                }
            } else if op == '&' {
                result := 't'
                for _, e := range subExpr {
                    if e == 'f' {
                        result = 'f'
                        break
                    }
                }
                stack = append(stack, result)
            } else if op == '|' {
                result := 'f'
                for _, e := range subExpr {
                    if e == 't' {
                        result = 't'
                        break
                    }
                }
                stack = append(stack, result)
            }
        } else if c != ',' {
            stack = append(stack, c)
        }
    }

    return stack[0] == 't'
}