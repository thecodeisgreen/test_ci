package version

import "time"

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

	if date == "__DATE__" {
		date = time.Now().String()
	} else {
		t, _ := time.Parse(time.ANSIC, date)
		date = t.String()
	}

	return Version{
		Value: value,
		Date:  date,
	}
}

func (version Version) String() string {
	return version.Value + " " + version.Date
}
