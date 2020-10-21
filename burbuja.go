package main

func Burbuja(s []int64) {
	acept := true
	var aux int64
	for i, val := range s {
		if i == 0 {
			aux = val
		}
		if aux != -100 {
			if aux > val {
				s[i-1] = s[i]
				s[i] = aux
				acept = false
			}
			if aux <= s[i] {
				aux = s[i]
			}
		}
	}
	if !acept {
		Burbuja(s)
	}
}

func main() {
}
