package main

import (
	"flag"
	"fmt"
	"os"

	spcgo "github.com/mttchpmn07/csv2spc/pkg/spcgo"
)

type file struct {
	bin  *[]byte
	path *string
}

func checkEmpty(s string) string {
	if s == "" {
		return "empty string"
	} else {
		return s
	}
}

func assembleFile(spc *spcgo.SPCfile, path *string, verbose bool) *file {
	var output file

	return &output
}

func saveFile(File *file, verbose bool) error {
	var err error

	return err
}

func main() {
	var filename string
	var verbose bool
	var reportdata bool
	flag.StringVar(&filename, "filename", "RAMAN.SPC", "filename to read defaults to RAMAN.SPC")
	flag.BoolVar(&verbose, "verbose", false, "boolen to print the details or not defaults to false")
	flag.BoolVar(&reportdata, "reportdata", false, "boolen to print the spectra to the console defaults to false")
	flag.Parse()
	if !verbose {
		fmt.Printf("Requested minimal output. Use -verbose=true to see more.\n")
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("File <%s> does not exist\n", checkEmpty(filename))
		fmt.Printf("Check that everything is right and use -filename=<file.spc>\n")
	} else {
		fmt.Printf("Reading <%s>\n", filename)
	}
	SPC := spcgo.ReadSPC(filename, verbose)

	var numpts int32 = 5
	if reportdata {
		fmt.Printf("\nData in file:\n")
		for i := int32(0); i < numpts; i++ {
			fmt.Printf("%d: %f, %f\n", i+1, (*SPC.Data.X)[i], (*SPC.Data.Y)[i])
		}
		fmt.Printf("...\n...\n...\n")
		for i := SPC.Head.Fnpts - numpts; i < SPC.Head.Fnpts; i++ {
			fmt.Printf("%d: %f, %f\n", i+1, (*SPC.Data.X)[i], (*SPC.Data.Y)[i])
		}
	}
}
