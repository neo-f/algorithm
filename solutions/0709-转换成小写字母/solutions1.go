package main

func toLowerCase(s string) string {
	var ns []rune
	const delta = 'a' - 'A'
	for _, w := range []rune(s) {
		if w >= 'A' && w <= 'Z'{
			ns = append(ns, w + delta)
		}else{
			ns = append(ns, w)
		}
	}
	return string(ns)
}
