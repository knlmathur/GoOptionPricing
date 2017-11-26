package CallOption

import (
	"fmt"
	"math"
	"math/rand"
)

// Struct with input variables
type OptInput struct {
	S0    float64 // Spot
	K     float64 // Strike
	T     float64 // Time
	R     float64 // Rate
	Sigma float64 // Vol
	I     int     // Number of Montecarlo Sims
}

// Function to sum values in an Array
func sumSlices(x []float64) float64 {

	totalx := 0.0
	for _, valuex := range x {
		totalx += valuex
	}
	return totalx
}

// Call Option Pricer
func (p *OptInput) call() float64 {

	//Generate an array of Random Numbers
	randnum := make([]float64, p.I)
	for i := 0; i < p.I; i++ {
		randnum[i] = rand.NormFloat64()
	}

	// Create Empty Arrays
	ST := make([]float64, p.I)
	hT := make([]float64, p.I)

	//Calc Raw Value
	for j := range randnum {
		ST[j] = p.S0 * math.Exp((p.R-0.5*math.Pow(p.Sigma, 2))*p.T+p.Sigma*math.Sqrt(p.T)*randnum[j])

		// Calc Time Value
		hT[j] = math.Max(ST[j]-p.K, 0.0)
	}

	//Calc average of values from Montecarlo
	C0 := math.Exp(-p.R*p.T) * (sumSlices(hT) / float64(p.I))

	return C0
}