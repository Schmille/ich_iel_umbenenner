package main

import (
	"io/fs"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

const (
	IchIEL      = "ich_iel"
	IchIELOne   = "ichiel"
	IchIELSpace = "ich iel"
	MeIRL       = "me_irl"
	MeIRLOne    = "meirl"
	MeIRLSpace  = "me irl"
)

func main() {
	dirpath := os.Args[1]

	if !fileExists(dirpath) {
		log.Fatalln("supplied path does not exist")
	}

	filepath.Walk(dirpath, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		name := info.Name()
		baseDir := path[:strings.Index(path, name)]
		extension := path[strings.Index(path, "."):]

		name = strings.ToLower(name)

		var newName string
		if isIchIEL(name) {
			newName = "ich_iel-"
		} else if isMeIRL(name) {
			newName = "me_irl-"
		} else {
			// Do not rename anything that is not iel
			return nil
		}

		for true {
			random := genRandom(5)
			possibleName := newName + random
			if !fileExists(possibleName) {
				newName += random
				break
			}
		}

		newPath := baseDir + newName + extension
		os.Rename(path, newPath)
		return nil
	})
}

func isIchIEL(name string) bool {
	return strings.Contains(name, IchIEL) || strings.Contains(name, IchIELOne) || strings.Contains(name, IchIELSpace)
}

func isMeIRL(name string) bool {
	return strings.Contains(name, MeIRL) || strings.Contains(name, MeIRLOne) || strings.Contains(name, MeIRLSpace)
}

func genRandom(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	out := make([]rune, length)
	for i := range out {
		out[i] = letters[rand.Intn(len(letters))]
	}
	return string(out)
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
