package main

import (
	"fmt"
)

func h(a float64, b float64, n int) float64 {
	return (b - a) / float64(n)
}

func divisor(a float64, b float64, n int) []float64 {
	aux := make([]float64, 0, n+1)
	step := (b - a) / float64(n)

	aux = append(aux, a)

	for i := 1; i < n; i++ {
		aux = append(aux, a+step*float64(i))
	}

	aux = append(aux, b)

	return aux
}

func trapeze(y []float64, h float64) []float64 {
	fmt.Printf("\n- Regra do Trapézio -\n")

	val := h / 2.0
	aux := 0.0

	for i, value := range y {
		if i == 0 {
			aux += value
		} else if i == len(y)-1 {
			aux += value
		} else {
			aux += 2 * value
		}
	}

	return []float64{val * aux, val}
}

func simpson1(y []float64, h float64) []float64 {
	if (len(y)-1)%2 != 0 {
		return []float64{0, 0}
	}

	fmt.Printf("\n- Primeira Regra de Simpson -\n")

	val := h / 3.0
	aux := 0.0

	for i, value := range y {
		if i == 0 {
			aux += value
		} else if i == len(y)-1 {
			aux += value
		} else if i%2 == 0 {
			aux += 2 * value
		} else {
			aux += 4 * value
		}
	}

	return []float64{val * aux, val}
}

func simpson2(y []float64, h float64) []float64 {
	if (len(y)-1)%3 != 0 {
		return []float64{0, 0}
	}

	fmt.Printf("\n- Segunda Regra de Simpson -\n")

	val := (3.0 * h) / 8.0
	aux := 0.0

	for i, value := range y {
		if i == 0 {
			aux += value
		} else if i == len(y)-1 {
			aux += value
		} else if i%3 == 0 {
			aux += 2 * value
		} else {
			aux += 3 * value
		}
	}

	return []float64{val * aux, val}
}

func main() {
	/*a := 0.7
	b := 2.0
	n := 7

	y := make([]float64, 0, n + 1)


	for _, x := range divisor(a, b, n) {
		y = append(y, math.Sin(x))
	}

	y1 := []float64{1.1778, 1.1176, 1.0265, 0.9093, 0.7709}
	y2 := []float64{0.7709, 0.2831, -0.196, -0.5351}

	h1 := h(0.7, 1.1, 4)
	h2 := h(1.1, 2.0, 3)

	fmt.Printf("%.7f", trapeze(y1, h1) + trapeze(y2, h2))*/

	/*a := 0.0
		b := math.Pi

		n1 := 1
		n2 := 2
		n3 := 3

		h1 := h(a, b, n1)
		h2 := h(a, b, n2)
		h3 := h(a, b, n3)

		y1 := make([]float64, 0, n1 + 1)
		y2 := make([]float64, 0, n2 + 1)
		y3 := make([]float64, 0, n3 + 1)

	    x1 := divisor(a, b, n1)
		x2 := divisor(a, b, n2)
		x3 := divisor(a, b, n3)

		for _, x := range x1 {
			y1 = append(y1, (math.Pow(x, 4)/4) + math.Pow(x, 2) + math.Sin(x))
		}

		for _, x := range x2 {
			y2 = append(y2, (math.Pow(x, 4)/4) + math.Pow(x, 2) + math.Sin(x))
		}

		for _, x := range x3 {
			y3 = append(y3, (math.Pow(x, 4)/4) + math.Pow(x, 2) + math.Sin(x))
		}


		fmt.Printf("\n- Regra do trapézio -\n" +
			              "n = %d h = %f val = %f I = %.7f\n",
						  n1, h2, trapeze(y1, h1)[1], trapeze(y1, h1)[0])

		fmt.Printf("x = %v\ny = %v\n", x1, y1)

		fmt.Printf("\n- Primeira Regra de Simpson -\n" +
			              "n = %d h = %f val = %f I = %.7f\n",
						  n2, h2, simpson1(y2, h2)[1], simpson1(y2, h2)[0])

		fmt.Printf("x = %v\ny = %v\n", x2, y2)

		fmt.Printf("\n- Segunda Regra de Simpson -\n" +
			              "n = %d h = %f val = %f I = %.7f\n",
						  n3, h3, simpson2(y3, h3)[1], simpson2(y3, h3)[0])

		fmt.Printf("x = %v\ny = %v\n", x3, y3)*/

	/*x1 := []float64{1.5, 1.8, 2.1, 2.4, 2.7, 3, 3.3}
	y1 := []float64{80, 75, 70, 65, 59, 53, 47}

	n1 := 6
	h1 := h(1.5, 3.3, n1)

	x2 := []float64{3.3, 3.5, 3.7, 3.9, 4.1, 4.3, 4.5}
	y2 := []float64{47, 44, 39, 35, 30, 26, 22}

	n2 := 6
	h2 := h(3.3, 4.5, n2)

	fmt.Printf("\n- Primeira Regra de Simpson -\n"+
		"n = %d h = %.4f val = %.4f I = %.4f\n",
		n1, h1, simpson1(y1, h1)[1], simpson1(y1, h1)[0])

	fmt.Printf("x = %v\ny = %v\n", x1, y1)

	fmt.Printf("\n- Primeira Regra de Simpson -\n"+
		"n = %d h = %.4f val = %.4f I = %.4f\n",
		n2, h2, simpson1(y2, h2)[1], simpson1(y2, h2)[0])

	fmt.Printf("x = %v\ny = %v\n", x2, y2)*/

	/*a := 30.0
	b := 45.0
	n := 3

	y := make([]float64, 0, n + 1)
	x := divisor(a, b, n)
	h := h(a, b, n)

	for _, x := range x {
		y = append(y, (1.0/30.0) * math.Pow(math.E, (-1.0/30.0) * x))
	}

	resp := simpson2(y, h)

	fmt.Printf("\n- Segunda Regra de Simpson -\n" +
		"n = %d h = %f val = %f I = %.7f\n",
		n, h, resp[1], resp[0])

	fmt.Printf("x = %.2v\ny = %.2v\n", x, y)*/

    /*a := -math.Sqrt(0.5)
	b := math.Sqrt(0.5)
	n := 5

	y := make([]float64, 0, n + 1)
	x := divisor(a, b, n)
	h := h(a, b, n)

	for _, x := range x {
		y = append(y, (1 / (1 + math.Pow(2 * x, 2))) - math.Pow(x, 2))
	}

	resp := trapeze(y, h)

	fmt.Printf("\n- Segunda do trapezio -\n" +
		"n = %d h = %f val = %f I = %.7f\n",
		n, h, resp[1], resp[0])

	fmt.Printf("x = %.4v\ny = %.4v\n", x, y)*/

	/*x1 := []float64{0, 0.5, 1, 1.5}
	y1 := []float64{62.0, 74.0, 73.5, 60.5}

	n1 := 3
	h1 := h(0, 1.5, n1)

	resp1 := trapeze(y1, h1)

	x2 := []float64{1.5, 48}
	y2 := []float64{60.5, 49.5}

	n2 := 1
	h2 := h(1.5, 48, n2)

	resp2 := trapeze(y2, h2)

	x3 := []float64{48, 48.5, 49}
	y3 := []float64{49.5, 42.5, 39.0}

	n3 := 2
	h3 := h(48, 49, n3)

	resp3 := trapeze(y3, h3)

	x4 := []float64{49, 59, 69, 79}
	y4 := []float64{39.0, 44.5, 58.0, 61.5}

	n4 := 3
	h4 := h(49, 79, n4)

	resp4 := trapeze(y4, h4)

	fmt.Printf("\n- Regra trapezio -\n"+
		"n = %d h = %.4f val = %.4f I = %.4f\n",
		n1, h1, resp1[1], resp1[0])

	fmt.Printf("x = %.4v\ny = %.4v\n", x1, y1)

	fmt.Printf("\n- Regra trapezio -\n"+
		"n = %d h = %.4f val = %.4f I = %.4f\n",
		n2, h2, resp2[1], resp2[0])

	fmt.Printf("x = %.4v\ny = %.4v\n", x2, y2)

	fmt.Printf("\n- Regra trapezio -\n"+
		"n = %d h = %.4f val = %.4f I = %.4f\n",
		n3, h3, resp3[1], resp3[0])

	fmt.Printf("x = %.4v\ny = %.4v\n", x3, y3)

	fmt.Printf("\n- Regra trapezio -\n"+
		"n = %d h = %.4f val = %.4f I = %.4f\n",
		n4, h4, resp4[1], resp4[0])

	fmt.Printf("x = %.4v\ny = %.4v\n", x4, y4)


	fmt.Printf("%f", resp1[0]+resp2[0]+resp3[0]+resp4[0])*/

	/*x1 := []float64{3,5,7,9,11,13,15}
	y1 := []float64{0.5, 0.5, 0.5, 1.5, 3.5, 3.0, 2.5}

	n1 := 6
	h1 := h(3, 15, n1)

	resp := simpson2(y1, h1)

	fmt.Printf("n = %d h = %.4f val = %.4f I = %.4f\n", n1, h1, resp[1], resp[0])

	fmt.Printf("x = %v\ny = %v\n", x1, y1)*/

	//a := 0.0
	//b := 30.0
	//n := 6
	//
	//y := make([]float64, 0, n + 1)
	//x := divisor(a, b, n)
	//h := h(a, b, n)
	//
	//for _, x := range x {
	//	y = append(y, 189+(0.11*math.Pow(x, 2)*math.Sqrt(30-x)))
	//}
	//
	//resp := trapeze(y, h)
	//
	//fmt.Printf("n = %d h = %f val = %f I = %.7f\n",
	//	n, h, resp[1], resp[0])
	//
	//fmt.Printf("x = %4.7v\ny = %4.7v\n", x, y)


	x := []float64{0,4,8,12,16,20,24}
	y := []float64{0,1.65,2.10,1.35,1.12,0.65,0}

	n := 6
	h := h(0, 24, n)

	resp := simpson2(y, h)

	fmt.Printf("n = %d h = %.4f val = %.4f I = %.4f\n", n, h, resp[1], resp[0])

	fmt.Printf("x = %v\ny = %v\n", x, y)

}
