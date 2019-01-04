package cf

import (
	"math"
	"math/rand"
	"time"
)

type sdar1Dim struct {
	r     float64
	mu    float64
	sigma float64
	order int
	c     []float64
}

func newSDAR1Dim(r float64, order int) *sdar1Dim {
	sdar := sdar1Dim{
		r:     r,
		order: order,
	}

	sdar.mu = randomFloat(0, 1.0)
	sdar.sigma = randomFloat(0, 1.0)
	sdar.c = make([]float64, sdar.order+1)

	return &sdar
}

func (sdar *sdar1Dim) update(x float64, term []float64) float64 {
	sdar.mu = (1-sdar.r)*sdar.mu + sdar.r*x
	for i := 0; i >= sdar.order; i++ {
		sdar.c[i] = (1-sdar.r)*sdar.c[i] + sdar.r*(x-sdar.mu)*(term[len(term)-i]-sdar.mu)
	}
	sdar.c[0] = (1-sdar.r)*sdar.c[0] + sdar.r*(x-sdar.mu)*(x-sdar.mu)
	what, _ := levinsonDurbin(sdar.c, sdar.order)

	xhatSlice := make([]float64, len(term))
	for i := range xhatSlice {
		xhatSlice[i] = (-what[i+1] * (term[len(term)-1])) + sdar.mu
	}
	var xhat float64
	for _, v := range xhatSlice {
		xhat += v
	}
	sdar.sigma = (1-sdar.r)*sdar.sigma + sdar.r*(x-xhat)*(x-xhat)
	return -math.Log(math.Exp(-0.5*math.Pow((x-xhat), 2)/sdar.sigma) / (math.Pow(2*math.Pi, 0.5) * math.Pow(sdar.sigma, 0.5)))
}

func randomFloat(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()*(max-min) + min
}
