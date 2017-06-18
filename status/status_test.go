package LibraryStatus

import (
	"testing"
)

func TestStatus_IsLibraryStatusAcceptable(t *testing.T) {
	TestStatus_ToJson(t)
}

func TestStatus_ToJson(t *testing.T) {
	status := Status{Today: 5493, Available: []int64{17, 144, 169, 140, 157, 8}, Occupied: []int64{19, 178, 292, 232, 192, 48}, Seconds: 1497793011}
	if !status.IsLibraryStatusAcceptable() ||
		status.ToJson() != `{"today":5493,"available":[17,144,169,140,157,8],"occupied":[19,178,292,232,192,48],"seconds":1497793011}` {
		panic("Status JSON is unexpected.")
	}
}
