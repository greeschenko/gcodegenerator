package main

import (
	"fmt"
	"os"
	"strings"
)

type product struct {
	id    int
	name  string
	stock int
}

var products []product

func listMenu() {
	fmt.Println("1. Create simple cutting off (one cut along X axis)")
	fmt.Println("2. Create cutting off with size (two cuts along X axis)")
	fmt.Println("3. Create cutting off with size (two cuts along X axis one direction)")
	fmt.Println("0. Exit")
	fmt.Print("Enter your choice: ")
}

func menu() {
	var choose int
	listMenu()
	for {
		fmt.Scan(&choose)
		if choose == 0 {
			fmt.Println("Thank You, exiting.....")
			break
		} else {
			switch choose {
			case 1:
				var (
					zsafe     float64
					tooldiam  float64
					cutlenght float64
					cutdepth  float64
					vfeed     float64
					hfeed     float64
					step      float64
					filename  string
					buf       strings.Builder
					depth     = 0.0
					isDirect  = true
				)

				fmt.Print("Z safe position: ")
				fmt.Scan(&zsafe)
				fmt.Print("choose tool diametr: ")
				fmt.Scan(&tooldiam)
				fmt.Print("choose cut lenght: ")
				fmt.Scan(&cutlenght)
				fmt.Print("choose cut depth: ")
				fmt.Scan(&cutdepth)
				fmt.Print("choose vertical feed: ")
				fmt.Scan(&vfeed)
				fmt.Print("choose horisontal feed: ")
				fmt.Scan(&hfeed)
				fmt.Print("choose step: ")
				fmt.Scan(&step)
				fmt.Print("choose gcode filename: ")
				fmt.Scan(&filename)

				fmt.Println("=============== RESULT ================")

				buf.WriteString("(Exported by greeschenko/gcodegenerator)\n")
				buf.WriteString("(Post Processor: grbl_post)\n")
				buf.WriteString("(Begin preamble)\n")
				buf.WriteString("G17 G90\n")
				buf.WriteString("G21\n")
				buf.WriteString("(Path: G54)\n")
				buf.WriteString(fmt.Sprintf("G0 Z%f\n", zsafe))
				buf.WriteString(fmt.Sprintf("(Compensated Tool Path. Diameter: %f)\n", tooldiam))
				buf.WriteString(fmt.Sprintf("G0 X%f Y%f\n", tooldiam*-1, -1*tooldiam/2))

				for depth >= -1*cutdepth {
					if isDirect {
						buf.WriteString(fmt.Sprintf("G1 X%f Y%f Z%f F%f\n", tooldiam*-1, -1*tooldiam/2, depth, vfeed))
						buf.WriteString(fmt.Sprintf("G1 X%f Y%f Z%f F%f\n", tooldiam+cutlenght, -1*tooldiam/2, depth, hfeed))
					} else {
						buf.WriteString(fmt.Sprintf("G1 X%f Y%f Z%f F%f\n", tooldiam+cutlenght, -1*tooldiam/2, depth, vfeed))
						buf.WriteString(fmt.Sprintf("G1 X%f Y%f Z%f F%f\n", tooldiam*-1, -1*tooldiam/2, depth, hfeed))
					}
					isDirect = !isDirect
					depth = depth - step
				}

				buf.WriteString(fmt.Sprintf("G0 Z%f\n", zsafe))
				buf.WriteString("(Begin postamble)\n")
				buf.WriteString("M5\n")
				buf.WriteString("G17 G90\n")
				buf.WriteString("M2\n")
				fmt.Print(buf.String())
				err := os.WriteFile(filename, []byte(buf.String()), 0644)
				if err != nil {
					panic(err)
				}

				fmt.Println("DONE!!!")
			case 2:
				var (
					zsafe     float64
					tooldiam  float64
					size      float64
					cutlenght float64
					cutdepth  float64
					vfeed     float64
					hfeed     float64
					step      float64
					filename  string
					buf       strings.Builder
					depth     = 0.0
					isDirect  = true
				)

				fmt.Print("Z safe position: ")
				fmt.Scan(&zsafe)
				fmt.Print("choose tool diametr: ")
				fmt.Scan(&tooldiam)
				fmt.Print("choose size between cuts: ")
				fmt.Scan(&size)
				fmt.Print("choose cut lenght: ")
				fmt.Scan(&cutlenght)
				fmt.Print("choose cut depth: ")
				fmt.Scan(&cutdepth)
				fmt.Print("choose vertical feed: ")
				fmt.Scan(&vfeed)
				fmt.Print("choose horisontal feed: ")
				fmt.Scan(&hfeed)
				fmt.Print("choose step: ")
				fmt.Scan(&step)
				fmt.Print("choose gcode filename: ")
				fmt.Scan(&filename)

				fmt.Println("=============== RESULT ================")

				buf.WriteString("(Exported by greeschenko/gcodegenerator)\n")
				buf.WriteString("(Post Processor: grbl_post)\n")
				buf.WriteString("(Begin preamble)\n")
				buf.WriteString("G17 G90\n")
				buf.WriteString("G21\n")
				buf.WriteString("(Path: G54)\n")
				buf.WriteString(fmt.Sprintf("G0 Z%f\n", zsafe))
				buf.WriteString(fmt.Sprintf("(Compensated Tool Path. Diameter: %f)\n", tooldiam))
				buf.WriteString(fmt.Sprintf("G0 X%f Y%f\n", tooldiam*-1, -1*tooldiam/2))

				for depth >= -1*cutdepth {
					if isDirect {
						buf.WriteString(fmt.Sprintf("G1 X%f Y%f Z%f F%f\n", tooldiam*-1, -1*tooldiam/2, depth, vfeed))
						buf.WriteString(fmt.Sprintf("G1 X%f Y%f Z%f F%f\n", tooldiam+cutlenght, -1*tooldiam/2, depth, hfeed))
					} else {
						buf.WriteString(fmt.Sprintf("G1 X%f Y%f Z%f F%f\n", tooldiam+cutlenght, -1*tooldiam/2, depth, vfeed))
						buf.WriteString(fmt.Sprintf("G1 X%f Y%f Z%f F%f\n", tooldiam*-1, -1*tooldiam/2, depth, hfeed))
					}
					isDirect = !isDirect
					depth = depth - step
				}

				buf.WriteString(fmt.Sprintf("G0 Z%f\n", zsafe))

				buf.WriteString(fmt.Sprintf("G0 X%f Y%f\n", tooldiam*-1, size + tooldiam/2))

                depth = 0.0

				for depth >= -1*cutdepth {
					if isDirect {
						buf.WriteString(fmt.Sprintf("G1 X%f Y%f Z%f F%f\n", tooldiam*-1, size + tooldiam/2, depth, vfeed))
						buf.WriteString(fmt.Sprintf("G1 X%f Y%f Z%f F%f\n", tooldiam+cutlenght, size + tooldiam/2, depth, hfeed))
					} else {
						buf.WriteString(fmt.Sprintf("G1 X%f Y%f Z%f F%f\n", tooldiam+cutlenght, size + tooldiam/2, depth, vfeed))
						buf.WriteString(fmt.Sprintf("G1 X%f Y%f Z%f F%f\n", tooldiam*-1, size + tooldiam/2, depth, hfeed))
					}
					isDirect = !isDirect
					depth = depth - step
				}

				buf.WriteString(fmt.Sprintf("G0 Z%f\n", zsafe))
				buf.WriteString("(Begin postamble)\n")
				buf.WriteString("M5\n")
				buf.WriteString("G17 G90\n")
				buf.WriteString("M2\n")
				fmt.Print(buf.String())
				err := os.WriteFile(filename, []byte(buf.String()), 0644)
				if err != nil {
					panic(err)
				}

				fmt.Println("DONE!!!")
			case 3:
				var (
					zsafe     float64
					tooldiam  float64
					size      float64
					cutlenght float64
					cutdepth  float64
					vfeed     float64
					hfeed     float64
					step      float64
					filename  string
					buf       strings.Builder
					depth     = 0.0
					isDirect  = true
				)

				fmt.Print("Z safe position: ")
				fmt.Scan(&zsafe)
				fmt.Print("choose tool diametr: ")
				fmt.Scan(&tooldiam)
				fmt.Print("choose size between cuts: ")
				fmt.Scan(&size)
				fmt.Print("choose cut lenght: ")
				fmt.Scan(&cutlenght)
				fmt.Print("choose cut depth: ")
				fmt.Scan(&cutdepth)
				fmt.Print("choose vertical feed: ")
				fmt.Scan(&vfeed)
				fmt.Print("choose horisontal feed: ")
				fmt.Scan(&hfeed)
				fmt.Print("choose step: ")
				fmt.Scan(&step)
				fmt.Print("choose gcode filename: ")
				fmt.Scan(&filename)

				fmt.Println("=============== RESULT ================")

				buf.WriteString("(Exported by greeschenko/gcodegenerator)\n")
				buf.WriteString("(Post Processor: grbl_post)\n")
				buf.WriteString("(Begin preamble)\n")
				buf.WriteString("G17 G90\n")
				buf.WriteString("G21\n")
				buf.WriteString("(Path: G54)\n")
				buf.WriteString(fmt.Sprintf("G0 Z%f\n", zsafe))
				buf.WriteString(fmt.Sprintf("(Compensated Tool Path. Diameter: %f)\n", tooldiam))
				buf.WriteString(fmt.Sprintf("G0 X%f Y%f\n", tooldiam*-1, -1*tooldiam/2))

				for depth >= -1*cutdepth {
					if isDirect {
						buf.WriteString(fmt.Sprintf("G1 X%f Y%f Z%f F%f\n", tooldiam*-1, -1*tooldiam/2, depth, vfeed))
						buf.WriteString(fmt.Sprintf("G1 X%f Y%f Z%f F%f\n", tooldiam+cutlenght, -1*tooldiam/2, depth, hfeed))
					} else {
						//buf.WriteString(fmt.Sprintf("G1 X%f Y%f Z%f F%f\n", tooldiam+cutlenght, -1*tooldiam/2, depth, vfeed))
						buf.WriteString(fmt.Sprintf("G1 X%f Y%f F%f\n", tooldiam*-1, -1*tooldiam/2, hfeed))
					}
					isDirect = !isDirect
					depth = depth - step
				}

				buf.WriteString(fmt.Sprintf("G0 Z%f\n", zsafe))

				buf.WriteString(fmt.Sprintf("G0 X%f Y%f\n", tooldiam*-1, size + tooldiam/2))

                depth = 0.0

				for depth >= -1*cutdepth {
					if isDirect {
						buf.WriteString(fmt.Sprintf("G1 X%f Y%f Z%f F%f\n", tooldiam*-1, size + tooldiam/2, depth, vfeed))
						buf.WriteString(fmt.Sprintf("G1 X%f Y%f Z%f F%f\n", tooldiam+cutlenght, size + tooldiam/2, depth, hfeed))
					} else {
						//buf.WriteString(fmt.Sprintf("G1 X%f Y%f Z%f F%f\n", tooldiam+cutlenght, size + tooldiam/2, depth, vfeed))
						buf.WriteString(fmt.Sprintf("G1 X%f Y%f F%f\n", tooldiam*-1, size + tooldiam/2, hfeed))
					}
					isDirect = !isDirect
					depth = depth - step
				}

				buf.WriteString(fmt.Sprintf("G0 Z%f\n", zsafe))
				buf.WriteString("(Begin postamble)\n")
				buf.WriteString("M5\n")
				buf.WriteString("G17 G90\n")
				buf.WriteString("M2\n")
				fmt.Print(buf.String())
				err := os.WriteFile(filename, []byte(buf.String()), 0644)
				if err != nil {
					panic(err)
				}

				fmt.Println("DONE!!!")
			default:
				fmt.Println("Re-enter your choice!")
				listMenu()
			}
		}
	}
}

func main() {
	menu()
}
