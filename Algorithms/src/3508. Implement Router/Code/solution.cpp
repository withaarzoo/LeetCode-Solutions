#include <bits/stdc++.h>
using namespace std;

class Router {
private:
    int memoryLimit;
    deque<tuple<int,int,int>> q;                 // global FIFO queue of packets
    unordered_set<string> seen;                  // set of "s#d#t" keys to detect duplicates
    unordered_map<int, vector<int>> times;       // per-destination appended timestamps
    unordered_map<int,int> head;                 // per-destination head index (how many removed)

    // helper to make a unique string key for a packet
    string makeKey(int s, int d, int t) {
        return to_string(s) + "#" + to_string(d) + "#" + to_string(t);
    }

public:
    Router(int memoryLimit) : memoryLimit(memoryLimit) {}

    bool addPacket(int source, int destination, int timestamp) {
        string key = makeKey(source, destination, timestamp);
        if (seen.count(key)) return false; // duplicate present

        // Evict oldest until we have room (normally at most one)
        while ((int)q.size() >= memoryLimit) {
            auto [os, od, ot] = q.front();
            q.pop_front();
            seen.erase(makeKey(os, od, ot));
            // advance head for that destination (lazy removal)
            head[od] += 1;
            // we do not physically remove from times[od] vector to keep append O(1)
        }

        // Add new packet
        q.emplace_back(source, destination, timestamp);
        seen.insert(key);
        times[destination].push_back(timestamp);
        // head[destination] defaults to 0 if not present
        return true;
    }

    vector<int> forwardPacket() {
        if (q.empty()) return {}; // no packets

        auto [s, d, t] = q.front();
        q.pop_front();
        seen.erase(makeKey(s, d, t));
        head[d] += 1; // advance head for that destination

        return {s, d, t};
    }

    int getCount(int destination, int startTime, int endTime) {
        auto it = times.find(destination);
        if (it == times.end()) return 0;
        vector<int> &arr = it->second;
        int h = head[destination]; // might be 0 if not present

        // search in arr[h .. end)
        auto loIt = lower_bound(arr.begin() + h, arr.end(), startTime);
        auto hiIt = upper_bound(arr.begin() + h, arr.end(), endTime);
        return (int)(hiIt - loIt);
    }
};

/**
 * Your Router object will be instantiated and called as such:
 * Router* obj = new Router(memoryLimit);
 * bool param_1 = obj->addPacket(source,destination,timestamp);
 * vector<int> param_2 = obj->forwardPacket();
 * int param_3 = obj->getCount(destination,startTime,endTime);
 */
