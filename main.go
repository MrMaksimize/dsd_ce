package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/gocarina/gocsv"
	"github.com/scoutred/opendsd"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
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

	dat, err := ioutil.ReadFile(*filePath)
	check(err)

	buf := bytes.NewBufferString(string(dat))
	cases, err := opendsd.DecodeCodeEnforcementCases(buf)
	check(err)
	//fmt.Printf("%+v\n", cases)
	ce_cases := cases.Cases
	spew.Dump(ce_cases[0])
	fmt.Println(reflect.TypeOf(ce_cases))
	ce_cases2 := CeCase(ce_cases)
	csvContent, err := gocsv.MarshalString(&ce_cases2)
	check(err)
	fmt.Print(string(csvContent))
	//fmt.Println("fpath: ", *filePath)

	/*project, err := client.ProjectByID(*projectID)
	    if err != nil {
		log.Fatal(err)
	    }
	    //log.Printf("project %v", project)

	    b, err := json.Marshal(project)
	    if err != nil {
		log.Println("error:", err)
	    }
	    os.Stdout.Write(b)*/
}
