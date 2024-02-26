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

package fasta

import (
	"bufio"
	"fmt"
	"io"
)

// Record represents a single sequence record in a FASTA file
type Record struct {
	Header string
	Seq    string
}

// Decoder is a FASTA file decoder
type Decoder struct {
	source io.Reader
}

// NewDecoder returns a new FastaDecoder.
func NewDecoder(in io.Reader) *Decoder {
	return &Decoder{source: in}
}

// Decode reads from the source and writes the output to out.
// out must be a pointer to a slice of fasta.Record.
// Returns an error if the source cannot be read.
func (d *Decoder) Decode(out *[]Record) error {

	// We use a scanner to read the source line by line
	scanner := bufio.NewScanner(d.source)

	// We reuse the same Record struct for each sequence
	var seq Record
	for scanner.Scan() {

		// If there is an error reading from the source, return it
		if scanner.Err() != nil {
			return fmt.Errorf("error reading from source: %w", scanner.Err())
		}

		line := scanner.Text()
		// If the line starts with '>', we have a new sequence
		if line[0] == '>' {
			// If the old sequence is not empty, append it to the output now...
			if seq.Header != "" {
				*out = append(*out, seq)
			}
			// ...and create a new sequence record
			seq = Record{Header: line[1:]}
		} else {
			// If the line does not start with '>', it is part of the sequence
			seq.Seq += line
		}
	}
	// Append the last sequence to the output
	*out = append(*out, seq)
	return nil
}
