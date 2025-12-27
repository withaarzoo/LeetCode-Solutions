import (
	"container/heap"
	"sort"
)

// ---------- Free Rooms Heap (min room number) ----------
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// ---------- Busy Rooms Heap (min end time) ----------
type Room struct {
	end  int
	room int
}

type RoomHeap []Room

func (h RoomHeap) Len() int { return len(h) }
func (h RoomHeap) Less(i, j int) bool {
	if h[i].end == h[j].end {
		return h[i].room < h[j].room
	}
	return h[i].end < h[j].end
}
func (h RoomHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *RoomHeap) Push(x interface{}) {
	*h = append(*h, x.(Room))
}

func (h *RoomHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// ---------- Main Logic ----------
func mostBooked(n int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	freeRooms := &IntHeap{}
	heap.Init(freeRooms)
	for i := 0; i < n; i++ {
		heap.Push(freeRooms, i)
	}

	busyRooms := &RoomHeap{}
	heap.Init(busyRooms)

	count := make([]int, n)

	for _, m := range meetings {
		start, end := m[0], m[1]
		duration := end - start

		for busyRooms.Len() > 0 && (*busyRooms)[0].end <= start {
			room := heap.Pop(busyRooms).(Room).room
			heap.Push(freeRooms, room)
		}

		if freeRooms.Len() > 0 {
			room := heap.Pop(freeRooms).(int)
			heap.Push(busyRooms, Room{end, room})
			count[room]++
		} else {
			r := heap.Pop(busyRooms).(Room)
			heap.Push(busyRooms, Room{r.end + duration, r.room})
			count[r.room]++
		}
	}

	ans := 0
	for i := 1; i < n; i++ {
		if count[i] > count[ans] {
			ans = i
		}
	}
	return ans
}
