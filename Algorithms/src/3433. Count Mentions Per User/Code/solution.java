
// Java
import java.util.*;

class Solution {
    public int[] countMentions(int numberOfUsers, List<List<String>> events) {
        // group events by timestamp (sorted by timestamp via TreeMap)
        TreeMap<Integer, List<List<String>>> byTime = new TreeMap<>();
        for (List<String> ev : events) {
            int t = Integer.parseInt(ev.get(1));
            byTime.computeIfAbsent(t, k -> new ArrayList<>()).add(ev);
        }

        int[] mentions = new int[numberOfUsers];
        boolean[] isOnline = new boolean[numberOfUsers];
        int[] offlineUntil = new int[numberOfUsers];
        Arrays.fill(isOnline, true); // initially all online

        for (Map.Entry<Integer, List<List<String>>> entry : byTime.entrySet()) {
            int t = entry.getKey();
            List<List<String>> evs = entry.getValue();

            // 1) expirations: bring users back online if offlineUntil <= t
            for (int i = 0; i < numberOfUsers; ++i) {
                if (!isOnline[i] && offlineUntil[i] <= t) {
                    isOnline[i] = true;
                    offlineUntil[i] = 0;
                }
            }

            // 2) apply OFFLINE events first at this timestamp
            for (List<String> ev : evs) {
                if (ev.get(0).equals("OFFLINE")) {
                    int id = Integer.parseInt(ev.get(2));
                    isOnline[id] = false;
                    offlineUntil[id] = t + 60;
                }
            }

            // 3) process MESSAGE events
            for (List<String> ev : evs) {
                if (!ev.get(0).equals("MESSAGE"))
                    continue;
                String mentionsStr = ev.get(2);
                String[] tokens = mentionsStr.split("\\s+");
                for (String token : tokens) {
                    if (token.equals("ALL")) {
                        for (int i = 0; i < numberOfUsers; ++i)
                            mentions[i]++;
                    } else if (token.equals("HERE")) {
                        for (int i = 0; i < numberOfUsers; ++i)
                            if (isOnline[i])
                                mentions[i]++;
                    } else if (token.startsWith("id")) {
                        int id = Integer.parseInt(token.substring(2));
                        if (id >= 0 && id < numberOfUsers)
                            mentions[id]++;
                    }
                }
            }
        }

        return mentions;
    }
}
