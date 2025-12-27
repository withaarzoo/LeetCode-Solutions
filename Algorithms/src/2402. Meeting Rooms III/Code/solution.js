var mostBooked = function (n, meetings) {
  meetings.sort((a, b) => a[0] - b[0]);

  const freeRooms = new MinPriorityQueue({
    priority: (x) => x,
  });

  for (let i = 0; i < n; i++) {
    freeRooms.enqueue(i);
  }

  const busyRooms = new MinPriorityQueue({
    priority: (x) => x.end,
  });

  const count = Array(n).fill(0);

  for (const [start, end] of meetings) {
    const duration = end - start;

    while (!busyRooms.isEmpty() && busyRooms.front().element.end <= start) {
      freeRooms.enqueue(busyRooms.dequeue().element.room);
    }

    if (!freeRooms.isEmpty()) {
      const room = freeRooms.dequeue().element;
      busyRooms.enqueue({ end: end, room: room });
      count[room]++;
    } else {
      const { end: finish, room } = busyRooms.dequeue().element;
      busyRooms.enqueue({ end: finish + duration, room: room });
      count[room]++;
    }
  }

  let ans = 0;
  for (let i = 1; i < n; i++) {
    if (count[i] > count[ans]) ans = i;
  }
  return ans;
};
