class Solution {
    public int bestClosingTime(String customers) {
        int totalY = 0;
        for (char c : customers.toCharArray()) {
            if (c == 'Y')
                totalY++;
        }

        int openPenalty = 0;
        int closedPenalty = totalY;
        int minPenalty = closedPenalty;
        int answer = 0;

        for (int i = 0; i < customers.length(); i++) {
            if (customers.charAt(i) == 'N') {
                openPenalty++;
            } else {
                closedPenalty--;
            }

            int currentPenalty = openPenalty + closedPenalty;
            if (currentPenalty < minPenalty) {
                minPenalty = currentPenalty;
                answer = i + 1;
            }
        }

        return answer;
    }
}
