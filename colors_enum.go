// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package main

import (
	"fmt"
	"strings"
)

const (
	// ColorRed is a Color of type Red.
	ColorRed Color = iota
	// ColorGreen is a Color of type Green.
	ColorGreen
	// ColorBlue is a Color of type Blue.
	ColorBlue
)

var ErrInvalidColor = fmt.Errorf("not a valid Color, try [%s]", strings.Join(_ColorNames, ", "))

const _ColorName = "RedGreenBlue"

var _ColorNames = []string{
	_ColorName[0:3],
	_ColorName[3:8],
	_ColorName[8:12],
}

// ColorNames returns a list of possible string values of Color.
func ColorNames() []string {
	tmp := make([]string, len(_ColorNames))
	copy(tmp, _ColorNames)
	return tmp
}

var _ColorMap = map[Color]string{
	ColorRed:   _ColorName[0:3],
	ColorGreen: _ColorName[3:8],
	ColorBlue:  _ColorName[8:12],
}

// String implements the Stringer interface.
func (x Color) String() string {
	if str, ok := _ColorMap[x]; ok {
		return str
	}
	return fmt.Sprintf("Color(%d)", x)
}

var _ColorValue = map[string]Color{
	_ColorName[0:3]:  ColorRed,
	_ColorName[3:8]:  ColorGreen,
	_ColorName[8:12]: ColorBlue,
}

// ParseColor attempts to convert a string to a Color.
func ParseColor(name string) (Color, error) {
	if x, ok := _ColorValue[name]; ok {
		return x, nil
	}
	return Color(0), fmt.Errorf("%s is %w", name, ErrInvalidColor)
}
