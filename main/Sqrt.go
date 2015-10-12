package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z, zNew, delta := float64(1), 0., .001
	for !((z - zNew > 0 && z - zNew < delta) || (zNew - z > 0 && zNew - z < delta)) {
		if (zNew != 0.) {
			z = zNew
		}
		fmt.Println(z, zNew)
		zNew = z - (z * z - x) / (2 * z)
	}
	return zNew
}

func main() {
	fmt.Println(int(Sqrt(4)))
}
