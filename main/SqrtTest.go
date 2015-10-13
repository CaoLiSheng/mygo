package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return "cannot Sqrt negative number: " + fmt.Sprint(float64(e))
}

func Sqrt(x float64) (float64, error) {
	if (x < 0) {
		return x, ErrNegativeSqrt(x)
	}

	z, zNew, delta := float64(1), 0., .001
	for !((z - zNew > 0 && z - zNew < delta) || (zNew - z > 0 && zNew - z < delta)) {
		if (zNew != 0.) {
			z = zNew
		}
		fmt.Println(z, zNew)
		zNew = z - (z * z - x) / (2 * z)
	}
	return zNew, nil
}

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(-4))
}
