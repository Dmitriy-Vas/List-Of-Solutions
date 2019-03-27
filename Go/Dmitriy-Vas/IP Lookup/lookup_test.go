package IP_Lookup

import "testing"

var ip = []struct{
	ip string
	out IP
	country string
}{
	{"1.1.1.1", IP{
		IP: "1.1.1.1",
		HostName: "one.one.one.one",
		City: "",
		Region: "",
		Country: "AU",
		Loc: "-33.4940,143.2100",
	}, "Australia"},
	{"8.8.8.8", IP{
		IP: "8.8.8.8",
		HostName: "google-public-dns-a.google.com",
		City: "Mountain View",
		Region: "California",
		Country: "US",
		Loc: "37.3860,-122.0840",
	}, "United States of America"},
}

func TestLookup(t *testing.T) {
	for _, d := range ip {
		if IPInfo(d.ip) != d.out {
			t.Fatal(d)
		}
		if CountryInfo(d.ip).Name != d.country {
			t.Fatal(d)
		}
	}
}
