/*
The code in this project is a rewrite version of the changefinder package written
in Python with golang. The URL of the original code is
https://pypi.org/project/changefinder/, and it is published under MIT Licence below.

The MIT License (MIT)

Copyright (c) 2013 Shunsuke Aihara

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package cf

// CF is a main struct of changefinder.
type CF struct {
	smooth         int
	order          int
	r              float64
	ts             []float64
	firstScores    []float64
	smoothedScores []float64
	secondScores   []float64
	convolve       []float64
	smooth2        int
	convolve2      []float64
	sdarFirst      *sdar1Dim
	sdarSecond     *sdar1Dim
}

// ChangeFinder returns a new CF struct.
func ChangeFinder(r float64, order, smooth int) *CF {
	cf := CF{
		smooth:         smooth,
		order:          order,
		r:              r,
		ts:             make([]float64, 0, 64),
		firstScores:    make([]float64, 0, smooth),
		smoothedScores: make([]float64, 0, smooth),
		secondScores:   make([]float64, 0, smooth),
		smooth2:        int(smooth / 2.0),
		sdarFirst:      newSDAR1Dim(r, order),
		sdarSecond:     newSDAR1Dim(r, order),
	}

	// ones
	cf.convolve = make([]float64, cf.smooth)
	for i := range cf.convolve {
		cf.convolve[i] = 1.0
	}
	cf.convolve2 = make([]float64, cf.smooth2)
	for i := range cf.convolve2 {
		cf.convolve2[i] = 1.0
	}

	return &cf
}

func (cf *CF) addOne(one float64, ts *[]float64, size int) {
	*ts = append(*ts, one)
	if len(*ts) == size+1 {
		*ts = (*ts)[1:]
	}
}

func (cf *CF) smoothing(ts []float64) float64 {
	var ave float64
	for i := 0; i < cf.smooth; i++ {
		ave += ts[i] * cf.convolve[i]
	}
	return ave
}

func (cf *CF) smoothing2(ts []float64) float64 {
	var ave float64
	for i := 0; i < cf.smooth2; i++ {
		ave += ts[i] * cf.convolve2[i]
	}
	return ave
}

// Update returns a new score of CF.
func (cf *CF) Update(x float64) float64 {
	score := float64(0.0)

	// First step learning
	if len(cf.ts) == cf.order {
		score = cf.sdarFirst.update(x, cf.ts)
		cf.addOne(score, &cf.firstScores, cf.smooth)
	}
	cf.addOne(x, &cf.ts, cf.order)

	var secondTarget float64

	// Smoothing
	if len(cf.firstScores) == cf.smooth {
		secondTarget = cf.smoothing(cf.firstScores)
		// log.Printf("firstScores: %v\n", cf.firstScores)
		// log.Printf("secondTarget: %v\n", secondTarget)
	}
	// Second step learning
	if secondTarget != 0.0 && len(cf.smoothedScores) == cf.order {
		score = cf.sdarSecond.update(secondTarget, cf.smoothedScores)
		cf.addOne(score, &cf.secondScores, cf.smooth2)
	}

	if secondTarget != 0.0 {
		cf.addOne(secondTarget, &cf.smoothedScores, cf.order)
	}
	if len(cf.secondScores) == cf.smooth2 {
		score = cf.smoothing2(cf.secondScores)
		// log.Printf("secondScores: %v\n", cf.secondScores)
		// log.Printf("secondSmoothed: %v\n", score)
		return score
	}

	return 0.0
}
