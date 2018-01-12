package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)



type tl struct {
	filename  string
	cattype   string
	typeparse []string
	tlstring  string
}

func checktype(what string) (thing bool) {
	types := []string{"AREA", "BTLGRADE", "BTLTUTO", "EQUIP", "INN", "JUMPPOINT", "MAP", "RETURNPOINT", "TEST", "TOOLS", "STR"}
	for _, y := range types {
		if strings.Contains(what, y) {
			thing = true
		}
	}
	return thing
}

func parsefile(filepath string) (lines []tl) {
	fname := strings.Split(filepath, "/")
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	//var lines []tl
	line := &tl{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line.filename = fname[2]
		if checktype(scanner.Text()) {
			line.cattype = scanner.Text()
		} else {
			line.tlstring = scanner.Text()
		}

		if line.cattype != "" && line.tlstring != "" {
			lines = append(lines, *line)
			line = &tl{}

		}
	}
	return lines
}


func dirparse(searchDir string){
	fileList := []string{}
	filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})
	var lines []tl
		pad := 0

	for _, file := range fileList {
		if strings.Contains(file, ".txt") {
		for _, h := range parsefile(file) {
			lines = append(lines, h)
		}
	}
	}

	for p, h := range lines {
		lines[p].typeparse = strings.Split(strings.Replace(h.cattype, "=", "", -1), "_")

		if pad < len(lines[p].typeparse) {
			pad = len(lines[p].typeparse)

		}

	}
	for p, h := range lines {

		if pad > len(h.typeparse) {

			calc := pad - len(h.typeparse)
			for uu := 1; uu <= calc; uu++ {
				lines[p].typeparse = append(lines[p].typeparse, " ")
			}
		}
	}

	for _, u := range lines {
		fmt.Println(u.filename+","+u.cattype+","+strings.Join(u.typeparse, ",")+","+"\""+u.tlstring+"\"")
	}

}




func main() {
	if len(os.Args) > 1 {
		dirparse(os.Args[1])
	} else {
		fmt.Println("Usage : " + os.Args[0] + " input_file_name output_file_name")
		os.Exit(1)
	}
}
