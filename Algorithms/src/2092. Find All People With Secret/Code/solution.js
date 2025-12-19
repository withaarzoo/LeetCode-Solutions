var findAllPeople = function (n, meetings, firstPerson) {
  meetings.sort((a, b) => a[2] - b[2]);

  const parent = Array.from({ length: n }, (_, i) => i);
  const knows = Array(n).fill(false);
  knows[0] = knows[firstPerson] = true;

  const find = (x) => {
    if (parent[x] !== x) parent[x] = find(parent[x]);
    return parent[x];
  };

  const union = (x, y) => {
    x = find(x);
    y = find(y);
    if (x !== y) parent[y] = x;
  };

  let i = 0;
  while (i < meetings.length) {
    let time = meetings[i][2];
    let people = [];

    let j = i;
    while (j < meetings.length && meetings[j][2] === time) {
      union(meetings[j][0], meetings[j][1]);
      people.push(meetings[j][0], meetings[j][1]);
      j++;
    }

    const good = new Set();
    for (let p of people) {
      if (knows[p]) good.add(find(p));
    }

    for (let p of people) {
      if (good.has(find(p))) {
        knows[p] = true;
      } else {
        parent[p] = p;
      }
    }
    i = j;
  }

  return knows.map((v, i) => (v ? i : -1)).filter((v) => v !== -1);
};
