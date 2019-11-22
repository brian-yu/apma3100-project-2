package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

type generator struct {
	count int
}

func NewGenerator() *generator {
	return &generator{count: 1}
}

func (g *generator) generate() float64 {
	p := g.rng(g.count)
	g.count += 1

	t := 57.0
	a := 1 / t

	return math.Sqrt(-2 * math.Log(1-p) / math.Pow(a, 2))
}

func (g *generator) rng(i int) float64 {
	multiplier := 24693
	increment := 3517
	modulus := math.Pow(2, 17)
	seed := 1000

	res := seed
	for j := 1; j <= i; j++ {
		res = (multiplier*res + increment) % int(modulus)
	}

	return float64(res) / modulus
}

func getSampleMean(n int, g *generator) float64 {
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += g.generate()
	}
	return sum / float64(n)
}

func main() {
	fmt.Println("Hello!")

	g := NewGenerator()

	fmt.Printf("u_1: %f\n", g.rng(1))
	fmt.Printf("u_2: %f\n", g.rng(2))
	fmt.Printf("u_3: %f\n", g.rng(3))
	fmt.Printf("u_51: %f\n", g.rng(51))
	fmt.Printf("u_52: %f\n", g.rng(52))
	fmt.Printf("u_53: %f\n", g.rng(53))

	err := os.Remove("go_log.csv")
	if err != nil {
		log.Println(err)
	}

	f, err := os.OpenFile("go_log.csv",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	n_vals := []int{10, 30, 50, 100, 150, 250, 500, 1000}
	z_vals := []float64{-1.4, -1.0, -.5, 0, .5, 1.0, 1.4}

	theta_vals := []float64{0.0808, 0.1587, 0.3085, 0.5, 0.6915, 0.8413, 0.9192}

	a := 1 / 57.0
	mu := (1 / a) * math.Sqrt(math.Pi/2.0)
	variance := (4 - math.Pi) / (2 * math.Pow(a, 2))
	std := math.Sqrt(variance)

	printCDF := func(n int, cdf []int) {
		for i, num := range cdf {
			fmt.Printf("%f,%f\n", z_vals[i], float64(num)/110.0)
		}
	}

	for _, n := range n_vals {

		cdf := make([]int, len(z_vals))

		fmt.Printf("n = %d\n", n)
		for i := 0; i < 110; i++ {
			m := getSampleMean(n, g)
			z := (m - mu) / (std / math.Sqrt(float64(n)))

			for j, z_j := range z_vals {
				if z <= z_j {
					cdf[j] += 1
				}
			}

			_, err := f.WriteString(fmt.Sprintf("%d,%g\n", n, m))
			if err != nil {
				log.Println(err)
			}
		}

		maxDiff := 0.0
		for i, num := range cdf {
			p := float64(num) / 110.0
			diff := math.Abs(p - theta_vals[i])
			maxDiff = math.Max(maxDiff, diff)
		}

		printCDF(n, cdf)
		fmt.Printf("MAD: %f\n", maxDiff)

	}
}
