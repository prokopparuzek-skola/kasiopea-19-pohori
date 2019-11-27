package main

import "fmt"

type hora struct {
	pozice int
	vyska  int
}

func reconstruct(pohori []hora) (denik []int) {
	denik = make([]int, len(pohori))
	for i, h := range pohori {
		var vzd1, vzd2 int
		var test1, test2 bool = false, false
		for j := i + 1; j < len(pohori); j++ {
			if pohori[j].vyska == h.vyska {
				vzd1 = pohori[j].pozice - h.pozice
				test1 = true
				break
			} else if pohori[j].vyska > h.vyska {
				vzd1 = h.vyska + (pohori[j].pozice - h.pozice - pohori[j].vyska)
				test1 = true
				break
			}
		}
		for j := i - 1; j >= 0; j-- {
			if pohori[j].vyska == h.vyska {
				vzd2 = h.pozice - pohori[j].pozice
				test2 = true
				break
			} else if pohori[j].vyska > h.vyska {
				vzd2 = h.vyska + (h.pozice - pohori[j].pozice - pohori[j].vyska)
				test2 = true
				break
			}
		}
		if !test1 && !test2 {
			denik[i] = -1
		} else {
			if !test1 {
				denik[i] = vzd2
			} else if !test2 {
				denik[i] = vzd1
			} else if vzd1 > vzd2 {
				denik[i] = vzd2
			} else {
				denik[i] = vzd1
			}
		}
	}
	return
}

func main() {
	var T int
	fmt.Scanf("%d", &T)
	for ; T > 0; T-- {
		var N int
		fmt.Scanf("%d", &N)
		pohori := make([]hora, N)
		for i := range pohori {
			fmt.Scanf("%d%d", &pohori[i].pozice, &pohori[i].vyska)
		}
		denik := reconstruct(pohori)
		for i := 0; i < len(denik)-1; i++ {
			fmt.Printf("%d ", denik[i])
		}
		fmt.Printf("%d\n", denik[len(denik)-1])
	}
}
