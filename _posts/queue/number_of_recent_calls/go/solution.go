package number_of_recent_calls

type RecentCounter struct {
	calls []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		calls: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.calls = append(this.calls, t)

	for this.calls[0] < t-3000 {
		this.calls = this.calls[1:]
	}

	return len(this.calls)
}
