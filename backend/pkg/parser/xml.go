package parser

import (
	"regexp"
	"strings"
)

// Segment represents a part of the line (either static text or an editable attribute)
type Segment struct {
	IsAttribute bool   `json:"isAttribute"`
	Content     string `json:"content"`             // For static text
	AttrName    string `json:"attrName,omitempty"`  // For attribute
	AttrValue   string `json:"attrValue,omitempty"` // For attribute
}

// XMLLine represents a line in the XML file
type XMLLine struct {
	Index    int       `json:"index"`
	Segments []Segment `json:"segments"`
	Depth    int       `json:"depth"`
	TagName  string    `json:"tagName"`
}

// XMLFile represents the parsed XML file
type XMLFile struct {
	Lines []*XMLLine `json:"lines"`
}

// Regex to find attributes: key="value"
var xmlAttributeRegex = regexp.MustCompile(`(\w+)="([^"]*)"`)

// Regex to find tag name: <TagName
var tagNameRegex = regexp.MustCompile(`<\/?([a-zA-Z0-9_]+)`)

// ParseXML parses the content of a .xml file line by line
func ParseXML(content string) (*XMLFile, error) {
	file := &XMLFile{Lines: make([]*XMLLine, 0)}
	lines := strings.SplitAfter(content, "\n")
	inComment := false

	for i, lineStr := range lines {
		line := &XMLLine{Index: i, Segments: make([]Segment, 0)}

		// Calculate Depth (indentation)
		trimmedLeft := strings.TrimLeft(lineStr, " \t")
		line.Depth = len(lineStr) - len(trimmedLeft)

		// Extract Tag Name
		tagMatch := tagNameRegex.FindStringSubmatch(trimmedLeft)
		if len(tagMatch) > 1 {
			line.TagName = tagMatch[1]
		}

		currentStr := lineStr

		for len(currentStr) > 0 {
			if inComment {
				// Look for closing -->
				idx := strings.Index(currentStr, "-->")
				if idx == -1 {
					// No closing tag, whole string is comment
					line.Segments = append(line.Segments, Segment{IsAttribute: false, Content: currentStr})
					currentStr = ""
				} else {
					// Found closing tag
					commentPart := currentStr[:idx+3]
					line.Segments = append(line.Segments, Segment{IsAttribute: false, Content: commentPart})
					currentStr = currentStr[idx+3:]
					inComment = false
				}
			} else {
				// Not in comment, look for opening <!--
				idx := strings.Index(currentStr, "<!--")
				if idx == -1 {
					// No comment start, process as normal XML
					segs := parseAttributes(currentStr)
					line.Segments = append(line.Segments, segs...)
					currentStr = ""
				} else {
					// Found comment start
					normalPart := currentStr[:idx]
					if len(normalPart) > 0 {
						segs := parseAttributes(normalPart)
						line.Segments = append(line.Segments, segs...)
					}

					// The rest starts with <!--
					currentStr = currentStr[idx:]
					inComment = true
				}
			}
		}

		file.Lines = append(file.Lines, line)
	}

	return file, nil
}

func parseAttributes(raw string) []Segment {
	segments := make([]Segment, 0)

	// Find all attribute matches with indices
	matches := xmlAttributeRegex.FindAllStringSubmatchIndex(raw, -1)

	if len(matches) == 0 {
		segments = append(segments, Segment{IsAttribute: false, Content: raw})
		return segments
	}

	lastIdx := 0
	for _, match := range matches {
		start := match[0]
		end := match[1]

		// Add static content before this match
		if start > lastIdx {
			segments = append(segments, Segment{
				IsAttribute: false,
				Content:     raw[lastIdx:start],
			})
		}

		keyStart, keyEnd := match[2], match[3]
		valStart, valEnd := match[4], match[5]

		key := raw[keyStart:keyEnd]
		val := raw[valStart:valEnd]

		segments = append(segments, Segment{
			IsAttribute: true,
			AttrName:    key,
			AttrValue:   val,
		})

		lastIdx = end
	}

	// Add remaining static content
	if lastIdx < len(raw) {
		segments = append(segments, Segment{
			IsAttribute: false,
			Content:     raw[lastIdx:],
		})
	}

	return segments
}

func parseXMLLine(raw string, index int) *XMLLine {
	// Deprecated, logic moved to ParseXML
	return nil
}

// String reconstructs the file content
func (f *XMLFile) String() string {
	var sb strings.Builder
	for _, line := range f.Lines {
		for _, seg := range line.Segments {
			if seg.IsAttribute {
				// Reconstruct attribute: name="value"
				sb.WriteString(seg.AttrName + `="` + seg.AttrValue + `"`)
			} else {
				sb.WriteString(seg.Content)
			}
		}
	}
	return sb.String()
}
