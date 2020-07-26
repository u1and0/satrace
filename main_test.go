package main

import (
	"math"
	"testing"
)

func Test_db2mw(t *testing.T) {
	var actual []float64
	for _, f := range []float64{0, 3, 6, 10} {
		actual = append(actual, math.Ceil(db2mw(f)))
	}
	// 10**0.1=1, 10**0.3~=1.99, 10**0.6~=3.98, 10**1=10
	expected := []float64{1, 2, 4, 10}
	for i, e := range expected {
		if actual[i] != e {
			t.Fatalf("got: %v want: %v\ndump all: %v", actual[i], e, actual)
		}
	}
}

// func Test_signalBand(t *testing.T) {
// 	f := conternArray{}
// }
func Test_readTrace(t *testing.T) {
	filename := "data/20200627_180505.txt"
	usecol := 1
	actualConf, actualCont, err := readTrace(filename, usecol)
	if err != nil {
		panic(err)
	}

	// Config test
	expectedConfig := map[string]string{
		":INP:COUP":              "DC",
		":BAND:RES":              "1 Hz",
		":AVER:COUNT":            "10",
		":SWE:POIN":              "11",
		":FREQ:CENT":             "5 MHz",
		":FREQ:SPAN":             "1 MHz",
		":TRAC1:TYPE":            "AVER",
		":INIT:CONT":             "0",
		":FORM":                  "REAL,32",
		":FORM:BORD":             "SWAP",
		":INIT:IMM":              "",
		":POW:ATT":               "0",
		":DISP:WIND:TRAC:Y:RLEV": "-30 dBm",
	}
	for k, v := range actualConf {
		if expectedConfig[k] != v {
			t.Fatalf("got: %v want: %v", v, expectedConfig[k])
		}
	}

	// Content test
	expectedCont := []float64{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	for i, e := range expectedCont {
		if actualCont[i] != e {
			t.Fatalf("got: %v want: %v\ndump all: %v", actualCont[i], e, actualCont)
		}
	}
}

/*
func bench(b *testing.B, a []string) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		writeBuffer(a)
	}
}

func Benchmark(b *testing.B) {
	files := []string{
		"data/20200627_180505.txt",
	}
	bench(b, files)
}
*/