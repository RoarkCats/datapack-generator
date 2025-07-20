package main

import (
	"embed"
	"fmt"
	"os"
	s "strings"
)

//go:embed templates/*
var templates embed.FS

var version = "2.0.0"

func main() {
	fmt.Printf("\n  -- Welcome to Datapack Generator %s by RoarkCats --\n", version)

	pack_name := Input("  Datapack Name: ")
	pack_namespace := Input("  Datapack Namespace: ")
	pack_author := Input("  Datapack Author: ")
	pack_ver := getVer()
	if pack_ver < 18 && s.ToLower(Input("  Import Rng? (y/n) ")) == "y" {
		// make rng
	}
	fmt.Printf("%s %s %s %d", pack_name, pack_namespace, pack_author, pack_ver)

	dat, err := os.ReadFile("./wah.txt")
	if err != nil {
		// panic(err)
	} else {
		fmt.Println(string(dat))
	}

	// for compilation
	// fmt.Scanf("%s")
}

// Get pack format from "version" (MC or format)
func getVer() (version int) {
	ver := Input("  Version: ")

	// basic format number
	_, err := fmt.Sscanf(ver, "%d ", &version)
	if err == nil {
		return version
	}

	var minor, patch int

	// '1.x' version
	_, err = fmt.Sscanf(ver, "1.%d ", &minor)
	if err == nil {
		version = map[int]int{
			13: 4, 14: 4,
			15: 5, 16: 5,
			17: 7,
			18: 8,
			19: 10,
			20: 15,
			21: 48,
		}[minor]
		return version
	}

	// 1.x.y version
	_, err = fmt.Sscanf(ver, "1.%d.%d ", &minor, &patch)
	if err == nil {
		version = map[int]int{
			1301: 4, 1302: 4, 1401: 4, 1402: 4, 1403: 4, 1404: 4,
			1501: 5, 1502: 5, 1601: 5,
			1602: 6, 1603: 6, 1604: 6, 1605: 6,
			1701: 7,
			1801: 8,
			1802: 9,
			1901: 10, 1902: 10, 1903: 10,
			1904: 12,
			2001: 15,
			2002: 18,
			2003: 26, 2004: 26,
			2005: 41, 2006: 41,
			2101: 48,
			2102: 57, 2103: 57,
			2104: 61,
			2105: 71,
			2106: 80,
			2107: 81, 2108: 81,
		}[minor*100+patch]
		return version
	}

	return 81
}
