package main

import "fmt"

func DeleteLongKeys(m map[string]int) map[string]int {
	otv := make(map[string]int)
	for k, v := range m {
		if len(k) >= 6 {
			otv[k] = v
		}

	}

	return otv
}

func main() {
	m := map[string]string{
		"Яндекс":        "+74957397000",
		"Музей Яндекса": "+74991101133",
	}
	fmt.Println(SwapKeysAndValues(m))
}
