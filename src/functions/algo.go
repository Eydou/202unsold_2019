//
// EPITECH PROJECT, 2020
// 202unsold_2019
// File description:
// algo
//

package functions

import (
	"fmt"
	"math"
)

func (s *algo) AlgoLaw(algoX float64, algoY float64) float64 {
	return ((s.nb1 - algoX) * (s.nb1 - algoY)) / ((5*s.nb1 - 150) * (5*s.nb2 - 150))
}

func (s *algo) ExValVariance(c string, loop []int) {
	expectV := float64(0)
	variance := float64(0)
	varTotal := make([]float64, 6)

	switch c {
	case "X":
		varTotal = s.xtotal
	case "Y":
		varTotal = s.ytotal
	case "Z":
		varTotal = s.ztotal
	}
	for i := 0; i != len(varTotal); i++ {
		expectV += (float64(loop[i]) * varTotal[i])
		variance += (math.Pow(float64(loop[i]), 2) * varTotal[i])
	}
	variance -= (math.Pow(expectV, 2))
	fmt.Printf("expected value of %s:\t%.1f\n", c, expectV)
	fmt.Printf("variance of %s:\t\t%.1f\n", c, variance)
}

func (s *algo) ZLaw(z int) {
	zres := float64(0)

	for y := 1; y != 6; y++ {
		algoY := float64(y) * 10
		for x := 0; x != 5; x++ {
			algoX := float64(x) * 10
			algo := s.AlgoLaw(algoX, algoY)
			if x+y == z && y < 6 {
				zres += algo
			}
		}
	}
	s.ztotal = append(s.ztotal, zres)
	fmt.Printf("\t%0.3f", zres)
}

func (s *algo) JointLaw(algoY float64) {
	y := float64(0)
	ylaw := float64(0)

	for i := 0; i < 6; i++ {
		algoX := make([]float64, 6)
		y += 10
		algoX[i] += y

		if i != 5 && algoY != 50 {
			fmt.Printf("\t%.3f", s.AlgoLaw(algoX[i], algoY))
			ylaw += s.AlgoLaw(algoX[i], algoY)
		} else if algoY != 50 {
			s.ytotal = append(s.ytotal, ylaw)
			fmt.Printf("\t%.3f\n", ylaw)
			s.endRes += ylaw
		}
		if s.xinc[i] == int(algoX[i]) {
			if algoY != 50 {
				s.xlaw[i] += s.AlgoLaw(algoX[i], algoY)
			}
			if algoY == 50 {
				s.xtotal = append(s.xtotal, s.xlaw[i])
				fmt.Printf("\t%.3f", s.xlaw[i])
			}
		}
		if algoY == 50 && int(algoX[i]) == 50 {
			fmt.Printf("\t%.3f", s.endRes)
		}
	}
}
