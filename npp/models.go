package npp

import "encoding/xml"

type NotepadPlusPlus struct {
	XMLName  xml.Name `xml:"NotepadPlus"`
	Language *Language
}

type Language struct {
	XMLName      xml.Name `xml:"UserLang"`
	Name         string   `xml:"name,attr"`
	Extension    string   `xml:"ext,attr"`
	Version      string   `xml:"udlVersion,attr"`
	Settings     *Settings
	KeywordLists *KeywordLists
	Styles       *Styles
}

type KeywordLists struct {
	XMLName  xml.Name `xml:"KeywordLists"`
	Keywords []*Keywords
}

type Keywords struct {
	XMLName xml.Name `xml:"Keywords"`
	Name    string   `xml:"name,attr"`
	Text    string   `xml:",chardata"`
}

type Styles struct {
	XMLName    xml.Name `xml:"Styles"`
	WordsStyle []*WordsStyle
}

type WordsStyle struct {
	XMLName    xml.Name `xml:"WordsStyle"`
	Name       string   `xml:"name,attr"`
	FgColor    string   `xml:"fgColor,attr"`
	BgColor    string   `xml:"bgColor,attr"`
	ColorStyle int      `xml:"colorStyle,attr"`
	FontStyle  int      `xml:"FontStyle,attr"`
	Nesting    int      `xml:"Nesting,attr"`
}

type Settings struct {
	XMLName xml.Name `xml:"Settings"`
	Global  *SettingsGlobal
	Prefix  *SettingsPrefix
}

type SettingsGlobal struct {
	XMLName             xml.Name `xml:"Global"`
	CaseIgnored         string   `xml:"caseIgnored,attr"`
	AllowFoldOfComments string   `xml:"allowFoldOfComments,attr"`
	FoldCompact         string   `xml:"foldCompact,attr"`
	ForcePureLC         int      `xml:"forcePureLC,attr"`
	DecimalSeparator    int      `xml:"decimalSeparator,attr"`
}

type SettingsPrefix struct {
	XMLName   xml.Name `xml:"Prefix"`
	Keywords1 string   `xml:"Keywords1,attr"`
	Keywords2 string   `xml:"Keywords2,attr"`
	Keywords3 string   `xml:"Keywords3,attr"`
	Keywords4 string   `xml:"Keywords4,attr"`
	Keywords5 string   `xml:"Keywords5,attr"`
	Keywords6 string   `xml:"Keywords6,attr"`
	Keywords7 string   `xml:"Keywords7,attr"`
	Keywords8 string   `xml:"Keywords8,attr"`
}
