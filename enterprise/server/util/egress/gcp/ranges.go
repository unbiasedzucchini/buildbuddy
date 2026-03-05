package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"log"
	"os"
)

type gcpIPRanges struct {
	Prefixes []struct {
		IPv4Prefix string `json:"ipv4Prefix"`
		IPv6Prefix string `json:"ipv6Prefix"`
		Scope      string `json:"scope"`
	} `json:"prefixes"`
}

// This simple Go program converts GCP IP block ranges as specified in
// https://www.gstatic.com/ipranges/cloud.json into this CSV format:
// GCP,<REGION>,<IP-BLOCK>
func main() {
	inputPath := flag.String("input", "", "Path to GCP ranges.json")
	outputPath := flag.String("output", "", "Path to output CSV")
	flag.Parse()

	if *inputPath == "" || *outputPath == "" {
		log.Fatal("--input and --output are required")
	}

	input, err := os.Open(*inputPath)
	if err != nil {
		log.Fatalf("open input: %v", err)
	}
	defer input.Close()

	var ranges gcpIPRanges
	if err := json.NewDecoder(input).Decode(&ranges); err != nil {
		log.Fatalf("decode JSON: %v", err)
	}

	output, err := os.Create(*outputPath)
	if err != nil {
		log.Fatalf("create output: %v", err)
	}
	defer output.Close()

	w := csv.NewWriter(output)
	defer w.Flush()

	if err := w.Write([]string{"GCP", "REGION", "IP-BLOCK"}); err != nil {
		log.Fatalf("write header: %v", err)
	}
	for _, p := range ranges.Prefixes {
		if p.IPv4Prefix != "" {
			if err := w.Write([]string{"GCP", p.Scope, p.IPv4Prefix}); err != nil {
				log.Fatalf("write IPv4 row: %v", err)
			}
		}
		if p.IPv6Prefix != "" {
			if err := w.Write([]string{"GCP", p.Scope, p.IPv6Prefix}); err != nil {
				log.Fatalf("write IPv6 row: %v", err)
			}
		}
	}

	if err := w.Error(); err != nil {
		log.Fatalf("flush CSV: %v", err)
	}
}
