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

package pansn

import (
	"bufio"
	"fmt"
	"io"

	"github.com/mwmahlberg/fasta2pansn/fasta"
)

type Encoder struct {
	SampleName  string
	HaplotypeID string
	delimiter   string
	target      io.Writer
}

// EncoderOption is a functional option for the PansnEncoder
type EncoderOption func(*Encoder) error

// WithDelimiter sets the delimiter for the PansnEncoder
func WithDelimiter(d string) EncoderOption {
	return func(e *Encoder) error {
		e.delimiter = d
		return nil
	}
}

// WithSampleName sets the sample name for the PansnEncoder
func WithSampleName(s string) EncoderOption {
	return func(e *Encoder) error {
		e.SampleName = s
		return nil
	}
}

// WithHaplotypeID sets the haplotype ID for the PansnEncoder
func WithHaplotypeID(h string) EncoderOption {
	return func(e *Encoder) error {
		e.HaplotypeID = h
		return nil
	}
}

// NewEncoder returns a new PansnEncoder.
// Unless WithDelimiter is used, the default delimiter is "#".
func NewEncoder(target io.Writer, opts ...EncoderOption) (*Encoder, error) {
	e := &Encoder{
		target:    target,
		delimiter: "#",
	}
	for _, opt := range opts {
		if err := opt(e); err != nil {
			return nil, err
		}
	}
	return e, nil
}

func (e *Encoder) write(sampleName, haplotypeID, contigOrScaffoldName, payload string) error {
	buf := bufio.NewWriter(e.target)
	defer buf.Flush()
	_, err := buf.WriteString(fmt.Sprintf(
		">%s%s%s%s%s\n%s\n",
		sampleName, e.delimiter,
		haplotypeID, e.delimiter,
		contigOrScaffoldName,
		payload))
	return err
}

// Encode writes the input to the target io.Writer.
// Currently, only fasta.Record and []fasta.Record are supported.
func (e *Encoder) Encode(in any) error {

	switch in.(type) {
	case []fasta.Record:
		for _, s := range in.([]fasta.Record) {
			if err := e.write(e.SampleName, e.HaplotypeID, s.Header, s.Seq); err != nil {
				return err
			}
		}
	case fasta.Record:
		s := in.(fasta.Record)
		if err := e.write(e.SampleName, e.HaplotypeID, s.Header, s.Seq); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported type: %T", in)
	}
	return nil
}
