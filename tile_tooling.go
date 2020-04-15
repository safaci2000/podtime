package podtime

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

type DurationType struct {
	Label  string
	Length int
}

var DurationMapping = map[string]DurationType{
	"s": {Label: "seconds", Length: 1},
	"m": {Label: "minutes", Length: 60},
	"h":   {Label: "hours", Length: 60 * 60},
	"d":    {Label: "Days", Length: 60 * 60 * 24},
}

//Time helper methods
// For some reason goLang has no UnixMilis so definiting helper function here
func  unixMilis(now time.Time) int64 {
	unixNano := now.UnixNano()
	umillisec := unixNano / 1000000
	return umillisec
}

func GetTimeRange(unit string) (int64, int64) {
	delta, _  := WindowDuration("15m")
	unitTime, _ := WindowDuration(unit)
	end := unixMilis(time.Now()) - delta
	begin := end - unitTime
	return begin, end
}


func WindowDuration(w string) (int64, error) {
	// window should be two parts, a number and a letter if it's a
	// range based index, e.g "1h".
	r, _ := regexp.Compile("([0-9]+)([smhd])")
	parts := r.FindStringSubmatch(w)
	if len(parts) >= 3 {
		num, _ := strconv.Atoi(parts[1])
		unit := parts[2]
		return (int64)(num * DurationMapping[unit].Length * 1000), nil
	}

	return -1, errors.New("Invalid ")
}


// Returns the current Title for a given string.  Example would be:
// 7d, 90d, 1d, 1m etc.
func CurrentTile(rawString string) string {
	d := time.Now()
	duration, _  := WindowDuration(rawString)
	unixMilis  := unixMilis(d)
	i := unixMilis / duration
	return fmt.Sprintf("%s-%d", rawString, i)
}

