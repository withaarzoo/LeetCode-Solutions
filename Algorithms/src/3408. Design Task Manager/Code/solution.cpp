#include <bits/stdc++.h>
using namespace std;

class TaskManager {
private:
    unordered_map<int, pair<int,int>> mp; // taskId -> {userId, priority}
    priority_queue<pair<int,int>> pq;     // {priority, taskId}

public:
    TaskManager(vector<vector<int>>& tasks) {
        // Initialize from list of [userId, taskId, priority]
        for (auto &t : tasks) {
            if (t.size() < 3) continue;
            int user = t[0], task = t[1], pr = t[2];
            mp[task] = {user, pr};
            pq.push({pr, task});
        }
    }
    
    void add(int userId, int taskId, int priority) {
        mp[taskId] = {userId, priority};
        pq.push({priority, taskId});
    }
    
    void edit(int taskId, int newPriority) {
        // guaranteed taskId exists
        auto &entry = mp[taskId];
        entry.second = newPriority; // update priority in map
        pq.push({newPriority, taskId});
    }
    
    void rmv(int taskId) {
        // guaranteed taskId exists
        mp.erase(taskId); // heap entry becomes stale and will be skipped later
    }
    
    int execTop() {
        while (!pq.empty()) {
            auto top = pq.top();
            int pr = top.first, id = top.second;
            pq.pop();
            auto it = mp.find(id);
            if (it == mp.end()) continue;           // task was removed/stale
            if (it->second.second != pr) continue;  // stale priority entry
            int userId = it->second.first;
            mp.erase(it); // remove executed task
            return userId;
        }
        return -1;
    }
};

/**
 * Your TaskManager object will be instantiated and called as such:
 * TaskManager* obj = new TaskManager(tasks);
 * obj->add(userId,taskId,priority);
 * obj->edit(taskId,newPriority);
 * obj->rmv(taskId);
 * int param_4 = obj->execTop();
 */
