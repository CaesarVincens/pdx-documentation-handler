package npp

import (
	"encoding/xml"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	"bahmut.de/pdx-documentation-manager/parser"
)

const (
	documentationFolder  = "docs"
	outputFolder         = "out"
	nppFolder            = "npp"
	xmlBoolFalse         = "no"
	quote                = "&#34;"
	separatorPlaceholder = "||"
	separator            = "&#x000D;&#x000A;"
	separatorLength      = 16
	maxChar              = 30720
)

var currentKeyword = 1

func Generate() error {
	currentKeyword = 1
	err := generateMode("vic3_npp_language.xml", false)
	if err != nil {
		return err
	}

	currentKeyword = 1
	err = generateMode("vic3_npp_language-dark.xml", true)
	if err != nil {
		return err
	}

	return nil
}

func generateMode(file string, darkMode bool) error {
	err, nppPath := setupFolders()
	if err != nil {
		return err
	}

	nppXml := setupXml(darkMode)
	documentationPath := path.Join(documentationFolder, "new")
	documentation, err := parser.ParseScriptDocumentation(documentationPath)
	if err != nil {
		return err
	}

	if !darkMode {
		addKeywords(
			createEffectMapping(documentation),
			"8000FF",
			nppXml,
		)
		addKeywords(
			createTriggerMapping(documentation),
			"FF8000",
			nppXml,
		)
		addKeywords(
			createEventTargetMapping(documentation),
			"000080",
			nppXml,
		)
	} else {
		addKeywords(
			createEffectMapping(documentation),
			"EDD6ED",
			nppXml,
		)
		addKeywords(
			createTriggerMapping(documentation),
			"FF8040",
			nppXml,
		)
		addKeywords(
			createEventTargetMapping(documentation),
			"E3CEAB",
			nppXml,
		)
	}

	out, err := xml.MarshalIndent(nppXml, "", "\t")
	if err != nil {
		return err
	}
	outputString := strings.ReplaceAll(string(out), separatorPlaceholder, separator)
	outputString = strings.ReplaceAll(outputString, quote, "\"")
	err = os.WriteFile(path.Join(nppPath, file), []byte(outputString), 0644)
	if err != nil {
		return err
	}

	return nil
}

func addKeywords(elements []string, color string, nppXml *NotepadPlusPlus) {
	activeKeywords := &Keywords{
		Name: "Keywords" + strconv.Itoa(currentKeyword),
	}
	lineBuilder := strings.Builder{}
	lineLength := 0
	keywords := make([]*Keywords, 0)
	for _, element := range elements {
		if lineLength+len(element)+separatorLength > maxChar {
			activeKeywords.Text = lineBuilder.String()
			currentKeyword++
			keywords = append(keywords, activeKeywords)
			activeKeywords = &Keywords{
				Name: "Keywords" + strconv.Itoa(currentKeyword),
			}
			lineBuilder.Reset()
			lineLength = 0
		} else if lineLength != 0 {
			lineBuilder.WriteString(separatorPlaceholder)
			lineLength += separatorLength
		}
		lineBuilder.WriteString(element)
		lineLength += len(element)
	}
	activeKeywords.Text = lineBuilder.String()
	if activeKeywords.Text != "" {
		keywords = append(keywords, activeKeywords)
		currentKeyword++
	}
	for _, keyword := range keywords {
		nppXml.Language.KeywordLists.Keywords = append(nppXml.Language.KeywordLists.Keywords, keyword)
		nppXml.Language.Styles.WordsStyle = append(nppXml.Language.Styles.WordsStyle, &WordsStyle{
			Name:       strings.ToUpper(keyword.Name),
			FgColor:    color,
			BgColor:    "FFFFFF",
			ColorStyle: 1,
			FontStyle:  0,
			Nesting:    0,
		})
	}
}

func createEffectMapping(docs *parser.ScriptDocumentation) []string {
	keywords := make([]string, len(docs.EffectDocumentation.Elements))
	for index, effect := range docs.EffectDocumentation.Elements {
		keywords[index] = effect.ElementName()
	}
	for _, iterator := range docs.IteratorDocumentation.Elements {
		keywords = append(keywords, "every_"+iterator.ElementName())
		keywords = append(keywords, "random_"+iterator.ElementName())
		keywords = append(keywords, "ordered_"+iterator.ElementName())
	}
	return keywords
}

func createTriggerMapping(docs *parser.ScriptDocumentation) []string {
	keywords := make([]string, len(docs.TriggerDocumentation.Elements)+len(docs.IteratorDocumentation.Elements))
	for index, trigger := range docs.TriggerDocumentation.Elements {
		keywords[index] = trigger.ElementName()
	}
	for index, iterator := range docs.IteratorDocumentation.Elements {
		keywords[index+len(docs.EffectDocumentation.Elements)-1] = "any_" + iterator.ElementName()
	}
	return keywords
}

func createEventTargetMapping(docs *parser.ScriptDocumentation) []string {
	keywords := make([]string, len(docs.EventTargetDocumentation.Elements))
	for index, eventTarget := range docs.EventTargetDocumentation.Elements {
		keywords[index] = eventTarget.ElementName()
	}
	return keywords
}

func setupFolders() (error, string) {
	digestPath := path.Join(outputFolder, nppFolder)
	err := os.MkdirAll(digestPath, 0755)
	if err != nil {
		return fmt.Errorf("could not create output directory: %v", err), digestPath
	}
	return nil, digestPath
}

func setupXml(darkMode bool) *NotepadPlusPlus {
	npp := &NotepadPlusPlus{}
	npp.Language = &Language{
		Name:      "Vic3 Script",
		Extension: "TXT File",
		Version:   "2.1",
	}
	if darkMode {
		npp.Language.Name += " (Dark Mode)"
	}
	npp.Language.Settings = &Settings{
		Global: &SettingsGlobal{
			CaseIgnored:         xmlBoolFalse,
			AllowFoldOfComments: xmlBoolFalse,
			FoldCompact:         xmlBoolFalse,
			ForcePureLC:         0,
			DecimalSeparator:    0,
		},
		Prefix: &SettingsPrefix{
			Keywords1: xmlBoolFalse,
			Keywords2: xmlBoolFalse,
			Keywords3: xmlBoolFalse,
			Keywords4: xmlBoolFalse,
			Keywords5: xmlBoolFalse,
			Keywords6: xmlBoolFalse,
			Keywords7: xmlBoolFalse,
			Keywords8: xmlBoolFalse,
		},
	}
	npp.Language.KeywordLists = &KeywordLists{
		Keywords: []*Keywords{
			{
				Name: "Comments",
				Text: "00# 01 02 03 04",
			},
			{
				Name: "Delimiters",
				Text: "00$ 01 02$ 03\" 04 05\" 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20 21 22 23",
			},
			{
				Name: "Numbers, prefix1",
			},
			{
				Name: "Numbers, prefix2",
			},
			{
				Name: "Numbers, extras1",
			},
			{
				Name: "Numbers, extras2",
			},
			{
				Name: "Numbers, suffix1",
			},
			{
				Name: "Numbers, suffix2",
			},
			{
				Name: "Numbers, range",
			},
			{
				Name: "Folders in code1, open",
			},
			{
				Name: "Folders in code1, middle",
			},
			{
				Name: "Folders in code1, close",
			},
			{
				Name: "Folders in code2, open",
			},
			{
				Name: "Folders in code2, middle",
			},
			{
				Name: "Folders in code2, close",
			},
			{
				Name: "Folders in comment, open",
			},
			{
				Name: "Folders in comment, middle",
			},
			{
				Name: "Folders in comment, close",
			},
		},
	}
	if !darkMode {
		npp.Language.Styles = &Styles{
			WordsStyle: []*WordsStyle{
				{
					Name:       "DEFAULT",
					FgColor:    "004080",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "COMMENTS",
					FgColor:    "000000",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "LINE COMMENTS",
					FgColor:    "008000",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "NUMBERS",
					FgColor:    "DF0000",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "FOLDER IN CODE1",
					FgColor:    "FF8000",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "FOLDER IN CODE2",
					FgColor:    "000000",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "DELIMITERS1",
					FgColor:    "808080",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "DELIMITERS2",
					FgColor:    "808080",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "DELIMITERS3",
					FgColor:    "808080",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "DELIMITERS4",
					FgColor:    "808080",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "DELIMITERS5",
					FgColor:    "808080",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "DELIMITERS6",
					FgColor:    "808080",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "DELIMITERS7",
					FgColor:    "808080",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "DELIMITERS8",
					FgColor:    "808080",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "OPERATORS",
					FgColor:    "DFC47D",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "FOLDER IN COMMENT",
					FgColor:    "000000",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
			},
		}
	} else {
		npp.Language.Styles = &Styles{
			WordsStyle: []*WordsStyle{
				{
					Name:       "DEFAULT",
					FgColor:    "CEDF99",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "COMMENTS",
					FgColor:    "008000",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "LINE COMMENTS",
					FgColor:    "008000",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "NUMBERS",
					FgColor:    "8CD0D3",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "FOLDER IN CODE1",
					FgColor:    "DCDCCC",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "FOLDER IN CODE2",
					FgColor:    "DCDCCC",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "DELIMITERS1",
					FgColor:    "E3CEAB",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "DELIMITERS2",
					FgColor:    "E3CEAB",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "DELIMITERS3",
					FgColor:    "E3CEAB",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "DELIMITERS4",
					FgColor:    "E3CEAB",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "DELIMITERS5",
					FgColor:    "E3CEAB",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "DELIMITERS6",
					FgColor:    "E3CEAB",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "DELIMITERS7",
					FgColor:    "E3CEAB",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "DELIMITERS8",
					FgColor:    "E3CEAB",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "OPERATORS",
					FgColor:    "8080FF",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
				{
					Name:       "FOLDER IN COMMENT",
					FgColor:    "DCDCCC",
					BgColor:    "FFFFFF",
					ColorStyle: 1,
					FontStyle:  0,
					Nesting:    0,
				},
			},
		}
	}

	return npp
}
