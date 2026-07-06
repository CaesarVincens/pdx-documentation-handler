package main

import (
	"flag"
	"os"
	"slices"

	"bahmut.de/pdx-documentation-manager/cwt"
	"bahmut.de/pdx-documentation-manager/digest"
	"bahmut.de/pdx-documentation-manager/logging"
	"bahmut.de/pdx-documentation-manager/npp"
	"bahmut.de/pdx-documentation-manager/parser"
)

const (
	FlagVersion = "version"
)

func main() {
	logging.Info("Starting documentation handler")

	if slices.Contains(os.Args, "digest") {
		version := flag.String(FlagVersion, "", "Game Version the digest is for")
		flag.Parse()

		parsedVersion := "x.x.x"
		if version == nil || *version == "" {
			logging.Warnf("The parameter %s%s%s is missing.\n", logging.AnsiBoldOn, FlagVersion, logging.AnsiAllDefault)
		} else {
			parsedVersion = *version
		}

		logging.Info("Generating digest")
		err := digest.Generate(parsedVersion)
		if err != nil {
			logging.Fatal(err.Error())
			return
		}
	}

	if slices.Contains(os.Args, "cwt") {
		logging.Info("Generating CWT files")
		err := cwt.Generate()
		if err != nil {
			logging.Fatal(err.Error())
			return
		}
	}

	if slices.Contains(os.Args, "npp") {
		logging.Info("Generating Notepad++ language files")
		err := npp.Generate()
		if err != nil {
			logging.Fatal(err.Error())
			return
		}
	}

	if slices.Contains(os.Args, "json") {
		logging.Info("Generating JSON files")
		err := parser.GenerateJson()
		if err != nil {
			logging.Fatal(err.Error())
			return
		}
	}

}
