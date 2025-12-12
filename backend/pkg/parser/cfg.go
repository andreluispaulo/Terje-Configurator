package parser

import (
	"regexp"
	"strings"
)

// LineType defines the type of a line in the CFG file
type LineType int

const (
	LineTypeUnknown LineType = iota
	LineTypeConfig
	LineTypeComment
	LineTypeEmpty
)

// Metadata holds the parsed info from comments
type Metadata struct {
	Type        string `json:"type"`
	Default     string `json:"default"`
	Description string `json:"description"`
}

// CFGLine represents a single line in the configuration file
type CFGLine struct {
	Index      int      `json:"index"`
	RawContent string   `json:"-"`
	Type       LineType `json:"type"`
	Key        string   `json:"key"`
	Value      string   `json:"value"`

	// Parts for reconstruction
	Prefix string `json:"-"`
	Suffix string `json:"-"`

	Metadata Metadata `json:"metadata"`
}

// CFGFile represents the parsed configuration file
type CFGFile struct {
	Lines []*CFGLine `json:"lines"`
}

// Regex to match config lines: Key = Value; // [type: ...; default: ...] Description
var configRegex = regexp.MustCompile(`^([^=]+=[ \t]*)([^;]+)(;.*)$`)
var metadataRegex = regexp.MustCompile(`//\s*\[type:\s*([^;]+);\s*default:\s*([^\]]+)\]\s*(.*)`)

// ParseCFG parses the content of a .cfg file
func ParseCFG(content string) (*CFGFile, error) {
	file := &CFGFile{Lines: make([]*CFGLine, 0)}
	lines := strings.SplitAfter(content, "\n")

	for i, lineStr := range lines {
		line := parseLine(lineStr, i)
		file.Lines = append(file.Lines, line)
	}

	return file, nil
}

func parseLine(raw string, index int) *CFGLine {
	trimmed := strings.TrimSpace(raw)
	line := &CFGLine{Index: index, RawContent: raw, Type: LineTypeUnknown}

	if trimmed == "" {
		line.Type = LineTypeEmpty
		return line
	}
	if strings.HasPrefix(trimmed, "//") {
		line.Type = LineTypeComment
		return line
	}

	// Handle newline stripping for regex
	hasNewline := strings.HasSuffix(raw, "\n")
	contentToMatch := raw
	if hasNewline {
		contentToMatch = raw[:len(raw)-1]
		if strings.HasSuffix(contentToMatch, "\r") {
			contentToMatch = contentToMatch[:len(contentToMatch)-1]
		}
	}

	matches := configRegex.FindStringSubmatch(contentToMatch)
	if matches != nil {
		line.Type = LineTypeConfig
		line.Prefix = matches[1]
		line.Value = strings.TrimSpace(matches[2])

		val := matches[2]
		trimmedVal := strings.TrimSpace(val)

		keyParts := strings.Split(line.Prefix, "=")
		line.Key = strings.TrimSpace(keyParts[0])

		line.Prefix = matches[1]

		trailingSpace := ""
		if len(val) > len(trimmedVal) {
			idx := strings.LastIndex(val, trimmedVal)
			if idx != -1 && idx+len(trimmedVal) < len(val) {
				trailingSpace = val[idx+len(trimmedVal):]
			}
		}

		line.Suffix = trailingSpace + matches[3]

		if hasNewline {
			if strings.HasSuffix(raw, "\r\n") {
				line.Suffix += "\r\n"
			} else {
				line.Suffix += "\n"
			}
		}

		parseMetadata(matches[3], line)
	}

	return line
}

func parseMetadata(commentPart string, line *CFGLine) {
	matches := metadataRegex.FindStringSubmatch(commentPart)
	if matches != nil {
		line.Metadata.Type = strings.TrimSpace(matches[1])
		line.Metadata.Default = strings.TrimSpace(matches[2])
		line.Metadata.Description = strings.TrimSpace(matches[3])
	}
}

// String reconstructs the file content
func (f *CFGFile) String() string {
	var sb strings.Builder
	for _, line := range f.Lines {
		if line.Type == LineTypeConfig {
			sb.WriteString(line.Prefix + line.Value + line.Suffix)
		} else {
			sb.WriteString(line.RawContent)
		}
	}
	return sb.String()
}
