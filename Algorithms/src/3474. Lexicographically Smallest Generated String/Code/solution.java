class Solution {
    public String generateString(String str1, String str2) {
        int n = str1.length();
        int m = str2.length();
        int len = n + m - 1;

        char[] ans = new char[len];
        boolean[] fixed = new boolean[len];

        for (int i = 0; i < len; i++) {
            ans[i] = '?';
        }

        // Apply all 'T' constraints
        for (int i = 0; i < n; i++) {
            if (str1.charAt(i) == 'T') {
                for (int j = 0; j < m; j++) {
                    int pos = i + j;
                    char ch = str2.charAt(j);

                    if (ans[pos] != '?' && ans[pos] != ch) {
                        return "";
                    }

                    ans[pos] = ch;
                    fixed[pos] = true;
                }
            }
        }

        // Fill remaining positions with 'a'
        for (int i = 0; i < len; i++) {
            if (ans[i] == '?')
                ans[i] = 'a';
        }

        // Process all 'F' constraints
        for (int i = 0; i < n; i++) {
            if (str1.charAt(i) == 'F') {
                boolean same = true;

                for (int j = 0; j < m; j++) {
                    if (ans[i + j] != str2.charAt(j)) {
                        same = false;
                        break;
                    }
                }

                if (!same)
                    continue;

                boolean changed = false;

                for (int j = m - 1; j >= 0; j--) {
                    int pos = i + j;

                    if (fixed[pos])
                        continue;

                    for (char c = 'a'; c <= 'z'; c++) {
                        if (c != ans[pos] && c != str2.charAt(j)) {
                            ans[pos] = c;
                            changed = true;
                            break;
                        }
                    }

                    if (changed)
                        break;
                }

                if (!changed)
                    return "";
            }
        }

        return new String(ans);
    }
}