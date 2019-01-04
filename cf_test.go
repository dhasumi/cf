package cf_test

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/dhasumi/cf"
)

func TestCF(t *testing.T) {
	f, err := os.Open("./testdata/data.out")
	defer f.Close()
	if err != nil {
		t.Errorf("could not open a test file")
	}

	sc := bufio.NewScanner(f)
	data := make([]float64, 0, 1200)
	for sc.Scan() {
		if err := sc.Err(); err != nil {
			break
		}
		f64, err := strconv.ParseFloat(sc.Text(), 64)
		if err != nil {
			t.Errorf("float parsing error")
		}
		data = append(data, f64)
	}

	cpd := cf.ChangeFinder(0.01, 1, 7)
	ret := make([]float64, 0, 1200)
	for _, v := range data {
		score := cpd.Update(v)
		ret = append(ret, score)
	}

	// write ret
	fw, err := os.OpenFile("./testdata/ret.out", os.O_WRONLY|os.O_CREATE, 0644)
	defer fw.Close()

	for _, v := range ret {
		fw.WriteString(fmt.Sprintf("%f\n", v))
	}
}

func BenchmarkCF(b *testing.B) {
	f, err := os.Open("./testdata/data.out")
	defer f.Close()
	if err != nil {
		b.Errorf("could not open a test file")
	}

	sc := bufio.NewScanner(f)
	data := make([]float64, 0, 1200)
	for sc.Scan() {
		if err := sc.Err(); err != nil {
			break
		}
		f64, err := strconv.ParseFloat(sc.Text(), 64)
		if err != nil {
			b.Errorf("float parsing error")
		}
		data = append(data, f64)
	}

	b.ReportAllocs()
	b.ResetTimer()

	cpd := cf.ChangeFinder(0.01, 1, 7)
	ret := make([]float64, 0, 120)
	for _, v := range data {
		score := cpd.Update(v)
		ret = append(ret, score)
	}

	b.StopTimer()
}
