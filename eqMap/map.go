package eqMap

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"strconv"
)

type EqMap struct {
	Lines  []Line
	Points []Point
	Ms     []M
}

type Color struct {
	R, G, B uint8
}

type (
	Point struct {
		X, Y, Z float64
		Color   Color
		Size    uint8
		Text    string
	}
	Line struct {
		X1, X2, Y1, Y2, Z1, Z2 float64
		Color                  Color
	}
	M struct{}
)

const (
	pointLen = 8
	lineLen  = 9
)

func (m *EqMap) Load(r io.Reader) error {
	s := bufio.NewScanner(r)

	var b []byte
	for s.Scan() {
		b = s.Bytes()

		switch b[0] {
		case 'l':
			l, err := ParseLine(b[1:])
			if err != nil {
				return err
			}

			m.Lines = append(m.Lines, l)
		case 'p':
			p, err := ParsePoint(b[1:])
			if err != nil {
				return err
			}

			m.Points = append(m.Points, p)
		default:
			slog.Error("unable to parse: ", "line", string(b))
		}
	}

	return nil
}

func ParseLine(b []byte) (Line, error) {
	d := bytes.Split(b, []byte{','})

	l := Line{
		Color: Color{},
	}

	if len(d) != lineLen {
		return l, fmt.Errorf("invalid length: want %d, got %d", lineLen, len(d))
	}

	var err error
	if l.X1, err = strconv.ParseFloat(string(d[0]), 64); err != nil {
		return l, err
	}

	if l.X2, err = strconv.ParseFloat(string(d[1]), 64); err != nil {
		return l, err
	}

	if l.Y1, err = strconv.ParseFloat(string(d[2]), 64); err != nil {
		return l, err
	}

	if l.Y2, err = strconv.ParseFloat(string(d[3]), 64); err != nil {
		return l, err
	}

	if l.Z1, err = strconv.ParseFloat(string(d[4]), 64); err != nil {
		return l, err
	}

	if l.Z2, err = strconv.ParseFloat(string(d[5]), 64); err != nil {
		return l, err
	}

	if l.Color.R, err = parseUint8(string(d[6])); err != nil {
		return l, err
	}

	if l.Color.G, err = parseUint8(string(d[7])); err != nil {
		return l, err
	}

	if l.Color.B, err = parseUint8(string(d[8])); err != nil {
		return l, err
	}

	return l, nil
}

func ParsePoint(b []byte) (Point, error) {
	d := bytes.Split(b, []byte{','})

	p := Point{
		Color: Color{},
	}

	if len(d) != pointLen {
		return p, fmt.Errorf("invalid length: want %d, got %d", pointLen, len(d))
	}

	var err error
	if p.X, err = strconv.ParseFloat(string(d[0]), 64); err != nil {
		return p, err
	}

	if p.X, err = strconv.ParseFloat(string(d[1]), 64); err != nil {
		return p, err
	}

	if p.Y, err = strconv.ParseFloat(string(d[2]), 64); err != nil {
		return p, err
	}

	if p.Z, err = strconv.ParseFloat(string(d[3]), 64); err != nil {
		return p, err
	}

	if p.Size, err = parseUint8(string(d[4])); err != nil {
		return p, err
	}

	p.Text = string(d[5])
	if p.Color.R, err = parseUint8(string(d[6])); err != nil {
		return p, err
	}

	if p.Color.G, err = parseUint8(string(d[7])); err != nil {
		return p, err
	}

	if p.Color.B, err = parseUint8(string(d[8])); err != nil {
		return p, err
	}

	return p, nil
}

func parseUint8(s string) (uint8, error) {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint8(i), nil
}
