package podtime

import (
	"bou.ke/monkey"
	"testing"
	"time"
)

func getTestDate() time.Time {
	layout := "2006-01-02T15:04:05.000Z"
	str := "2014-11-12T11:45:26.371Z"
	t, _ := time.Parse(layout, str)
	return t
}

func TestCurrentTileDay(t *testing.T) {
	patch := monkey.Patch(time.Now, func() time.Time { return getTestDate() })
	expected := "1d-16386"
	result := CurrentTile("1d")
	if expected != result {
		t.Errorf("CurrentTitle logic Failed")
	}

	result = CurrentTile("7d")
	if "7d-2340" != result {
		t.Errorf("CurrentTitle logic for 7d")
	}

	defer patch.Unpatch()
}

func TestCurrentTileMonth(t *testing.T) {
	patch := monkey.Patch(time.Now, func() time.Time { return getTestDate() })
	result := CurrentTile("1m")
	if "1m-23596545" != result {
		t.Errorf("CurrentTitle logic Failed for 1m")
	}

	result = CurrentTile("30m")
	if "30m-786551" != result {
		t.Errorf("CurrentTitle logic for 30m")
	}

	defer patch.Unpatch()
}

func TestCurrentTileHour(t *testing.T) {
	patch := monkey.Patch(time.Now, func() time.Time { return getTestDate() })
	result := CurrentTile("1h")
	if "1h-393275" != result {
		t.Errorf("CurrentTitle logic Failed for 1m")
	}

	result = CurrentTile("30h")
	if "30h-13109" != result {
		t.Errorf("CurrentTitle logic for 30h")
	}

	defer patch.Unpatch()
}

func TestCurrentTileSecond(t *testing.T) {
	patch := monkey.Patch(time.Now, func() time.Time { return getTestDate() })
	result := CurrentTile("1s")
	if "1s-1415792726" != result {
		t.Errorf("CurrentTitle logic Failed for 1m")
	}

	result = CurrentTile("30s")
	if "30s-47193090" != result {
		t.Errorf("CurrentTitle logic for 30h")
	}

	defer patch.Unpatch()
}

// Ideally this pulls from pondtime.info
//func TestTodayData(t *testing.T) {
//	oneDay := CurrentTile("1d")
//	sevenDay := CurrentTile("7d")
//	nintyDay := CurrentTile("90d")
//
//	if oneDay != "1d-18366" {
//		t.Errorf("Todays' Day test failed")
//	}
//	if sevenDay!= "7d-2623" {
//		t.Errorf("Todays' 7 Day test failed")
//	}
//
//	if nintyDay != "90d-204" {
//		t.Errorf("Todays' 90 Day test failed")
//	}
//
//}

func TestTimeRange(t *testing.T) {
	patch := monkey.Patch(time.Now, func() time.Time { return getTestDate() })
	begin, end := GetTimeRange("1d")
	if begin != 1415705426371 {
		t.Errorf("Begin Time is invalid")
	}
	if end != 1415791826371 {
		t.Errorf("End Time is invalid")
	}
	defer patch.Unpatch()
}
