package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

const (
	documentationFolder = "docs"
	outputFolder        = "out"
	jsonFolder          = "json"
)

func (d *ScriptDocumentation) ToJson() (string, error) {
	out, err := json.MarshalIndent(d, "", "\t")
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func (d *DataTypeDocumentation) ToJson() (string, error) {
	out, err := json.MarshalIndent(d, "", "\t")
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func GenerateJson() error {
	err, jsonPath := setupFolders()
	if err != nil {
		return err
	}

	documentationPath := path.Join(documentationFolder, "new")

	scriptDocumentation, err := ParseScriptDocumentation(documentationPath)
	if err != nil {
		return err
	}
	scriptJson, err := scriptDocumentation.ToJson()
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(jsonPath, "script-documentation.json"), []byte(scriptJson), 0644)
	if err != nil {
		return err
	}

	dataTypeDocumentation, err := ParseDataTypeDocumentation(documentationPath)
	if err != nil {
		return err
	}
	dataTypeJson, err := dataTypeDocumentation.ToJson()
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(jsonPath, "data-type-documentation.json"), []byte(dataTypeJson), 0644)
	if err != nil {
		return err
	}

	return nil
}

func setupFolders() (error, string) {
	digestPath := path.Join(outputFolder, jsonFolder)
	err := os.MkdirAll(digestPath, 0755)
	if err != nil {
		return fmt.Errorf("could not create output directory: %v", err), digestPath
	}
	return nil, digestPath
}
