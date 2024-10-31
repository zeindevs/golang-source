package main

import (
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseIMEI(t *testing.T) {
	input := `Result: Parcel(
0x00000000: 00000000 0000000f 00350033 00350037 '........3.2.9.2.'
0x00000010: 00320030 00300031 00370037 00350039 '0.2.1.0.7.0.2.5.'
0x00000020: 00370034 00000030                   '5.2.0...        ')`
	re := regexp.MustCompile(`'([^']*)'`)
	matches := re.FindAllStringSubmatch(input, -1)
	var combined string
	for _, match := range matches {
		if len(match) > 1 {
			cleaned := strings.ReplaceAll(match[1], ".", "")
			combined += cleaned
		}
	}
	imei := strings.TrimSpace(combined)
	assert.Equal(t, "329202107025520", imei)
}
