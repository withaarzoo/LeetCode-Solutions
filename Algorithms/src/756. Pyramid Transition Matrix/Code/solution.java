class Solution {
    Map<String, List<Character>> rules = new HashMap<>();
    Set<String> bad = new HashSet<>();

    public boolean pyramidTransition(String bottom, List<String> allowed) {
        for (String s : allowed) {
            rules.computeIfAbsent(s.substring(0, 2), k -> new ArrayList<>())
                    .add(s.charAt(2));
        }
        return dfs(bottom, 0, new StringBuilder());
    }

    private boolean dfs(String row, int idx, StringBuilder next) {
        if (row.length() == 1)
            return true;

        if (idx == row.length() - 1) {
            String nextRow = next.toString();
            if (bad.contains(nextRow))
                return false;
            boolean ok = dfs(nextRow, 0, new StringBuilder());
            if (!ok)
                bad.add(nextRow);
            return ok;
        }

        String key = row.substring(idx, idx + 2);
        if (!rules.containsKey(key))
            return false;

        for (char c : rules.get(key)) {
            next.append(c);
            if (dfs(row, idx + 1, next))
                return true;
            next.deleteCharAt(next.length() - 1);
        }
        return false;
    }
}
