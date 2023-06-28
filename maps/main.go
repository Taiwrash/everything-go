package main

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#ff00ff",
		"white": "#ffffff",
	}

	print(colors)
}

func print(d map[string]string) {
	for color, hex := range d {
		println("color ", color, "its hex is ", hex)
	}
}
