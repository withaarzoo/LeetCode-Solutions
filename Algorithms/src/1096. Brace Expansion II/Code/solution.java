class Solution {
    private int i;
    private String expr;
    public List<String> braceExpansionII(String expression) {
        this.expr = expression;
        this.i = 0;
        Set<String> res = parse();
        List<String> ans = new ArrayList<>(res);
        Collections.sort(ans);
        return ans;
    }
    private Set<String> parse() {
        Set<String> res = new TreeSet<>();
        Set<String> cur = new TreeSet<>();
        cur.add("");
        while (i < expr.length() && expr.charAt(i) != '}') {
            if (expr.charAt(i) == '{') {
                i++;
                Set<String> next = parse();
                i++;
                cur = product(cur, next);
            } else if (expr.charAt(i) == ',') {
                res.addAll(cur);
                cur = new TreeSet<>();
                cur.add("");
                i++;
            } else {
                Set<String> next = new TreeSet<>();
                next.add(String.valueOf(expr.charAt(i)));
                i++;
                cur = product(cur, next);
            }
        }
        res.addAll(cur);
        return res;
    }
    private Set<String> product(Set<String> a, Set<String> b) {
        Set<String> res = new TreeSet<>();
        for (String x : a)
            for (String y : b)
                res.add(x + y);
        return res;
    }
}