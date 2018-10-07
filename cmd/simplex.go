package buildings

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"

	"github.com/tbogdala/noisey"
)

func Test(seed int) {
	// r := rand.New(rand.NewSource(int64(1)))
	// perlin := noisey.NewPerlinGenerator(r)
	// v := perlin.Get2D(0.4, 0.2)
	// fmt.Printf("Perlin: %v\n", v)

	// simplex := noisey.NewOpenSimplexGenerator(r)
	// x := 10
	// y := 10
	// for i := 0; i <= x; i++ {
	// 	for h := 0; h <= y; h++ {
	// 		floath := float64(h)
	// 		floati := float64(i)
	// 		b := simplex.Get2D(floati, floath)
	// 		fmt.Printf("Cord: %v,%v Value: %v\n", i, h, b)
	// 	}
	// }

	fmt.Println("Doing a noisey test.")

	const imageSize int = 50

	fmt.Printf("Seed: %v\n", seed)
	// make a test generator seeded to 1
	r := rand.New(rand.NewSource(int64(seed)))

	// create a new perlin noise generator
	perlin := noisey.NewPerlinGenerator(r)
	// simplex := noisey.NewOpenSimplexGenerator(r)

	// create the fractal Brownian motion generator based on perlin
	fbmPerlin := noisey.NewFBMGenerator2D(&perlin, 1.0, 0.5, 2.0, 1.0)

	// make an ascii pixel image by calculating random noise
	pixels := make([]float64, imageSize*imageSize)
	for y := 0; y < imageSize; y++ {
		for x := 0; x < imageSize; x++ {
			v := fbmPerlin.Get2D(float64(x)*0.1, float64(y)*0.1)
			v = v*0.5 + 0.5
			pixels[y*imageSize+x] = v
		}
	}

	// print the image out to the terminal using the symbols below
	symbols := []string{" ", "░", "▒", "▓", "█", "█"}
	out := bufio.NewWriter(os.Stdout)
	for y := 0; y < imageSize; y++ {
		for x := 0; x < imageSize; x++ {
			fmt.Fprint(out, symbols[int(pixels[y*imageSize+x]/0.2)])
		}
		fmt.Fprintln(out)
	}
	out.Flush()

}
