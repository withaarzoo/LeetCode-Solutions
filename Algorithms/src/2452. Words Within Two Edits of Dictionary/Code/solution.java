class Solution {
    public List<String> twoEditWords(String[] queries, String[] dictionary) {
        List<String> result = new ArrayList<>();

        // Check every query word
        for (String query : queries) {

            // Compare with every dictionary word
            for (String word : dictionary) {
                int diff = 0;

                // Count different characters
                for (int i = 0; i < query.length(); i++) {
                    if (query.charAt(i) != word.charAt(i)) {
                        diff++;
                    }

                    // Stop early if more than 2 edits are needed
                    if (diff > 2) {
                        break;
                    }
                }

                // If query can match this dictionary word
                if (diff <= 2) {
                    result.add(query);
                    break;
                }
            }
        }

        return result;
    }
}