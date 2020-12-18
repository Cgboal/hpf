package main

import (
	"bufio"
	"os"
	"github.com/Cgboal/DomainParser"
	"fmt"
	"sort"
	"flag"
)

type Domain struct {
	FQDN string
	Subdomain string
}

type FrequencyPair struct {
	Key string
	Value int
}

type FrequencyPairSlice []FrequencyPair

func (p FrequencyPairSlice) Len() int           { return len(p) }
func (p FrequencyPairSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p FrequencyPairSlice) Less(i, j int) bool { return p[i].Value < p[j].Value }

func DomainParser() func(string) Domain {
	parser := parser.NewDomainParser()
	return func(domain string) Domain {
		fqdn := parser.GetFQDN(domain)
		subdomain := parser.GetSubdomain(domain)
		return Domain{
			FQDN: fqdn,
			Subdomain: subdomain,
		}
	}
}

func buildFrequencyMap(domains []Domain) map[string]int {
	frequencyMap := make(map[string]int)
	for _, domain := range domains {
		if _, ok := frequencyMap[domain.FQDN]; ok {
			frequencyMap[domain.FQDN] = frequencyMap[domain.FQDN] + 1
		} else {
			frequencyMap[domain.FQDN] = 1
		}
	}

	return frequencyMap

}

func buildSortedFrequencySlice(frequencyMap map[string]int) FrequencyPairSlice {
	frequencySlice := make(FrequencyPairSlice, len(frequencyMap))
	i := 0
	for k, v := range frequencyMap {
		frequencySlice[i] = FrequencyPair{k, v}
		i++
	}
	sort.Sort(frequencySlice)
	return frequencySlice
}

func filterAndPrint(frequencyMap map[string]int, domains []Domain, filter int) {
	for _, domain := range domains {
		if frequencyMap[domain.FQDN] <= filter {
			if domain.Subdomain == "" {
				fmt.Printf("%s\n", domain.FQDN)
			} else {
				fmt.Printf("%s.%s\n", domain.Subdomain, domain.FQDN)
			}
		}
	}
}


func main () {
	var highPass = flag.Int("f", 0, "Frequency cut off, FQDNs which appear more than this many times will be excluded")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	domains := []Domain{}
	parse := DomainParser()
	for scanner.Scan() {
		domains = append(domains, parse(scanner.Text()))
	}

	frequencyMap := buildFrequencyMap(domains)

	if *highPass != 0 {
		filterAndPrint(frequencyMap, domains, *highPass)
	} else {
		frequencySlice := buildSortedFrequencySlice(frequencyMap)
		for _, f := range frequencySlice {
			fmt.Printf("%d: %s\n", f.Value, f.Key)
		}
	}
}
