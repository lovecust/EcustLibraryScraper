package LibraryStatus

import (
	"github.com/lovecust/EcustLibraryScraper/status/constants"
	"encoding/json"
)

// Library status entity.
type Status struct {
	// Current students traffic.
	Today     int64 `json:"today"`
	Available []int64`json:"available"`
	Occupied  []int64`json:"occupied"`
	// The Unix time in seconds.
	Seconds int64`json:"seconds"`
}

// Check whether the library status parsed is acceptable logically.
func (m Status) IsLibraryStatusAcceptable() bool {
	if 0 > m.Today ||
		constants.FLOORS_COUNT != len(m.Available) ||
		constants.FLOORS_COUNT != len(m.Occupied) ||
		0 > m.Seconds {
		return false
	}
	for i := 0; i < constants.FLOORS_COUNT; i++ {
		if m.Available[i]+m.Occupied[i] != constants.SEATS_OF_EACH_FLOOR[i] {
			return false
		}
	}
	return true
}

// To JSON string.
func (m Status) ToJson() string {
	str, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return string(str)
}
