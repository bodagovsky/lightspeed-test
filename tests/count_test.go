package tests

import (
	"fmt"
	"testing"

	"github.com/bodagovsky/lightspeed-test/counter"
	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	cases := []struct {
		title    string
		input    [4]uint8
		expected uint32
	}{
		{
			title:    "case01",
			input:    [4]uint8{125, 168, 0, 1},
			expected: 0b1111101101010000000000000000001,
		},
		{
			title:    "case02",
			input:    [4]uint8{0, 0, 0, 0},
			expected: 0b0,
		},
		{
			title:    "case03",
			input:    [4]uint8{255, 255, 255, 255},
			expected: 0b11111111111111111111111111111111,
		},
		{
			title:    "case04",
			input:    [4]uint8{0, 168, 0, 1},
			expected: 0b00000000101010000000000000000001,
		},
		{
			title:    "case05",
			input:    [4]uint8{1, 1, 1, 1},
			expected: 0b00000001000000010000000100000001,
		},
	}

	for _, tt := range cases {
		t.Run(tt.title, func(t *testing.T) {
			assert.Equal(t, tt.expected, counter.Encode(tt.input))
		})
	}
}

func TestCountSequential(t *testing.T) {
	cases := []struct {
		title    string
		input    []string
		expected uint64
	}{
		{
			title: "case 01",
			input: []string{
				"125.168.0.1",
				"125.168.0.1",
				"125.168.0.1",
				"125.168.0.1",
				"125.168.0.1",
			},
			expected: 1,
		},
		{
			title: "case 02",
			input: []string{
				"125.168.0.1",
				"125.168.0.2",
				"125.168.0.3",
				"125.168.0.4",
				"125.168.0.5",
			},
			expected: 5,
		},
		{
			title: "case 03",
			input: []string{
				"125.168.0.1",
				"10.168.0.2",
				"125.11.0.3",
				"125.168.0.255",
				"125.11.0.3",
				"255.255.255.0",
				"255.255.255.1",
				"125.168.0.1",
			},
			expected: 6,
		},
	}

	for _, tt := range cases {
		t.Run(tt.title, func(t *testing.T) {
			ctr := counter.NewIpAdressCounter()
			var actual uint64
			for _, line := range tt.input {
				actual = ctr.Process(line)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkCount(b *testing.B) {
	ctr := counter.NewIpAdressCounter()
	for a := 0; a < 64; a++ {
		for b := 0; b < 64; b++ {
			for c := 0; c < 64; c++ {
				for d := 0; d < 64; d++ {
					ctr.Process(fmt.Sprintf("%d.%d.%d.%d", a, b, c, d))
				}
			}
		}
	}
}
