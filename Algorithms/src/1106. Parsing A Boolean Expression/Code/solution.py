class Solution:
    def parseBoolExpr(self, expression: str) -> bool:
        stack = []

        for c in expression:
            if c == ')':
                sub_expr = []
                while stack[-1] != '(':
                    sub_expr.append(stack.pop())
                stack.pop()  # remove '('
                op = stack.pop()  # get the operator

                if op == '!':
                    stack.append('f' if sub_expr[0] == 't' else 't')
                elif op == '&':
                    result = 't'
                    for e in sub_expr:
                        if e == 'f':
                            result = 'f'
                            break
                    stack.append(result)
                elif op == '|':
                    result = 'f'
                    for e in sub_expr:
                        if e == 't':
                            result = 't'
                            break
                    stack.append(result)
            elif c != ',':
                stack.append(c)

        return stack[0] == 't'