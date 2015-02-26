package main

import (
        "encoding/csv"
        "fmt"
        "os"
        "regexp"
)

type Reader struct {
        Comma            rune // field delimiter (set to ',' by NewReader)
        Comment          rune // comment character for start of line
        FieldsPerRecord  int  // number of expected fields per record
        LazyQuotes       bool // allow lazy quotes
        TrailingComma    bool // ignored; here for backwards compatibility
        TrimLeadingSpace bool // trim leading space
        // contains filtered or unexported fields
}

func main() {

        // csvfile, err := os.Open("cdrs-teltech.csv")
        csvfile, err := os.Open("cdrs-teltech.csv")
        output, err := os.Create("all_others.csv")

        if err != nil {
                fmt.Println(err)
                return
        }

        defer csvfile.Close()
        defer output.Close()

        reader := csv.NewReader(csvfile)
        writer := csv.NewWriter(output)

        reader.FieldsPerRecord = 15 // see the Reader struct information below

        rawCSVdata, err := reader.ReadAll()

        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }

        // sanity check, display to standard output
        // for _, each := range rawCSVdata {
        //      if ok, _ := regexp.MatchString("UNITED STATES.*", each[0]); ok || each[0] == "CANADA" {
        //              err := writer.Write(each)
        //              if err != nil {
        //                      fmt.Println("Error:", err)
        //                      return
        //              }
        //      }
        // }

        // ALL OTHERS // ALL OTHERS //
        r, _ := regexp.Compile("UNITED STATES.*")
        c, _ := regexp.Compile("CANADA.*")

        for _, each := range rawCSVdata {
                if us := r.MatchString(each[0]); us == false {
                        if canada := c.MatchString(each[0]); canada == false {
                                fmt.Println(each)
                                err := writer.Write(each)
                                if err != nil {
                                        fmt.Println("Error:", err)
                                        return
                                }
                        }
                }
        }
        writer.Flush()
}
