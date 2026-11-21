class Solution {
    public int[] findSubtreeSizes(int[] parent, String s) {
        int n = parent.length;
        List<Integer>[] g = new ArrayList[n];
        for (int i = 0; i < n; i++) g[i] = new ArrayList<>();
        for (int i = 1; i < n; i++) g[parent[i]].add(i);
        int[] np = parent.clone();
        int[] last = new int[26];
        Arrays.fill(last, -1);
        dfs(0, s, g, np, last);
        List<Integer>[] ng = new ArrayList[n];
        for (int i = 0; i < n; i++) ng[i] = new ArrayList<>();
        for (int i = 1; i < n; i++) ng[np[i]].add(i);
        int[] ans = new int[n];
        size(0, ng, ans);
        return ans;
    }
    void dfs(int u, String s, List<Integer>[] g, int[] np, int[] last) {
        int c = s.charAt(u) - 'a';
        int prev = last[c];
        if (prev != -1) np[u] = prev;
        last[c] = u;
        for (int v : g[u]) dfs(v, s, g, np, last);
        last[c] = prev;
    }
    int size(int u, List<Integer>[] ng, int[] ans) {
        int sz = 1;
        for (int v : ng[u]) sz += size(v, ng, ans);
        return ans[u] = sz;
    }
}