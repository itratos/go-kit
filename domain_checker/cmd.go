package domain

import (
	"net"
	"strings"
)

type Info struct {
	Mx    []*net.MX
	Dmarc string
	Spf   string
}

func CheckDom(domainStr string) (Info, error) {
	var domainInfo Info
	var err error
	domainInfo.Mx, err = net.LookupMX(domainStr)

	txtRecords, err := net.LookupTXT(domainStr)

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf") {
			domainInfo.Spf = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domainStr)

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			domainInfo.Dmarc = record
			break
		}
	}

	return domainInfo, err
}
