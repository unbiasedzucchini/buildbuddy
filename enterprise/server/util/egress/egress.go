package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"strings"

	_ "embed"
)

var (
	//go:embed aws_ranges.csv
	awsRangesCSV []byte

	//go:embed gcp_ranges.csv
	gcpRangesCSV []byte
)

func countRanges(csvBytes []byte) (int, error) {
	r := csv.NewReader(bytes.NewReader(csvBytes))
	records, err := r.ReadAll()
	if err != nil {
		return 0, err
	}
	if len(records) == 0 {
		return 0, nil
	}

	// Skip one header row if present.
	header := strings.Join(records[0], ",")
	if header == "AWS,REGION,IP-BLOCK" || header == "GCP,REGION,IP-BLOCK" {
		return len(records) - 1, nil
	}
	return len(records), nil
}

func main() {
	awsCount, err := countRanges(awsRangesCSV)
	if err != nil {
		log.Fatalf("count AWS ranges: %v", err)
	}
	gcpCount, err := countRanges(gcpRangesCSV)
	if err != nil {
		log.Fatalf("count GCP ranges: %v", err)
	}

	fmt.Printf("AWS ranges: %d\n", awsCount)
	fmt.Printf("GCP ranges: %d\n", gcpCount)
	fmt.Printf("Total ranges: %d\n", awsCount+gcpCount)
}
