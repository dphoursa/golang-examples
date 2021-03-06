package main

import (
	"fmt"
	"math"
)

/**
 * The fraction 49/98 is a curious fraction, as an inexperienced mathematician in attempting 
 * to simplify it may incorrectly believe that 49/98 = 4/8, which is correct, is obtained by cancelling the 9s.
 *
 * We shall consider fractions like, 30/50 = 3/5, to be trivial examples.
 *
 * There are exactly four non-trivial examples of this type of fraction, less than one in value,
 * and containing two digits in the numerator and denominator.
 *
 * If the product of these four fractions is given in its lowest common terms, find the value of the denominator.
 */

type Fraction struct {
	n, d int
}

func (f *Fraction) multiplyBy(multiplier Fraction) {
	f.n *= multiplier.n
	f.d *= multiplier.d
}

func (f *Fraction) divideBy(divider int) {
	f.n /= divider
	f.d /= divider
}

func (f *Fraction) simplify() {
	for i := f.n; i > 1; i-- {
		for math.Mod(float64(f.n), float64(i)) == 0 && 
			math.Mod(float64(f.d), float64(i)) == 0 {
			f.divideBy(i)
		}
	}
}

func main() {

	fraction := Fraction{1, 1}

	for n := 1; n <= 8; n++ {
		for d := 2; d <= 9; d++ {
			if d <= n {
				continue
			}

			for x := 1; x <= 9; x ++ {

				xn := x * 10 + n;
				xd := x * 10 + d;
				nx := n * 10 + x;
				dx := d * 10 + x;

				if  (xn < xd && xn * d == xd * n) ||
					(xn < dx && xn * d == dx * n) ||
					(nx < xd && nx * d == xd * n) || 
					(nx < dx && nx * d == dx * n) {

					multiplier := Fraction{n, d}
					fraction.multiplyBy(multiplier)
				}
			}
		}
	}

	fraction.simplify();

	fmt.Println(fraction.n, " / ", fraction.d)
}

