package main

import (
    "math"
	"os"
	"time"
)

// Generate generates the logs with given options
func Generate(option *Option) error {
    frac_ns := 1e9 * (option.Created - math.Floor(option.Created))
    created := time.Unix(int64(option.Created), int64(frac_ns))
    writer := os.Stdout

    ival := time.Duration(option.Sleep)
    index := option.Index

    // Generates the logs until the certain number of lines is reached
    for line := 0; line < option.Number; line++ {
        every := DeriveEvery(index)
        log := NewLog(option.Format, created, every)
        writer.Write([]byte(log + "\n"))

        created = created.Add(ival)
        index++
    }

	return nil
}

// NewLog creates a log for given format
func NewLog(format string, t time.Time, every string) string {
	switch format {
	case "apache_common":
		return NewApacheCommonLog(t, every)
	case "apache_combined":
		return NewApacheCombinedLog(t, every)
	default:
		return ""
	}
}

func DeriveEvery(count int64) string {
    if count % 1e12 == 0 {
        return "every1t"
    } else if count % 1e11 == 0 {
        return "every1 every10 every100 every1k every10k every10b every1m every10m every100m every1b every10b every100b"
    } else if count % 1e10 == 0 {
        return "every1 every10 every100 every1k every10k every10b every1m every10m every100m every1b every10b"
    } else if count % 1e9 == 0 {
        return "every1 every10 every100 every1k every10k every100k every1m every10m every100m every1b"
    } else if count % 1e8 == 0 {
        return "every1 every10 every100 every1k every10k every100k every1m every10m every100m"
    } else if count % 1e7 == 0 {
        return "every1 every10 every100 every1k every10k every100k every1m every10m"
    } else if count % 1e6 == 0 {
        return "every1 every10 every100 every1k every10k every100k every1m"
    } else if count % 1e5 == 0 {
        return "every1 every10 every100 every1k every10k every100k"
    } else if count % 1e4 == 0 {
        return "every1 every10 every100 every1k every10k"
    } else if count % 1e3 == 0 {
        return "every1 every10 every100 every1k"
    } else if count % 1e2 == 0 {
        return "every1 every10 every100"
    } else if count % 1e1 == 0 {
        return "every1 every10"
    } else {
        return "every1"
    }
}
