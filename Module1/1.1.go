package main

func main() {
	var arr = [...]string{"I", "am", "stupid", "and", "weak"}

	for i, v := range arr {
		if v == "stupid" {
			arr[i] = "smart"
		}

		if v == "weak" {
			arr[i] = "strong"
		}
	}
}
