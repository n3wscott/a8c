/*
Copyright 2022 Scott Nichols.
SPDX-License-Identifier: Apache-2.0
*/

package addition

import (
	"fmt"
	"io"
	"math/rand"
	"strings"
)

type Interface interface {
	// Generate will write `count` equations to `out`, each digit less than `max`.
	Generate(out io.Writer, min, max, count int) error
}

func New() Interface {
	return &impl{}
}


type impl struct {}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Generate outputs a random addition equation like:
//    aaaa
//  + bbbb
// --------
//
//
//
func (i impl) Generate(out io.Writer, min, max, count int) error {
	padding := len(fmt.Sprintf("%d", max))
	line := strings.Repeat("-", padding+5)
	for i := 0; i < count; i++ {
		a := randomInt(min, max)
		b := randomInt(min, max)
		fmt.Fprintf(out, "%*d\n", padding+3, a)
		fmt.Fprintf(out, "%*s%*d\n", 3, "+ ", padding, b)
		fmt.Fprintf(out, "%s\n", line)
		fmt.Fprintf(out, "\n\n\n")
	}
	return nil
}
