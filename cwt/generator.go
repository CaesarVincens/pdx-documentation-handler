package cwt

import (
	"fmt"
	"os"
	"path"

	"bahmut.de/pdx-documentation-manager/parser"
)

const (
	documentationFolder = "docs"
	outputFolder        = "out"
	cwtFolder           = "cwt"
)

func Generate() error {
	err, cwtPath := setupFolders()
	if err != nil {
		return err
	}

	documentationPath := path.Join(documentationFolder, "new")
	documentation, err := parser.ParseScriptDocumentation(documentationPath)
	if err != nil {
		return err
	}

	iterators := PrintIterators(documentation)
	err = os.WriteFile(path.Join(cwtPath, "lists_generic.cwt"), []byte(iterators), 0644)
	if err != nil {
		return err
	}

	scriptedLists := PrintScriptedListEnum(documentation)
	err = os.WriteFile(path.Join(cwtPath, "list_base.cwt"), []byte(scriptedLists), 0644)
	if err != nil {
		return err
	}

	onActions := PrintOnActions(documentation)
	err = os.WriteFile(path.Join(cwtPath, "on_actions.cwt"), []byte(onActions), 0644)
	if err != nil {
		return err
	}

	modifiers := PrintModifiers(documentation)
	err = os.WriteFile(path.Join(cwtPath, "modifiers.cwt"), []byte(modifiers), 0644)
	if err != nil {
		return err
	}

	modifierCategories := PrintModifierCategories(documentation)
	err = os.WriteFile(path.Join(cwtPath, "modifier_categories.cwt"), []byte(modifierCategories), 0644)
	if err != nil {
		return err
	}

	scopes := PrintScopes(documentation)
	err = os.WriteFile(path.Join(cwtPath, "scopes.cwt"), []byte(scopes), 0644)
	if err != nil {
		return err
	}

	scopeEnum := PrintScopeEnum(documentation)
	err = os.WriteFile(path.Join(cwtPath, "enum_scopes.cwt"), []byte(scopeEnum), 0644)
	if err != nil {
		return err
	}

	return nil
}

func setupFolders() (error, string) {
	digestPath := path.Join(outputFolder, cwtFolder)
	err := os.MkdirAll(digestPath, 0755)
	if err != nil {
		return fmt.Errorf("could not create output directory: %v", err), digestPath
	}
	return nil, digestPath
}
