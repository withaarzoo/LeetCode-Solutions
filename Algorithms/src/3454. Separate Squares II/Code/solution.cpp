class Solution {
    vector<int> cnt;
    vector<double> segLen;
    vector<double> xs;

    void update(int idx, int l, int r, int ql, int qr, int val) {
        if (qr <= l || r <= ql) return;
        if (ql <= l && r <= qr) {
            cnt[idx] += val;
        } else {
            int m = (l + r) >> 1;
            update(idx<<1, l, m, ql, qr, val);
            update(idx<<1|1, m, r, ql, qr, val);
        }

        if (cnt[idx] > 0) {
            segLen[idx] = xs[r] - xs[l];
        } else if (r - l == 1) {
            segLen[idx] = 0.0;
        } else {
            segLen[idx] = segLen[idx<<1] + segLen[idx<<1|1];
        }
    }

public:
    double separateSquares(vector<vector<int>>& squares) {
        for (auto &s : squares) {
            xs.push_back((double)s[0]);
            xs.push_back((double)s[0] + (double)s[2]);
        }
        sort(xs.begin(), xs.end());
        xs.erase(unique(xs.begin(), xs.end()), xs.end());

        struct Event {
            double y, x1, x2;
            int type;
        };

        vector<Event> events;
        for (auto &s : squares) {
            double x = (double)s[0];
            double y = (double)s[1];
            double l = (double)s[2];
            events.push_back({y, x, x + l, 1});
            events.push_back({y + l, x, x + l, -1});
        }

        sort(events.begin(), events.end(),
             [](auto &a, auto &b) { return a.y < b.y; });

        int n = xs.size();
        cnt.assign(4 * n, 0);
        segLen.assign(4 * n, 0.0);

        vector<array<double,3>> strips;
        double total = 0.0;
        double prevY = events[0].y;

        for (auto &e : events) {
            if (e.y > prevY) {
                double w = segLen[1];
                double h = e.y - prevY;
                total += w * h;
                strips.push_back({prevY, h, w});
                prevY = e.y;
            }
            int l = lower_bound(xs.begin(), xs.end(), e.x1) - xs.begin();
            int r = lower_bound(xs.begin(), xs.end(), e.x2) - xs.begin();
            update(1, 0, n - 1, l, r, e.type);
        }

        double half = total / 2.0;
        double acc = 0.0;

        for (auto &s : strips) {
            double area = s[1] * s[2];
            if (acc + area >= half) {
                return s[0] + (half - acc) / s[2];
            }
            acc += area;
        }
        return 0.0;
    }
};
