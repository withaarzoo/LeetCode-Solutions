class Solution {
    static class Event {
        double y, x1, x2;
        int type;
        Event(double y, double x1, double x2, int type) {
            this.y = y; this.x1 = x1; this.x2 = x2; this.type = type;
        }
    }

    double[] xs, len;
    int[] cover;

    void update(int node, int l, int r, int ql, int qr, int val) {
        if (qr <= l || r <= ql) return;
        if (ql <= l && r <= qr) {
            cover[node] += val;
        } else {
            int m = (l + r) / 2;
            update(node*2, l, m, ql, qr, val);
            update(node*2+1, m, r, ql, qr, val);
        }

        if (cover[node] > 0) {
            len[node] = xs[r] - xs[l];
        } else if (r - l == 1) {
            len[node] = 0;
        } else {
            len[node] = len[node*2] + len[node*2+1];
        }
    }

    public double separateSquares(int[][] squares) {
        ArrayList<Double> xList = new ArrayList<>();
        for (int[] s : squares) {
            xList.add((double)s[0]);
            xList.add(s[0] + s[2] * 1.0);
        }
        xs = xList.stream().distinct().sorted().mapToDouble(d->d).toArray();

        ArrayList<Event> events = new ArrayList<>();
        for (int[] s : squares) {
            events.add(new Event(s[1], s[0], s[0]+s[2], 1));
            events.add(new Event(s[1]+s[2], s[0], s[0]+s[2], -1));
        }
        events.sort((a,b)->Double.compare(a.y,b.y));

        int n = xs.length;
        cover = new int[4*n];
        len = new double[4*n];

        double total = 0, prevY = events.get(0).y;
        ArrayList<double[]> strips = new ArrayList<>();

        for (Event e : events) {
            if (e.y > prevY) {
                double w = len[1];
                double h = e.y - prevY;
                total += w * h;
                strips.add(new double[]{prevY, h, w});
                prevY = e.y;
            }
            int l = Arrays.binarySearch(xs, e.x1);
            int r = Arrays.binarySearch(xs, e.x2);
            update(1, 0, n-1, l, r, e.type);
        }

        double half = total / 2, acc = 0;
        for (double[] s : strips) {
            double area = s[1] * s[2];
            if (acc + area >= half) {
                return s[0] + (half - acc) / s[2];
            }
            acc += area;
        }
        return 0;
    }
}
