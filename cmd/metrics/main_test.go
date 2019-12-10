package main

import (
	"log"
	"testing"
)

type DataWriter interface {
	Write()
}

type Data struct {
	Value string
}

func (d *Data) Write() {
	log.Printf("wrote %s", d.Value)
}

func gen() ([]DataWriter, error) {
	d := make([]DataWriter, 0)
	d = append(d, DataWriter(&Data{Value: "some"}))
	return d, nil
}

func TestTest(t *testing.T) {
	d, e := gen()
	if e != nil {
		t.Error("something bad happened")
	} else {
		for _, w := range d {
			w.Write()
		}
	}
}
