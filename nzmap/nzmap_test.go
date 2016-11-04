package nzmap

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

// these tests also output SVG to svg_test/ for visual inspection.

func TestIconWellington(t *testing.T) {
	var b bytes.Buffer
	p := Point{Longitude: 174.7,
		Latitude: -41.2,
		Value:    -1,
	}
	p.Icon(&b)

	if !p.Visible() {
		t.Error("point should be on the icon map")
	}

	b.WriteString(fmt.Sprintf("<circle cx=\"%d\" cy=\"%d\" r=\"5\" fill=\"black\"/></svg>", p.X(), p.Y()))

	if err := ioutil.WriteFile("svg_test/nzicon-wellington.svg", b.Bytes(), 0644); err != nil {
		t.Fatal(err)
	}
}

func TestIconRaoul(t *testing.T) {
	var b bytes.Buffer
	p := Point{Longitude: -177.9286,
		Latitude: -29.2684,
		Value:    -1,
	}
	p.Icon(&b)

	if !p.Visible() {
		t.Error("point should be on the icon map")
	}

	b.WriteString(fmt.Sprintf("<circle cx=\"%d\" cy=\"%d\" r=\"5\" fill=\"black\"/></svg>", p.X(), p.Y()))

	if err := ioutil.WriteFile("svg_test/nzicon-raoul.svg", b.Bytes(), 0644); err != nil {
		t.Fatal(err)
	}
}

func TestIconAucklandIsland(t *testing.T) {
	var b bytes.Buffer
	p := Point{Longitude: 166.102,
		Latitude: -50.72,
		Value:    -1,
	}
	p.Icon(&b)

	if !p.Visible() {
		t.Error("point should be on the icon map")
	}

	b.WriteString(fmt.Sprintf("<circle cx=\"%d\" cy=\"%d\" r=\"5\" fill=\"black\"/></svg>", p.X(), p.Y()))

	if err := ioutil.WriteFile("svg_test/nzicon-aucklandisland.svg", b.Bytes(), 0644); err != nil {
		t.Fatal(err)
	}
}

func TestIconCanberra(t *testing.T) {
	var b bytes.Buffer
	p := Point{Longitude: 149.1300,
		Latitude: -35.2809,
		Value:    -1,
	}
	p.Icon(&b)

	if p.Visible() {
		t.Error("point should not be on the icon map")
	}

	b.WriteString(fmt.Sprintf("<circle cx=\"%d\" cy=\"%d\" r=\"5\" fill=\"black\"/></svg>", p.X(), p.Y()))

	if err := ioutil.WriteFile("svg_test/nzicon-canberra.svg", b.Bytes(), 0644); err != nil {
		t.Fatal(err)
	}
}

func TestMediumWellington(t *testing.T) {
	var b bytes.Buffer

	var pt Points

	pt = append(pt, Point{Longitude: 174.7,
		Latitude: -41.2,
		Value:    -1,
	})

	pt.Medium(&b)

	for _, p := range pt {
		if !p.Visible() {
			t.Error("point should be on the map")
		}

		b.WriteString(fmt.Sprintf("<circle cx=\"%d\" cy=\"%d\" r=\"5\" fill=\"black\"/>", p.X(), p.Y()))
	}

	b.WriteString("</svg>")

	if err := ioutil.WriteFile("svg_test/nzmedium-wellington.svg", b.Bytes(), 0644); err != nil {
		t.Fatal(err)
	}
}

func TestMediumRaoul(t *testing.T) {
	var b bytes.Buffer

	var pt Points

	pt = append(pt, Point{Longitude: -177.9286,
		Latitude: -29.2684,
		Value:    -1,
	})

	pt.Medium(&b)

	for _, p := range pt {
		if p.Visible() {
			t.Error("point should not be on the map")
		}

		b.WriteString(fmt.Sprintf("<circle cx=\"%d\" cy=\"%d\" r=\"5\" fill=\"black\"/>", p.X(), p.Y()))
	}

	b.WriteString("</svg>")

	if err := ioutil.WriteFile("svg_test/nzmedium-raoul.svg", b.Bytes(), 0644); err != nil {
		t.Fatal(err)
	}
}
