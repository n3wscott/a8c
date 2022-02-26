/*
Copyright 2022 Scott Nichols.
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/n3wscott/a8c/pkg/addition"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a := addition.New()

	err := a.Generate(os.Stdout, 1, 9999, 1000)
	if err != nil {
		panic(err)
	}
}
