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

package pansn_test

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"

	"github.com/mwmahlberg/fasta2pansn/fasta"
	"github.com/mwmahlberg/fasta2pansn/pansn"
	"github.com/stretchr/testify/assert"
)

var records = []fasta.Record{{Header: "Chr01", Seq: "ACGTACGTACGTACGT"}, {Header: "Chr02", Seq: "TACGTACGTACGTACG"}}

func TestEncoderSingle(t *testing.T) {
	randomId := rand.Int63n(1000)

	buf := bytes.NewBuffer([]byte{})
	enc, err := pansn.NewEncoder(buf, pansn.WithSampleName("Sample1"), pansn.WithHaplotypeID(fmt.Sprintf("Hap%d", randomId)))
	assert.NoError(t, err)
	assert.NotNil(t, enc)

	err = enc.Encode(records[0])
	assert.NoError(t, err)
	assert.Equal(t, fmt.Sprintf(">Sample1#Hap%d#Chr01\nACGTACGTACGTACGT\n", randomId), buf.String())
}

func TestEncoderMultiple(t *testing.T) {
	randomId := rand.Int63n(1000)

	buf := bytes.NewBuffer([]byte{})
	enc, err := pansn.NewEncoder(buf, pansn.WithSampleName("Sample1"), pansn.WithHaplotypeID(fmt.Sprintf("Hap%d", randomId)))
	assert.NoError(t, err)
	assert.NotNil(t, enc)

	err = enc.Encode(records)
	assert.NoError(t, err)
	assert.Contains(t, buf.String(), fmt.Sprintf(">Sample1#Hap%d#Chr01\nACGTACGTACGTACGT\n", randomId))
	assert.Contains(t, buf.String(), fmt.Sprintf(">Sample1#Hap%d#Chr02\nTACGTACGTACGTACG\n", randomId))
}
