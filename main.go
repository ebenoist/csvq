package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "csvq"
	app.Usage = "A Simple CSV Tool"
	app.UsageText = "cat my.csv | csvq"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "col, c",
			Usage: "Specify the columns to print by name or number",
		},
		cli.BoolFlag{
			Name:  "json, j",
			Usage: "convert to json for use in jq",
		},
	}

	app.Action = func(c *cli.Context) {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			csv := readCSV(os.Stdin)
			cols := parseColumns(c.GlobalString("col"), csv[0])

			if c.GlobalBool("json") {
				printJSON(csv, cols)
			} else {
				printCSV(csv, cols)
			}
		} else {
			cli.ShowAppHelp(c)
		}
	}

	app.Run(os.Args)
}

func printJSON(records [][]string, cols []int) {
	header := selectCols(records[0], cols)

	var out []map[string]interface{}

	for _, r := range records {
		s := selectCols(r, cols)

		jRow := make(map[string]interface{})
		for idx, val := range s {
			jRow[header[idx]] = val
		}

		out = append(out, jRow)
	}

	r, err := json.Marshal(out)
	if err != nil {
		panic(err)
	}

	fmt.Print(string(r))
}

func printCSV(records [][]string, cols []int) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 2, 1, ' ', tabwriter.DiscardEmptyColumns)
	for _, r := range records {
		s := selectCols(r, cols)
		fmt.Fprintln(writer, strings.Join(s, "\t"))
	}

	writer.Flush()
}

func readCSV(in io.Reader) [][]string {
	csv := csv.NewReader(in)
	var records [][]string

	for {
		record, err := csv.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		records = append(records, record)
	}

	return records
}

func parseColumns(col string, header []string) []int {
	var columns []int
	if col == "" {
		return columns
	}

	for _, c := range strings.Split(col, ",") {
		col, err := strconv.Atoi(c)

		if err != nil {
			col = indexOf(c, header)
		}

		if col > len(header)-1 || col < 0 {
			log.Fatalf("Column %q does not exist", col)
		}

		columns = append(columns, col)
	}

	return columns
}

func indexOf(s string, a []string) int {
	for i, b := range a {
		if strings.Compare(s, b) == 0 {
			return i
		}
	}

	return -1
}

func selectCols(record []string, columns []int) []string {
	if len(columns) == 0 {
		return record
	}

	var s []string

	for _, c := range columns {
		s = append(s, record[c])
	}

	return s
}
