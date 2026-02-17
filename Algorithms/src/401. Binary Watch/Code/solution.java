class Solution {
    public List<String> readBinaryWatch(int turnedOn) {
        List<String> result = new ArrayList<>();

        // Try all possible hours
        for (int hour = 0; hour < 12; hour++) {

            // Try all possible minutes
            for (int minute = 0; minute < 60; minute++) {

                // Count set bits of hour and minute
                int totalBits = Integer.bitCount(hour) + Integer.bitCount(minute);

                if (totalBits == turnedOn) {

                    // Format minute with leading zero
                    String time = hour + ":" +
                            (minute < 10 ? "0" + minute : minute);

                    result.add(time);
                }
            }
        }

        return result;
    }
}
