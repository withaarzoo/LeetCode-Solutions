class Solution
{
public:
    int mostBooked(int n, vector<vector<int>> &meetings)
    {
        sort(meetings.begin(), meetings.end());

        // Min heap of free rooms
        priority_queue<int, vector<int>, greater<int>> freeRooms;
        for (int i = 0; i < n; i++)
            freeRooms.push(i);

        // Min heap of busy rooms: {endTime, room}
        priority_queue<pair<long long, int>, vector<pair<long long, int>>, greater<>> busyRooms;

        vector<long long> count(n, 0);

        for (auto &m : meetings)
        {
            long long start = m[0], end = m[1];
            long long duration = end - start;

            // Free rooms that have completed meetings
            while (!busyRooms.empty() && busyRooms.top().first <= start)
            {
                freeRooms.push(busyRooms.top().second);
                busyRooms.pop();
            }

            if (!freeRooms.empty())
            {
                int room = freeRooms.top();
                freeRooms.pop();
                busyRooms.push({end, room});
                count[room]++;
            }
            else
            {
                auto [finish, room] = busyRooms.top();
                busyRooms.pop();
                busyRooms.push({finish + duration, room});
                count[room]++;
            }
        }

        int ans = 0;
        for (int i = 1; i < n; i++)
        {
            if (count[i] > count[ans])
                ans = i;
        }
        return ans;
    }
};
