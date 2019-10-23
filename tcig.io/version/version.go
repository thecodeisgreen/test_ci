package version

import (
	"time"
)

type Version struct {
	Value string
	Date  string
}

func Get() Version {
	value := "__VERSION__"
	date := "__DATE__"

	if value == "" || value == "__VERSION__" {
		value = "latest"
	}

	t, err := time.Parse(time.UnixDate, date)
	if err != nil {
		date = time.Now().Format(time.UnixDate)
	} else {
		date = t.Local().Format(time.UnixDate)
	}

	return Version{
		Value: value,
		Date:  date,
	}
}

func (version Version) String() string {
	return version.Value + " " + version.Date
}
