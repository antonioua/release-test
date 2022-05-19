package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
)

var (
	filePath string
	version  string
)

func main() {
	flag.StringVar(&filePath, "file", "changelog.md", "path to changelog file")
	flag.StringVar(&version, "version", "", "the new release version")
	flag.Parse()

	//const fileName = "./docs/docs/changelog.md"
	//const fileName = "changelog.md"
	const diffUrlPat = "https://github.com/TheWeatherCompany/cassandra-operator/compare/%s...%s"

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err, "Can't read file.")
	}

	start := bytes.Index(data, []byte("[Unreleased]"))
	end := bytes.Index(data[start:], []byte("\n"))
	prevUnreleased := data[start : start+end]
	//fmt.Println(string(prevUnreleased))
	newUnreleased := []byte(fmt.Sprintf("[Unreleased](%s)", fmt.Sprintf(diffUrlPat, version, "main")))
	//fmt.Println(string(newUnreleased))

	latestRelease := bytes.Replace(prevUnreleased, []byte("...main"), []byte("..."+version), 1)
	//fmt.Println(string(latestRelease))
	latestRelease = bytes.Replace(latestRelease, []byte("[Unreleased]"), []byte(fmt.Sprintf("## [%s]", version)), 1)
	//fmt.Println(string(latestRelease))

	newUnreleased = append(newUnreleased, append([]byte("\n\n"), latestRelease...)...)
	//fmt.Print(string(newUnreleased))
	data = bytes.Replace(data, prevUnreleased, newUnreleased, 1)
	fmt.Println(string(data))
}
