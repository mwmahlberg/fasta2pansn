/*
Copyright © 2024 Markus W Mahlberg

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mwmahlberg/fasta2pansn/fasta"
	"github.com/mwmahlberg/fasta2pansn/pansn"
)

var (
	delimiter   string = "#"
	haplotypeID string = "1"
	sampleName  string = "Sample1"
	fastaFile   string
)

func init() {
	flag.StringVar(&delimiter, "delimiter", delimiter, "Delimiter to use between fields in the pansn output")
	flag.StringVar(&haplotypeID, "haplotype-id", haplotypeID, "Haplotype ID to use in the pansn output")
	flag.StringVar(&sampleName, "sample", sampleName, "Sample name to use in the pansn output")
	flag.StringVar(&fastaFile, "fasta", "", "Path to the input FASTA file (required)")
}

func main() {

	flag.Parse()
	if fastaFile == "" {
		flag.PrintDefaults()
		return
	}

	// Check if the file exists, is not a directory
	// and we have permission to read it.
	// If not, print an error and exit with a non-zero status code.
	if fi, err := os.Stat(fastaFile); fi.IsDir() {
		fmt.Printf("file %s is a directory\n", fastaFile)
		flag.PrintDefaults()
		os.Exit(1)
	} else if err == os.ErrNotExist {
		fmt.Printf("file %s does not exist\n", fastaFile)
		os.Exit(2)
	} else if err == os.ErrPermission {
		fmt.Printf("no permission to access file %s: %s\n", fastaFile, fi.Mode())
		os.Exit(3)
	} else if err != nil {
		fmt.Printf("error accessing file %s: %s\n", fastaFile, err)
		os.Exit(4)
	}

	f, err := os.Open(fastaFile)
	if err != nil {
		fmt.Printf("error opening file %s: %s\n", fastaFile, err)
		os.Exit(5)
	}
	defer f.Close()

	dec := fasta.NewDecoder(f)
	seqs := make([]fasta.Record, 0)
	if err := dec.Decode(&seqs); err != nil {
		fmt.Printf("error decoding file %s: %s\n", fastaFile, err)
		os.Exit(6)
	}

	enc, err := pansn.NewEncoder(os.Stdout, pansn.WithDelimiter(delimiter), pansn.WithHaplotypeID(haplotypeID), pansn.WithSampleName(sampleName))
	if err != nil {
		fmt.Printf("error creating encoder: %s\n", err)
		os.Exit(7)
	}

	if err = enc.Encode(seqs); err != nil {
		fmt.Printf("error encoding sequences: %s\n", err)
		os.Exit(8)
	}
}
