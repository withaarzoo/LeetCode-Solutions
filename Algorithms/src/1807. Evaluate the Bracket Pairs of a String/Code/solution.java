class Solution {
    public String evaluate(String s, List<List<String>> knowledge) {
        Map<String, String> mp = new HashMap<>();
        for (List<String> p : knowledge) mp.put(p.get(0), p.get(1));
        StringBuilder res = new StringBuilder();
        int n = s.length();
        for (int i = 0; i < n; ) {
            if (s.charAt(i) == '(') {
                int j = i + 1;
                while (s.charAt(j) != ')') ++j;
                String key = s.substring(i + 1, j);
                res.append(mp.getOrDefault(key, "?"));
                i = j + 1;
            } else {
                res.append(s.charAt(i++));
            }
        }
        return res.toString();
    }
}