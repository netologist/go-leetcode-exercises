# Meeting Rooms II (Minimum Concurrent Server / Resource Allocation)

## LeetCode Information
- **Number:** 253
- **Difficulty:** Medium
- **FinTech Companies:** Bloomberg, Stripe, Citadel, Two Sigma, Robinhood, Brex

---

## FinTech Relevance & Real-World Application
Frequently asked in FinTech systems design & live coding as **"Peak Concurrency Provisioning"**:
1. **Payment Gateway Worker Allocation:** Calculating the minimum number of isolated background worker goroutines required to process overlapping transaction spans without blocking.
2. **Matching Engine Core Utilization:** Sizing multi-tenant financial auction queues to guarantee zero latency starvation during burst volume.
3. **Infrastructure Cost Optimization:** Computing true peak simultaneous resource usage to configure autoscaling boundaries for cloud computing instances.

---

## Algorithmic Approach
While this can be solved using a Min-Heap, the **Two-Pointer Chronological Sweep** approach is more cache-friendly and faster in Go:

1. Extract `starts` and `ends` into separate slices and sort both ascending in $O(N \log N)$.
2. Iterate `startIdx` through all intervals:
   - If `starts[startIdx] < ends[endIdx]`: A new transaction begins before the earliest ongoing transaction completes. We must allocate a new room/server (`rooms++`).
   - Else: The earliest transaction has finished, freeing its room for reuse (`endIdx++`).

---

## Complexity Analysis
- **Time Complexity:** $O(N \log N)$ - Dominated by sorting the `starts` and `ends` integer slices.
- **Space Complexity:** $O(N)$ - Auxiliary arrays for start and end timestamps.

---

## Critical Edge Cases
- Exact boundary equality: If interval A ends at time `T` and interval B starts at time `T` (e.g. `[1, 5]` and `[5, 10]`), the resource is freed and reused immediately; no extra allocation occurs.
- Completely non-overlapping intervals (yields `1`).
- Empty interval set (yields `0`).

---

## 💻 Production Go Implementation

```go
package meetingrooms

import "sort"

// MinMeetingRooms calculates the minimum number of conference rooms (or concurrent worker servers)
// required to accommodate all scheduled intervals.
// Time Complexity: O(N log N), Space Complexity: O(N)
func MinMeetingRooms(intervals [][]int) int {
    if len(intervals) == 0 {
        return 0
    }

    n := len(intervals)
    starts := make([]int, n)
    ends := make([]int, n)

    for i, interval := range intervals {
        starts[i] = interval[0]
        ends[i] = interval[1]
    }

    sort.Ints(starts)
    sort.Ints(ends)

    rooms := 0
    endIdx := 0

    for startIdx := range n {
        if starts[startIdx] < ends[endIdx] {
            // A new meeting starts before the earliest ending meeting finishes -> need a new room
            rooms++
        } else {
            // The earliest meeting ended -> room is reused
            endIdx++
        }
    }

    return rooms
}
```

---

## 🧪 Running Unit Tests

```bash
go test -v ./07-meeting-rooms-ii/...
go test -race ./07-meeting-rooms-ii/...
```

---

[⬅️ Back to All Exercises](../README.md)
