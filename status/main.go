package LibraryStatus

import (
	"github.com/lovecust/EcustLibraryScraper/status/constants"
	"strconv"
	"github.com/PuerkitoBio/goquery"
	"errors"
	"time"
)

// GetStatus returns the current status of Ecust Library, which includes the seats status and daily traffic.
func GetStatus() (Status, error) {
	status := Status{}
	doc, err := goquery.NewDocument(constants.URL_LIBRARY_STATUS)
	status.Seconds = time.Now().Unix()
	if err != nil {
		return status, err
	}
	today := doc.Find("#LabelAll").Text()
	status.Today, err = strconv.ParseInt(today, 10, 64)
	if err != nil {
		return status, err
	}
	for i := 0; i < constants.FLOORS_COUNT; i++ {
		iT, err := strconv.ParseInt(doc.Find("#Label" + strconv.Itoa(i+1) + "f1").Text(), 10, 64)
		status.Occupied = append(status.Occupied, iT)
		if err != nil {
			return status, err
		}
		iT, err = strconv.ParseInt(doc.Find("#Label" + strconv.Itoa(i+1) + "f2").Text(), 10, 64)
		if err != nil {
			return status, err
		}
		status.Available = append(status.Available, iT)
	}
	if !status.IsLibraryStatusAcceptable() {
		return status, errors.New("Status got is not expected.")
	}
	return status, nil
}
