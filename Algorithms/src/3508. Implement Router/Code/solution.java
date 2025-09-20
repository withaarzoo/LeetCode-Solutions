import java.util.*;

class Router {
    private int memoryLimit;
    private Deque<int[]> q;                       // global FIFO queue of [s,d,t]
    private Set<String> seen;                     // "s#d#t" keys
    private Map<Integer, ArrayList<Integer>> times; // dest -> list of timestamps (append-only)
    private Map<Integer, Integer> head;            // dest -> head index

    private String makeKey(int s, int d, int t) {
        return s + "#" + d + "#" + t;
    }

    public Router(int memoryLimit) {
        this.memoryLimit = memoryLimit;
        q = new ArrayDeque<>();
        seen = new HashSet<>();
        times = new HashMap<>();
        head = new HashMap<>();
    }

    public boolean addPacket(int source, int destination, int timestamp) {
        String key = makeKey(source, destination, timestamp);
        if (seen.contains(key)) return false;

        while (q.size() >= memoryLimit) {
            int[] old = q.pollFirst();
            String oldKey = makeKey(old[0], old[1], old[2]);
            seen.remove(oldKey);
            head.put(old[1], head.getOrDefault(old[1], 0) + 1);
        }

        q.addLast(new int[]{source, destination, timestamp});
        seen.add(key);
        times.computeIfAbsent(destination, k -> new ArrayList<>()).add(timestamp);
        return true;
    }

    public int[] forwardPacket() {
        if (q.isEmpty()) return new int[0];
        int[] pkt = q.pollFirst();
        String key = makeKey(pkt[0], pkt[1], pkt[2]);
        seen.remove(key);
        head.put(pkt[1], head.getOrDefault(pkt[1], 0) + 1);
        return pkt;
    }

    // Helper binary search methods using provided lo index
    private int lowerBound(ArrayList<Integer> arr, int target, int lo) {
        int l = lo, r = arr.size();
        while (l < r) {
            int m = l + (r - l) / 2;
            if (arr.get(m) < target) l = m + 1;
            else r = m;
        }
        return l;
    }
    private int upperBound(ArrayList<Integer> arr, int target, int lo) {
        int l = lo, r = arr.size();
        while (l < r) {
            int m = l + (r - l) / 2;
            if (arr.get(m) <= target) l = m + 1;
            else r = m;
        }
        return l;
    }

    public int getCount(int destination, int startTime, int endTime) {
        ArrayList<Integer> arr = times.get(destination);
        if (arr == null) return 0;
        int h = head.getOrDefault(destination, 0);
        int L = lowerBound(arr, startTime, h);
        int R = upperBound(arr, endTime, h);
        return R - L;
    }
}

/**
 * Your Router object will be instantiated and called as such:
 * Router obj = new Router(memoryLimit);
 * boolean param_1 = obj.addPacket(source,destination,timestamp);
 * int[] param_2 = obj.forwardPacket();
 * int param_3 = obj.getCount(destination,startTime,endTime);
 */
