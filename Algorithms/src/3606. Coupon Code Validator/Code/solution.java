class Solution {
    public List<String> validateCoupons(String[] code,
            String[] businessLine,
            boolean[] isActive) {

        Map<String, Integer> priority = new HashMap<>();
        priority.put("electronics", 0);
        priority.put("grocery", 1);
        priority.put("pharmacy", 2);
        priority.put("restaurant", 3);

        List<String[]> valid = new ArrayList<>();

        for (int i = 0; i < code.length; i++) {

            if (!isActive[i])
                continue;
            if (!priority.containsKey(businessLine[i]))
                continue;
            if (code[i].isEmpty())
                continue;

            boolean ok = true;
            for (char c : code[i].toCharArray()) {
                if (!Character.isLetterOrDigit(c) && c != '_') {
                    ok = false;
                    break;
                }
            }
            if (!ok)
                continue;

            valid.add(new String[] { businessLine[i], code[i] });
        }

        valid.sort((a, b) -> {
            int p1 = priority.get(a[0]);
            int p2 = priority.get(b[0]);
            if (p1 == p2)
                return a[1].compareTo(b[1]);
            return p1 - p2;
        });

        List<String> result = new ArrayList<>();
        for (String[] v : valid) {
            result.add(v[1]);
        }

        return result;
    }
}
