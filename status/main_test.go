package LibraryStatus

import (
	"testing"
)

func TestGetStatus(t *testing.T) {
	status, err := GetStatus()
	if err != nil ||
		!status.IsLibraryStatusAcceptable() {
		t.Fatal("Status got is not acceptable!", err)
	}
}
