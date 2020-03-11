//
// EPITECH PROJECT, 2020
// 201yams_2019
// File description:
// error
//

package functions

import "fmt"

type algo struct {
	nb1    float64
	nb2    float64
	endRes float64
	xlaw   [6]float64
	xinc   []int
	xtotal []float64
	ytotal []float64
	ztotal []float64
}

//PrintLine the line
func PrintLine(ret int) {
	if ret == 1 {
		fmt.Printf("\n--------------------------------------------------------------------------------\n")
	} else {
		fmt.Printf("--------------------------------------------------------------------------------\n")
	}
}

func (s *algo) SetupMath() {
	zloop := make([]int, 0)
	PrintLine(1)
	fmt.Printf("z")
	for nb := 20; nb <= 100; nb += 10 {
		fmt.Printf("\t%d", nb)
		zloop = append(zloop, nb)
	}
	fmt.Printf("\np(Z=z)")
	for z := 1; z != 10; z++ {
		s.ZLaw(z)
	}
	PrintLine(1)
	s.ExValVariance("X", s.xinc)
	s.ExValVariance("Y", s.xinc)
	s.ExValVariance("Z", zloop)
	PrintLine(0)
}

//Math algo
func Math(a float64, b float64) {
	s := algo{nb1: a, nb2: b, xlaw: [6]float64{0, 0, 0, 0, 0, 0},
		xinc: []int{10, 20, 30, 40, 50, 0}, endRes: 0}
	x := 0
	y := 0

	PrintLine(0)
	for i := 0; i < 6; i++ {
		x += 10
		if i != 5 {
			fmt.Printf("\tX=%d", x)
		} else {
			fmt.Printf("\tY law\n")
		}
	}
	for i := 0; i < 6; i++ {
		algoY := make([]float64, 6)
		fy := float64(y)
		y += 10
		algoY[i] += fy
		if i != 5 {
			fmt.Printf("Y=%d", y)
			s.JointLaw(algoY[i])
		} else {
			fmt.Printf("X law")
			s.JointLaw(algoY[i])
		}
	}
	s.SetupMath()
}
