class Solution {

    int count = 0;
    String result = "";

    private void dfs(int n, int k, StringBuilder curr) {

        if (!result.equals(""))
            return;

        if (curr.length() == n) {
            count++;
            if (count == k)
                result = curr.toString();
            return;
        }

        char[] chars = { 'a', 'b', 'c' };

        for (char c : chars) {

            if (curr.length() > 0 && curr.charAt(curr.length() - 1) == c)
                continue;

            curr.append(c);
            dfs(n, k, curr);
            curr.deleteCharAt(curr.length() - 1);
        }
    }

    public String getHappyString(int n, int k) {
        dfs(n, k, new StringBuilder());
        return result;
    }
}