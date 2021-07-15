package main

func main() {
	m := make(map[string]string)
	m["a"] = "q"
	m["b"] = "w"
	m["c"] = "e"
	m["d"] = "r"
	for k, v := range m {
		m[k] = v + "te"
		println(m[k])
	}
}
