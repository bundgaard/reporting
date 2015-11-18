package main

import (
	"flag"
	"fmt"
	"github.com/bundgaard/myconfig"
	"os"
	"strings"
)

// http://stackoverflow.com/a/29347148

var (
	actual_start_date = flag.String("startdate", "", "2006-01-02")
	actual_end_date   = flag.String("enddate", "", "2006-01-02")
)

func splitheader(s string) []string {

	return strings.Split(s, ",")
}

func init() {

	flag.Usage = func() {

		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "%s version %s\n", os.Args[0], Version)
		fmt.Fprintf(os.Stderr, "built %s\n", BuildTime)
		fmt.Fprintln(os.Stderr, "")
		flag.PrintDefaults()

	}
}

func main() {
	flag.Parse()
	/*
		if len(os.Args) < 2 {
			flag.Usage()
			os.Exit(-1)
		}
	*/

	// Test if we are initiated from Command line, that means a colleague will need a specific report.
	if len(*actual_start_date) > 0 && len(*actual_end_date) > 0 {
		fmt.Printf("Initiated from CLI.\n")
		fmt.Printf("Chosen values: %s - %s\n", *actual_start_date, *actual_end_date)
		fmt.Printf("Should do some heavy lifting\n")
		fmt.Printf("Will exit afterwards, as to not initiate other stupidities.\n")
		os.Exit(0)
	}

	/* if we are here, that means we should look for monthly, weekly and other requirements
	 */

	config, err := myconfig.LoadConfiguration()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}

	fmt.Printf("%v\n", config)

}
