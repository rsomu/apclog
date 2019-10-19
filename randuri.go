package main

import (
	"math/rand"
)

// RandURL generates a random url from the given list
func RandURI() string {
	uris := []string{
          "/administrator/",
          "/apache-log/access.log",
          "/apache-log/apache.gif",
          "/administrator/index.php",
          "/administrator/home.php",
          "/OrderEntry/placeOrder",
          "/OrderEntry/getOrder",
          "/OrderEntry/report/weeklyreport",
          "/OrderEntry/retrieveItem/catalog",
          "/OrderEntry/listItem/catalog",
          "/finance/annualreport/fy20",
          "/finance/annualreport/fy19",
          "/finance/annualreport/fy18",
          "/finance/analysis/revenues",
          "/finance/analysis/sales",
          "/finance/analysis/margins",
          "/inventory/replenishItem",
          "/inventory/purchaseItem",
          "/inventory/listItems",
          "/inventory/getItems",
          "/inventory/searchItems",
          "/wiki/bin/search/Main",
          "/wiki/bin/rdiff/ReadMe",
          "/wiki/bin/edit/Wiki/Service",
          "/wiki/bin/view/Main/homepage",
          "/wiki/bin/attach/Main/comments",
          "/wiki/bin/oops/Main/Error",
          "/wiki/bin/statistics/Main",
          "/search/blogs",
          "/search/solutions/oracle",
          "/search/solutions/vmware",
          "/search/solutions/windows",
          "/search/products/flasharray",
          "/search/products/flashblade",
          "/search/products/objectengine",
          "/display/products",
          "/display/pages",
          "/display/company",
          "/display/aboutus",
          "/display/contactus",
          "/en-US/app/",
          "/en-US/page/",
          "/en-US/config/autoload",
          "/en-US/app/monitoring",
          "/en-US/app/services",
	}
	return uris[rand.Intn(len(uris))]
}
