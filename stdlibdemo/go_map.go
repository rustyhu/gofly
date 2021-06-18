package stdlibdemo

import "fmt"

func GoMap() {
	am := map[string]int{"kk": 1, "aa": 2}

	for _, s := range []string{"kk", "cc", "aa", "ksdfe"} {
		if v, ok := am[s]; ok {
			fmt.Printf("find key %v - value %v\n", s, v)
		} else {
			// no modification
			fmt.Printf("Looking up failed: %v - value %v\n", s, am[s])
		}
	}

	fmt.Println("After looking up -", am)
}
