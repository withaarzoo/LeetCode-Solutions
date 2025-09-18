import java.util.*;

class TaskManager {
    private static class Item {
        int priority, taskId;
        Item(int p, int t) { priority = p; taskId = t; }
    }
    
    private final Map<Integer, int[]> map; // taskId -> [userId, priority]
    private final PriorityQueue<Item> pq;
    
    public TaskManager(List<List<Integer>> tasks) {
        map = new HashMap<>();
        pq = new PriorityQueue<>((a, b) -> {
            if (a.priority != b.priority) return Integer.compare(b.priority, a.priority);
            return Integer.compare(b.taskId, a.taskId);
        });
        for (List<Integer> t : tasks) {
            if (t.size() < 3) continue;
            int user = t.get(0), task = t.get(1), pr = t.get(2);
            map.put(task, new int[]{user, pr});
            pq.offer(new Item(pr, task));
        }
    }
    
    public void add(int userId, int taskId, int priority) {
        map.put(taskId, new int[]{userId, priority});
        pq.offer(new Item(priority, taskId));
    }
    
    public void edit(int taskId, int newPriority) {
        // guaranteed to exist
        int[] arr = map.get(taskId);
        arr[1] = newPriority;
        pq.offer(new Item(newPriority, taskId));
    }
    
    public void rmv(int taskId) {
        map.remove(taskId); // heap entry is left stale
    }
    
    public int execTop() {
        while (!pq.isEmpty()) {
            Item it = pq.poll();
            int id = it.taskId, pr = it.priority;
            int[] arr = map.get(id);
            if (arr == null) continue;           // removed
            if (arr[1] != pr) continue;          // stale priority
            int userId = arr[0];
            map.remove(id);
            return userId;
        }
        return -1;
    }
}

/**
 * Your TaskManager object will be instantiated and called as such:
 * TaskManager obj = new TaskManager(tasks);
 * obj.add(userId,taskId,priority);
 * obj.edit(taskId,newPriority);
 * obj.rmv(taskId);
 * int param_4 = obj.execTop();
 */
