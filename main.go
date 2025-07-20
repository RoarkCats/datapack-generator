package main

import (
	"embed"
	"fmt"
	"os"
	"strconv"
	s "strings"
)

//go:embed templates/*
var templates embed.FS

// Gets text content from embedded template file
func getTemplate(file string) (text string) {
	data, err := templates.ReadFile("templates/" + file)
	Check(err)
	return string(data)
}

var version = "2.0.0"

func main() {
	fmt.Printf("\n  -- Welcome to Datapack Generator %s by RoarkCats --\n\n", version)

	// --- Get datapack info
	pack_name := Input("  Datapack Name: ")
	pack_namespace := Input("  Datapack Namespace: ")
	pack_author := Input("  Datapack Author: ")
	pack_ver := getVer(Input("  Version: "))

	pack_ver_min := 0
	if pack_ver >= 18 {
		inp := Input("  Supported Version Range? : ")
		if inp != "" {
			pack_ver_min = getVer(inp)
			if pack_ver_min > pack_ver {
				pack_ver, pack_ver_min = pack_ver_min, pack_ver
			}
		}
	}

	// --- Create folders
	_s := "" // legacy folder name support
	if pack_ver < 48 {
		_s = "s"
	}
	Check(os.MkdirAll(pack_name+"/data/"+pack_namespace+"/function"+_s, 0755))
	Check(os.MkdirAll(pack_name+"/data/minecraft/tags/function"+_s, 0755))

	// --- Create files from templates
	fillTemplate := func(template string) string {
		template = s.ReplaceAll(template, "%name%", pack_name)
		template = s.ReplaceAll(template, "%namespace%", pack_namespace)
		template = s.ReplaceAll(template, "%author%", pack_author)
		template = s.ReplaceAll(template, "%format%", strconv.Itoa(pack_ver))
		template = s.ReplaceAll(template, "%min_format%", strconv.Itoa(pack_ver_min))
		return template
	} // helper funcs to get, fill, and write templates
	makePackTemplate := func(filedir string, template string) {
		Check(os.WriteFile(pack_name+"/"+filedir, []byte(fillTemplate(getTemplate(template))), 0644))
	}
	makePackTemplate("data/minecraft/tags/function"+_s+"/load.json", "load.json")
	makePackTemplate("data/minecraft/tags/function"+_s+"/tick.json", "tick.json")
	makePackTemplate("data/"+pack_namespace+"/function"+_s+"/load.mcfunction", "load.mcfunction")
	makePackTemplate("data/"+pack_namespace+"/function"+_s+"/main.mcfunction", "main.mcfunction")

	// --- New mcmeta formats
	if pack_ver_min == 0 {
		makePackTemplate("pack.mcmeta", "pack.mcmeta")
	} else {
		makePackTemplate("pack.mcmeta", "pack-range.mcmeta")
	}

	// --- Import legacy RNG
	if pack_ver < 18 && s.ToLower(Input("  Import Rng? (y/n) ")) == "y" {

		Check(os.MkdirAll(pack_name+"/data/"+pack_namespace+"/function"+_s+"/rng", 0755))

		AppendFile(pack_name+"/data/"+pack_namespace+"/function"+_s+"/load.mcfunction", "\n\nfunction "+pack_namespace+":rng/setup")

		makePackTemplate("data/"+pack_namespace+"/function"+_s+"/rng/lcg.mcfunction", "rng/lcg.mcfunction")
		makePackTemplate("data/"+pack_namespace+"/function"+_s+"/rng/next_int_lcg.mcfunction", "rng/next_int_lcg.mcfunction")
		makePackTemplate("data/"+pack_namespace+"/function"+_s+"/rng/range_lcg.mcfunction", "rng/range_lcg.mcfunction")
		makePackTemplate("data/"+pack_namespace+"/function"+_s+"/rng/setup.mcfunction", "rng/setup.mcfunction")

		uuid_reset := fillTemplate(getTemplate("rng/uuid_reset.mcfunction"))
		if pack_ver < 6 { // uuid reset source (super)legacy support
			uuid_reset = s.Replace(uuid_reset, "UUID[0] 1", "UUIDMost 0.00000000023283064365386962890625", 1)
		}
		Check(os.WriteFile(pack_name+"/"+"data/"+pack_namespace+"/function"+_s+"/rng/uuid_reset.mcfunction", []byte(uuid_reset), 0644))
	}

	// --- Done!
	fmt.Print("\n  Done!\n")

	// for compilation
	fmt.Scanf("%s")
}

// Prompt pack format from "version" (MC or format)
func getVer(ver string) (version int) {

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
	// else latest
	fmt.Println(" Invalid version! Selecting latest.")
	return 81
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
