package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

const version = "0.1.0"
const usage = `apclog is a fake log generator for common log formats

Usage: apclog [options]

Version: %s

Options:
  -f, --format string      choose log format. ("apache_common"|"apache_combined") (default "apache_common")
  -n, --number integer     number of events to generate.
  -c, --created numeric    creation start time for each log (in seconds since epoch).
  -s, --sleep numeric      creation time interval for each log (in nanoseconds). It does not actually sleep.
  -i, --index integer      initial index (count) in the overall log sequence
`

var validFormats = []string{"apache_common", "apache_combined"}

// Option defines log generator options
type Option struct {
	Format    string
	Number    int
	Created   float64
	Sleep     int64
    Index     int64
}

func init() {
	pflag.Usage = printUsage
}

func printUsage() {
	fmt.Printf(usage, version)
}

func printVersion() {
	fmt.Printf("apclog version %s\n", version)
}

func errorExit(err error) {
	os.Stderr.WriteString(err.Error() + "\n")
	os.Exit(-1)
}

func defaultOptions() *Option {
	return &Option{
		Format:    "apache_common",
		Number:    1000,
        Created:   0.0,
		Sleep:     0,
        Index:     1,
	}
}

// ParseFormat validates the given format
func ParseFormat(format string) (string, error) {
	if !containString(validFormats, format) {
		return "", fmt.Errorf("%s is not a valid format", format)
	}
	return format, nil
}

// ParseNumber validates the given number
func ParseNumber(lines int) (int, error) {
	if lines < 0 {
		return 0, errors.New("lines can not be negative")
	}
	return lines, nil
}

// ParseCreated validates the given start time
func ParseCreated(created float64) (float64, error) {
	if created < 0 {
		return 0.0, errors.New("created can not be negative")
	}
	return created, nil
}

// ParseSleep validates the given sleep
func ParseSleep(sleep int64) (int64, error) {
	if sleep < 0 {
		return 0, errors.New("sleep can not be negative")
	}
	return sleep, nil
}

// ParseIndex validates the given number
func ParseIndex(index int64) (int64, error) {
	if index < 0 {
		return 0, errors.New("start index can not be negative")
	}
	return index, nil
}


// ParseOptions parses given parameters from command line
func ParseOptions() *Option {
	var err error

	opts := defaultOptions()

	help := pflag.BoolP("help", "h", false, "Show usage")
	version := pflag.BoolP("version", "v", false, "Show version")
	format := pflag.StringP("format", "f", opts.Format, "Log format")
	number := pflag.IntP("number", "n", opts.Number, "Number of lines to generate")
	created := pflag.Float64P("created", "c", opts.Created, "Creation start time for each log (in seconds since epoch)")
	sleep := pflag.Int64P("sleep", "s", opts.Sleep, "Creation time interval for each log (in nanoseconds)")
	index := pflag.Int64P("index", "i", opts.Index, "Initial index (count) in the overall log sequence")

	pflag.Parse()

	if *help {
		printUsage()
		os.Exit(0)
	}
	if *version {
		printVersion()
		os.Exit(0)
	}
	if opts.Format, err = ParseFormat(*format); err != nil {
		errorExit(err)
	}
	if opts.Number, err = ParseNumber(*number); err != nil {
		errorExit(err)
	}
	if opts.Created, err = ParseCreated(*created); err != nil {
		errorExit(err)
	}
	if opts.Sleep, err = ParseSleep(*sleep); err != nil {
		errorExit(err)
	}
	if opts.Index, err = ParseIndex(*index); err != nil {
		errorExit(err)
	}
	return opts
}
