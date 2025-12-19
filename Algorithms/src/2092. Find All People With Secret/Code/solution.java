class Solution {
    int[] parent;

    int find(int x) {
        if (parent[x] == x)
            return x;
        return parent[x] = find(parent[x]);
    }

    void union(int x, int y) {
        x = find(x);
        y = find(y);
        if (x != y)
            parent[y] = x;
    }

    public List<Integer> findAllPeople(int n, int[][] meetings, int firstPerson) {
        Arrays.sort(meetings, (a, b) -> a[2] - b[2]);

        parent = new int[n];
        for (int i = 0; i < n; i++)
            parent[i] = i;

        boolean[] knows = new boolean[n];
        knows[0] = knows[firstPerson] = true;

        int i = 0;
        while (i < meetings.length) {
            int time = meetings[i][2];
            List<Integer> people = new ArrayList<>();

            int j = i;
            while (j < meetings.length && meetings[j][2] == time) {
                union(meetings[j][0], meetings[j][1]);
                people.add(meetings[j][0]);
                people.add(meetings[j][1]);
                j++;
            }

            Set<Integer> good = new HashSet<>();
            for (int p : people) {
                if (knows[p])
                    good.add(find(p));
            }

            for (int p : people) {
                if (good.contains(find(p))) {
                    knows[p] = true;
                } else {
                    parent[p] = p;
                }
            }
            i = j;
        }

        List<Integer> ans = new ArrayList<>();
        for (i = 0; i < n; i++)
            if (knows[i])
                ans.add(i);
        return ans;
    }
}
