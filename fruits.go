//go:generate stringer -type=Fruit
package main

type Fruit int

const (
	Apple Fruit = iota
	Banana
	Orange
)
