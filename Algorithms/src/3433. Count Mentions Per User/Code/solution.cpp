/* C++ */
#include <bits/stdc++.h>
using namespace std;

class Solution
{
public:
    vector<int> countMentions(int numberOfUsers, vector<vector<string>> &events)
    {
        // group events by timestamp
        map<int, vector<vector<string>>> byTime;
        for (auto &ev : events)
        {
            int t = stoi(ev[1]);
            byTime[t].push_back(ev);
        }

        vector<int> mentions(numberOfUsers, 0);
        vector<bool> isOnline(numberOfUsers, true);
        vector<int> offlineUntil(numberOfUsers, 0); // 0 means currently online

        for (auto &entry : byTime)
        {
            int t = entry.first;
            auto &evs = entry.second;

            // 1) process expirations (users that become online at or before t)
            for (int i = 0; i < numberOfUsers; ++i)
            {
                if (!isOnline[i] && offlineUntil[i] <= t)
                {
                    isOnline[i] = true;
                    offlineUntil[i] = 0;
                }
            }

            // 2) first apply all OFFLINE events at this timestamp
            for (auto &ev : evs)
            {
                if (ev[0] == "OFFLINE")
                {
                    int id = stoi(ev[2]);
                    // mark offline for exactly 60 units (t..t+59), back online at t+60
                    isOnline[id] = false;
                    offlineUntil[id] = t + 60;
                }
            }

            // 3) then handle MESSAGE events at this timestamp
            for (auto &ev : evs)
            {
                if (ev[0] != "MESSAGE")
                    continue;
                string mentionsStr = ev[2];
                // split by spaces
                string token;
                stringstream ss(mentionsStr);
                while (ss >> token)
                {
                    if (token == "ALL")
                    {
                        for (int i = 0; i < numberOfUsers; ++i)
                            mentions[i]++;
                    }
                    else if (token == "HERE")
                    {
                        for (int i = 0; i < numberOfUsers; ++i)
                            if (isOnline[i])
                                mentions[i]++;
                    }
                    else if (token.rfind("id", 0) == 0)
                    {
                        int id = stoi(token.substr(2));
                        if (0 <= id && id < numberOfUsers)
                            mentions[id]++;
                    }
                }
            }
        }

        return mentions;
    }
};
