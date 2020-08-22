package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"encoding/json"
	"log"
	"os"
	"strconv"
)

type Point struct {
        Name  string
	Value float64 
}

type Line struct {
	Name                string
	Series              []Point
}

type Graph struct {
      Variables []Line 
}

var data Graph 
var lineUSD Line
var pointUSD Point
var seriesUSD []Point

var lineARG Line
var pointARG Point
var seriesARG []Point

var pointTmp Point
var GraphTmp []Line
var lineTmp Line

func getFileHandler(filename string) *csv.Reader {
    csvfileInput, err := os.Open(filename)
    if (err != nil) {
        log.Fatalln("could not open csv file", err)
    }
    r := csv.NewReader(csvfileInput)
    return r
}

func main() {

    header := map[string]Line {
	"USD":{},
	"ARG":{},
    }
    body := map[string][]Point {
	"USD": {},
	"ARG": {},
    }
    variables := map[string]int {
        "USD": 6,
        "ARG": 2,
    }

    r := getFileHandler("uva_uvi_usd_final.csv")

    for k, _ := range variables {
	header[k] = Line { 
            Name: k,
        }
        fmt.Printf("[HEADERS]  %s\n", k)
    }

    fmt.Printf("[SETTING VARIABLES] ... \n", )
    for {
        record, err := r.Read()
        if err == io.EOF {
            break
        }
        for k, v := range variables {
	    csvValue, err := strconv.ParseFloat(record[v], 64)
	    if err != nil {
                fmt.Printf("[SETTING VARIABLES ERROR]  %s\n", err)
                fmt.Printf("[SETTING VARIABLES ERROR]  %v\n", record[v])
	    }
	    PointTmp := Point { Name: record[0], Value: csvValue}
            body[k] = append(body[k], PointTmp)
        }
	/* fmt.Printf("Current Row name: %s \n", record[0])
	fmt.Printf("Current Row value usd %f \n", record[6])
	fmt.Printf("Current Row value uva %f \n", record[2]) */
    }
    fmt.Printf("[STAGING] Header %s \n", header)
    //fmt.Printf("[STAGING] Body %s \n", body)
    fmt.Printf("[FINAL] Setting GraphTmp \n")

    for k, v := range variables {
        fmt.Printf("[FINAL] key variables %s \n", k)
        fmt.Printf("[FINAL] v variables %s \n", v)
        lineTmp := header[k]
	lineTmp.Series = body[k]
        GraphTmp = append(GraphTmp, lineTmp)
    }
    fmt.Printf("[FINAL] Writing GraphTmp to a file \n")
    //c, err := json.Marshal(seriesUSD) write json in one line, and not formated
    file, _ := json.MarshalIndent(GraphTmp, "", " ")
    _ = ioutil.WriteFile("graphLineTemplate.json", file, 0644)
    fmt.Printf("[FINAL] END ce ya \n")
}
