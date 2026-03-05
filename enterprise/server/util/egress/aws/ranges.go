package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"log"
	"os"
)

type awsIPRanges struct {
	Prefixes []struct {
		IPPrefix string `json:"ip_prefix"`
		Region   string `json:"region"`
	} `json:"prefixes"`
	IPv6Prefixes []struct {
		IPv6Prefix string `json:"ipv6_prefix"`
		Region     string `json:"region"`
	} `json:"ipv6_prefixes"`
}

// This simple Go program converts AWS IP block ranges as specified in
// https://ip-ranges.amazonaws.com/ip-ranges.json into this CSV format:
// AWS,<REGION>,<IP-BLOCK>
func main() {
	inputPath := flag.String("input", "", "Path to AWS raw-ranges.json")
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

	var ranges awsIPRanges
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

	if err := w.Write([]string{"AWS", "REGION", "IP-BLOCK"}); err != nil {
		log.Fatalf("write header: %v", err)
	}
	for _, p := range ranges.Prefixes {
		if err := w.Write([]string{"AWS", p.Region, p.IPPrefix}); err != nil {
			log.Fatalf("write IPv4 row: %v", err)
		}
	}
	for _, p := range ranges.IPv6Prefixes {
		if err := w.Write([]string{"AWS", p.Region, p.IPv6Prefix}); err != nil {
			log.Fatalf("write IPv6 row: %v", err)
		}
	}

	if err := w.Error(); err != nil {
		log.Fatalf("flush CSV: %v", err)
	}
}
