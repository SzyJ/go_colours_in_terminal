package main

import "fmt"

// Terminal Colour escapes
var HEADER string     = "\033[95m"   //   \033[95m
var OKBLUE string     = "\033[94m"   //   \033[94m
var OKGREEN string    = "\033[92m"   //   \033[92m
var WARNING string    = "\033[93m"   //   \033[93m
var FAIL string       = "\033[91m"   //   \033[91m
var ENDC string       = "\033[0m"    //   \033[0m
var BOLD string       = "\033[1m"    //   \033[1m
var UNDERLINE string  = "\033[4m"    //   \033[4m

func printColours() {
	fmt.Println(HEADER    + "HEADER"    + ENDC + ":    \t " + "033[95m")
	fmt.Println(OKBLUE    + "OKBLUE"    + ENDC + ":    \t " + "033[94m")
	fmt.Println(OKGREEN   + "OKGREEN"   + ENDC + ":    \t " + "033[92m")
	fmt.Println(WARNING   + "WARNING"   + ENDC + ":    \t " + "033[93m")
	fmt.Println(FAIL      + "FAIL"      + ENDC + ":    \t " + "033[91m")
	fmt.Println(ENDC      + "ENDC"      + ENDC + ":    \t " + "033[0m")
	fmt.Println(BOLD      + "BOLD"      + ENDC + ":    \t " + "033[1m")
	fmt.Println(UNDERLINE + "UNDERLINE" + ENDC + ":    \t " + "033[4m")
}

func main() {
	printColours();
}
