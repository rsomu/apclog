package main

import (
	"math/rand"
)

// RandURL generates a random url from the given list
func RandURL() string {
	urls := []string{
          "dev.purestorage.com",
          "benchmark.purestorage.com",
          "wiki.purestorage.com",
          "jira.purestorage.com",
          "jenkins.purestorage.com",
          "splunk.purestorage.com",
          "blog.purestorage.com",
          "webvpn.puretec.purestorage.com",
          "esx01.puretec.purestorage.com",
          "esx02.puretec.purestorage.com",
          "fa-g31.puretec.purestorage.com",
          "fb-a02.puretec.purestorage.com",
          "fb-g28.purestorage.com",
          "ad02.puretec.purestorage.com",
          "ad01.puretec.purestorage.com",
          "dns1.purestorage.com",
          "dns2.purestorage.com",
          "ntp1.purestorage.com",
          "ntp2.purestorage.com",
          "legal.purestorage.com",
          "sales.purestorage.com",
          "support.purestorage.com",
          "hub.purestorage.com",
          "elastic.purestorage.com",
          "sf.purestorage.com",
          "purestorage.okta.com",
          "pstg.salesforce.com",
          "www.purestorage.com",
          "splunk-cm01.pstg.com",
          "splunk-sh01.pstg.com",
          "docs.purestorage.com",
          "help.storreduce.com",
          "support.storreduce.com",
          "cloud.storreduce.com",
          "www.compuverde.com",
          "support.compuverde.com",
          "support.oracle.com",
          "www.oracle.com",
          "metalink.oracle.com",
          "help.compuverde.com",
          "cloud.purestorage.com",
	}
	return urls[rand.Intn(len(urls))]
}
