package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/scoutred/opendsd"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

type CeCase struct { // Our example struct, you can use "-" to ignore a field
	ID         string `csv:"case_id"`
	CaseSource string `csv:"case_source"`
}

func main() {
	filePath := flag.String("filePath", "", "xml file path of code enforcement data")
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "usage: dsd_ce [options]")
		flag.PrintDefaults()
	}
	flag.Parse()

	f, err := os.Open(*filePath)
	check(err)

	codeEnforcement, err := opendsd.DecodeCodeEnforcementCases(f)
	check(err)

	var ceCases []CeCase
	for _, c := range codeEnforcement.Cases {
		ceCases = append(ceCases, CeCase{
			ID:         c.ID,
			CaseSource: c.CaseSource,
		})
	}

	csvContent, err := gocsv.MarshalString(ceCases)
	check(err)

	fmt.Print(string(csvContent))
}
