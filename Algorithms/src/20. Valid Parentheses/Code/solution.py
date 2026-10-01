class Solution:

    def isValid(self, s: str) -> bool:

        st = []

        for c in s:

            if c in '({[':

                st.append(c)

            else:

                if not st:

                    return False

                top = st.pop()

                if (c == ')' and top != '(') or (c == '}' and top != '{') or (c == ']' and top != '['):

                    return False

        return not st