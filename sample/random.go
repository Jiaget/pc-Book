package sample

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.jiaget.com/pc-book/pb"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	return pb.Keyboard_Layout(rand.Intn(5) - 1)
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomStringFromSet(s ...string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	return s[rand.Intn(n)]
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Intel corei7 4790K",
			"Intel corei5 4590",
			"Intel corei3 4160",
			"Intel corei5 4460",
			"Intel corei7 4790",
		)
	}
	return randomStringFromSet(
		"Ryzen 3 Pro 1200",
		"Ryzen 3 1200",
		"Ryzen 3 Pro 1300",
		"Ryzen 3 1300X",
	)
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet(
			"GTX 1080",
			"GTX 1070",
			"Quadro P5000",
			"Tesla K40",
			"Tesla K40",
		)
	}
	return randomStringFromSet(
		"RX 6900 XT",
		"RX 6800",
		"Radeon VII",
		"R9 Fury X",
		"R9 380X",
	)
}

func randomScrennResolution() *pb.Screen_Resolution {
	height := randomInt(1000, 4320)
	width := height * 16 / 9

	resolution := &pb.Screen_Resolution{
		Height: uint32(height),
		Width:  uint32(width),
	}
	return resolution
}

func randomID() string {
	return uuid.New().String()
}

func randomLaptopBrand() string {
	return randomStringFromSet(
		"Apple",
		"Dell",
		"Lenovo",
	)
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet(
			"MacBook Air",
			"Macbook Pro",
		)
	case "Dell":
		return randomStringFromSet(
			"Insprion",
			"Vostro",
			"XPS",
			"Latitude",
		)
	default:
		return randomStringFromSet(
			"Thinkpad",
			"IdeaPad",
		)
	}
}
