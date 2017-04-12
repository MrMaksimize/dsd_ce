package dsd_ce

// Convert from dsd xml to csv for code enforcement

import (
	"flag"
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/scoutred/opendsd"
	"log"
	"os"
)


// Function that checks for errors.
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Struct for code enforcement case.
type CeCase struct { // Our example struct, you can use "-" to ignore a field
	ID                       string  `csv:"case_id"`
	APN                      string  `csv:"apn"`
	StreetAddress            string  `csv:"street_address"`
	CaseSource               string  `csv:"case_source"`
	Description              string  `csv:"description"`
	OpenDate                 string  `csv:"open_date"`
	CloseDate                string  `csv:"close_date"`
	CloseReason              string  `csv:"close_reason"`
	CloseNote                string  `csv:"close_note"`
	Lat                      float64 `csv:"latitude"`
	Lon                      float64 `csv:"longitude"`
	NAD83Northing            string  `csv:"nad83_northing"`
	NAD83Easting             string  `csv:"nad83_easting"`
	Workgroup                string  `csv:"workgroup"`
	InvestigatorName         string  `csv:"investigator_name"`
	InvestigatorPhoneNumber  string  `csv:"investigator_phone_number"`
	InvestigatorEmailAddress string  `csv:"investigator_email_address"`
	InvestigatorActive       string  `csv:"investigator_active"`
	LastAction               string  `csv:"last_action"`
	LastActionDueDate        string  `csv:"last_action_due_date"`
	RemedyMsg                string  `csv:"remedy_msg"`
}

// Complaint Struct
type CeComplaint struct {
	CaseID string `csv:"case_id"`
	TypeID string `csv:"complaint_type_id"`
	Type   string `csv:"complaint_type"`
}

// Main function
func main() {
	xmlPath := flag.String("xmlPath", "", "xml file path of code enforcement data")
	ceOutPath := flag.String("ceOutPath", "", "csv file to write cases to")
	cmplOutPath := flag.String("cmplOutPath", "", "csv file to write complaints to")
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "usage: dsd_ce [options]")
		flag.PrintDefaults()
	}
	flag.Parse()

	f, err := os.Open(*xmlPath)
	check(err)

	codeEnforcement, err := opendsd.DecodeCodeEnforcementCases(f)
	check(err)

	var ceCases []CeCase
	var ceComplaints []CeComplaint
	for _, c := range codeEnforcement.Cases {
		ceCases = append(ceCases, CeCase{
			ID:                       c.ID,
			CaseSource:               c.CaseSource,
			Description:              c.Description,
			OpenDate:                 c.OpenDate,
			CloseDate:                c.CloseDate,
			CloseReason:              c.CloseReason,
			CloseNote:                c.CloseNote,
			APN:                      c.APN,
			StreetAddress:            c.StreetAddress,
			Lat:                      c.Lat,
			Lon:                      c.Lon,
			NAD83Northing:            c.NAD83Northing,
			NAD83Easting:             c.NAD83Easting,
			Workgroup:                c.Workgroup,
			InvestigatorName:         c.InvestigatorName,
			InvestigatorPhoneNumber:  c.InvestigatorPhoneNumber,
			InvestigatorEmailAddress: c.InvestigatorEmailAddress,
			InvestigatorActive:       c.InvestigatorActive,
			LastAction:               c.LastAction,
			LastActionDueDate:        c.LastActionDueDate,
			RemedyMsg:                c.RemedyMsg,
		})
		for _, compl := range c.Complaints {
			ceComplaints = append(ceComplaints, CeComplaint{
				CaseID: c.ID,
				TypeID: compl.TypeID,
				Type:   compl.Type,
			})
		}
	}

	/*csvContent, err := gocsv.MarshalString(ceComplaints)
	check(err)
	fmt.Print(string(csvContent))*/
	// Write CSV
	ceFile, err := os.OpenFile(*ceOutPath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	check(err)
	cmplFile, err := os.OpenFile(*cmplOutPath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	check(err)
	defer ceFile.Close()
	defer cmplFile.Close()
	err = gocsv.MarshalFile(ceCases, ceFile)
	check(err)
	err = gocsv.MarshalFile(ceComplaints, cmplFile)
	check(err)
}
