package leetcode

import (
	"fmt"
	"testing"
)

func hardestWorker(n int, logs [][]int) int {
	var delta = make([]int, len(logs))
	var pre = 0
	for i := 0; i < len(logs); i++ {
		if i != 0 {
			pre = logs[i-1][1]
		}
		delta[i] = logs[i][1] - pre
	}
	var max = delta[0]
	var minWorkerId = n - 1
	for _, v := range delta {
		if v > max {
			max = v
		}
	}
	for i := 0; i < len(delta); i++ {
		if delta[i] == max {
			if logs[i][0] <= minWorkerId {
				minWorkerId = logs[i][0]
			}
		}
	}
	return minWorkerId

}

func hardestWorker2(n int, logs [][]int) int {
	var maxWorker = logs[0][1]
	var workId = logs[0][0]
	for i := 1; i < len(logs); i++ {
		if logs[i][1]-logs[i-1][1] > maxWorker {
			maxWorker = logs[i][1] - logs[i-1][1]
			workId = logs[i][0]
		} else if logs[i][1]-logs[i-1][1] == maxWorker && logs[i][0] < workId {
			workId = logs[i][0]
		}
	}
	return workId
}

func hardestWorker3(n int, logs [][]int) int {
	var workId = logs[0][0]
	var startTime = 0
	var lastCost = 0
	for _, v := range logs {

		if v[1]-startTime > lastCost {
			lastCost = v[1] - startTime
			workId = v[0]
		} else if v[1]-startTime == lastCost && v[0] < workId {
			workId = v[0]
		}
		startTime = v[1]
	}
	return workId
}

func TestHardestWorker(t *testing.T) {
	n := 17
	var worker = [][]int{{14, 2}, {10, 4}, {1, 6}, {0, 7}, {1, 9}, {5, 10}, {13, 11}, {11, 18}, {8, 23}, {3, 27}, {10, 28}, {3, 33}, {11, 34}, {3, 38}, {1, 39}, {10, 42}, {4, 47}, {12, 54}, {0, 56}, {7, 57}, {4, 58}, {5, 67}, {10, 68}, {7, 70}, {15, 78}, {4, 79}, {13, 83}, {3, 84}, {15, 87}, {1, 94}, {5, 96}, {4, 98}, {16, 105}, {5, 107}, {11, 108}, {12, 110}, {5, 114}, {2, 116}, {6, 119}, {12, 122}, {8, 127}, {6, 129}, {2, 131}, {3, 132}, {15, 140}, {3, 144}, {5, 145}, {11, 147}, {1, 152}, {5, 153}, {14, 154}, {0, 157}, {8, 158}, {10, 160}, {7, 167}, {4, 173}, {0, 175}, {15, 177}, {16, 181}, {2, 182}, {8, 187}, {6, 188}, {0, 189}, {10, 191}, {11, 193}}
	fmt.Println(hardestWorker(n, worker))
	fmt.Println(hardestWorker2(n, worker))
	fmt.Println(hardestWorker3(n, worker))

}
