class Solution:
    def separateSquares(self, squares):
        xs = []
        events = []

        for x, y, l in squares:
            xs.append(x)
            xs.append(x + l)
            events.append((y, 1, x, x + l))
            events.append((y + l, -1, x, x + l))

        xs = sorted(set(xs))
        events.sort()

        idx = {v: i for i, v in enumerate(xs)}
        n = len(xs)

        cnt = [0] * (4 * n)
        seg = [0.0] * (4 * n)

        def update(node, l, r, ql, qr, val):
            if qr <= l or r <= ql:
                return
            if ql <= l and r <= qr:
                cnt[node] += val
            else:
                m = (l + r) // 2
                update(node*2, l, m, ql, qr, val)
                update(node*2+1, m, r, ql, qr, val)

            if cnt[node] > 0:
                seg[node] = xs[r] - xs[l]
            elif r - l == 1:
                seg[node] = 0.0
            else:
                seg[node] = seg[node*2] + seg[node*2+1]

        strips = []
        total = 0.0
        prevY = events[0][0]

        for y, t, x1, x2 in events:
            if y > prevY:
                w = seg[1]
                h = y - prevY
                total += w * h
                strips.append((prevY, h, w))
                prevY = y

            update(1, 0, n - 1, idx[x1], idx[x2], t)

        half = total / 2.0
        acc = 0.0

        for y, h, w in strips:
            if acc + h * w >= half:
                return y + (half - acc) / w
            acc += h * w

        return 0.0
