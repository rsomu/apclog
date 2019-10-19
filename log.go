package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit"
)

const (
// ApacheCommonLog :   {host} {user-identifier} {auth-user-id} [{datetime}] "{method} {request} HTTP/1.0" {response-code} {bytes} {every}
	ApacheCommonLog = "%s - %s %d [%s] \"%s %s\" %d %d %s"
// ApacheCombinedLog : {host} {user-identifier} {auth-user-id} [{datetime}] "{method} {request} HTTP/1.0" {response-code} {bytes} "{referrer}" "{agent}" {every}
	ApacheCombinedLog = "%s - %s %d [%s] \"%s %s\" %d %d \"%s\" \"%s\" %s"
)

// NewApacheCommonLog creates a log string with apache common log format
func NewApacheCommonLog(t time.Time, every string) string {
	return fmt.Sprintf(
		ApacheCommonLog,
		RandIPV4(),
		RandUser(),
		gofakeit.Number(0, 1000),
		t.Format(time.RFC3339),
		gofakeit.HTTPMethod(),
		RandResourceURI(),
		gofakeit.StatusCode(),
		gofakeit.Number(0, 30000),
	        every,
	)
}

// NewApacheCombinedLog creates a log string with apache combined log format
func NewApacheCombinedLog(t time.Time, every string) string {
	return fmt.Sprintf(
		ApacheCombinedLog,
		RandIPV4(),
		RandUser(),
		gofakeit.Number(100, 999),
		t.Format(time.RFC3339),
		gofakeit.HTTPMethod(),
		RandURI(),
		gofakeit.StatusCode(),
		gofakeit.Number(100, 20000),
		RandURL(),
		gofakeit.UserAgent(),
	        every,
	)
}
