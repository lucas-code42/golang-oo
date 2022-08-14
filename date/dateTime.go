package date

import "time"

func DateTime() string {
	dt := time.Now()
	return dt.String()
}
