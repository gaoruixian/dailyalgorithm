package leetcode

import (
	"fmt"
	"testing"
)

func haveConflict(event1 []string, event2 []string) bool {
	return !(event2[0] > event1[1] || event2[1] < event1[0])
}
func TestHaveConflict(t *testing.T) {
	event1 := []string{"14:00", "15:00"}
	event2 := []string{"10:00", "11:00"}
	fmt.Println(haveConflict(event1, event2))
}
