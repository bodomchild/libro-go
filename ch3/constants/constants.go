package constants

import "fmt"

const (
	_ = 1 << (10*iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

const (
	KB = 1000
	MB = KB * 1000
	GB = MB * 1000
	TB = GB * 1000
	PB = TB * 1000
	EB = PB * 1000
	ZB = EB * 1000
	YB = ZB * 1000
)

func Mostrar2() {
	fmt.Printf("%T %[1]v\n", KiB)
	fmt.Printf("%T %[1]v\n", MiB)
	fmt.Printf("%T %[1]v\n", GiB)
	fmt.Printf("%T %[1]v\n", TiB)
	fmt.Printf("%T %[1]v\n", PiB)
	fmt.Printf("%T %[1]v\n", EiB)
	//fmt.Printf("%T %[1]v\n", ZiB) overflows int
	//fmt.Printf("%T %[1]v\n", YiB) overflows int
}

func Mostrar10() {
	fmt.Printf("%T %[1]v\n", KB)
	fmt.Printf("%T %[1]v\n", MB)
	fmt.Printf("%T %[1]v\n", GB)
	fmt.Printf("%T %[1]v\n", TB)
	fmt.Printf("%T %[1]v\n", PB)
	fmt.Printf("%T %[1]v\n", EB)
	//fmt.Printf("%T %[1]v\n", ZB) // overflows int
	//fmt.Printf("%T %[1]v\n", YB) // overflows int
}