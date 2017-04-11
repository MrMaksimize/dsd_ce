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
	ID                       string      `csv:"case_id"`
	CaseSource 		 string      `csv:"case_source"`
	Description              string      `csv:"description"`
	OpenDate                 string      `csv:"open_date"`
        CloseDate                string      `csv:"close_date"`
	CloseReason              string      `csv:"close_reason"`
	CloseNote                string      `csv:"close_note"`
	APN                      string      `csv:"apn"`
	StreetAddress            string      `csv:"street_address"`
	SortableStreetAddress    string      `csv:"sortable_street_address"`
	MapReference             string      `csv:"map_reference"`
	Lat                      float64     `csv:"latitude"`
	Lon                      float64     `csv:"longitude"`
	NAD83Northing            interface{} `csv:"nad83_northing"`
	NAD83Easting             interface{} `csv:"nad83_easting"`
	Workgroup                string      `csv:"workgroup"`
	InvestigatorName         string      `csv:"investigator_name"`
	InvestigatorPhoneNumber  string      `csv:"investigator_phone_number"`
	InvestigatorEmailAddress string      `csv:"investigator_email_address"`
	InvestigatorActive       string      `csv:"investigator_active"`
	LastAction               string      `csv:"last_action"`
	LastActionDueDate        string      `csv:"last_action_due_date"`
	RemedyMsg                string      `csv:"remedy_msg"`
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
			Description: c.Description,
			OpenDate: c.OpenDate,
			CloseDate: c.CloseDate,
			CloseReason: c.CloseReason,
			CloseNote: c.CloseNote,
			APN: c.APN,
			StreetAddress: c.StreetAddress,
			SortableStreetAddress: c.SortableStreetAddress,
			MapReference: c.MapReference,
			Lat: c.Lat,
			Lon: c.Lon,
			NAD83Northing: c.NAD83Northing,
			NAD83Easting: c.NAD83Easting,
			Workgroup: c.Workgroup,
			InvestigatorName: c.InvestigatorName,
			InvestigatorPhoneNumber: c.InvestigatorPhoneNumber,
			InvestigatorEmailAddress: c.InvestigatorEmailAddress,
			InvestigatorActive: c.InvestigatorActive,
			LastAction: c.LastAction,
			LastActionDueDate: c.LastActionDueDate,
			RemedyMsg: c.RemedyMsg,
		})
	}

	csvContent, err := gocsv.MarshalString(ceCases)
	check(err)

	fmt.Print(string(csvContent))
}
