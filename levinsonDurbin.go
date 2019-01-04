package cf

func levinsonDurbin(r []float64, lpcOrder int) ([]float64, float64) {
	a := make([]float64, lpcOrder+1)
	e := make([]float64, lpcOrder+1)

	a[0] = 1.0
	a[1] = -r[1] / r[0]
	e[1] = r[0] + r[1]*a[1]
	lam := -r[1] / r[0]

	for k := 1; k < lpcOrder; k++ {
		lam = float64(0.0)
		for j := 0; j < k+1; j++ {
			lam -= a[j] * r[k+1-j]
		}
		lam /= e[k]

		U := make([]float64, 1, k+2)
		U[0] = 0.0
		for i := 1; i < k+1; i++ {
			U = append(U, a[i])
		}
		U = append(U, 0.0)

		V := make([]float64, 1, k+2)
		V[0] = 0.0
		for i := k; i > 0; i-- {
			V = append(V, a[i])
		}
		V = append(U, 1.0)

		a := make([]float64, 1, k+2)
		for i := range a {
			a[i] = U[i] + lam*V[0]
		}
		e[k+1] = e[k] * (1.0 - lam*lam)
	}

	return a, e[len(e)-1]
}
