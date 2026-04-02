package cmd

import (
	"bytes"
	"strings"
)

var utf8BOM = []byte{0xEF, 0xBB, 0xBF}

// stripYAMLFrontmatter removes a leading YAML frontmatter block: the file must
// begin (after optional UTF-8 BOM) with a line that trims to "---", followed by
// a later line that trims to "---". If no closing delimiter exists, b is
// returned unchanged.
func stripYAMLFrontmatter(b []byte) []byte {
	orig := b
	b = bytes.TrimPrefix(b, utf8BOM)
	s := string(b)
	if !strings.HasPrefix(s, literalMarkdownTripleDash) {
		return orig
	}
	firstNL := strings.IndexByte(s, '\n')
	if firstNL < 0 {
		return orig
	}
	if strings.TrimSpace(s[:firstNL]) != literalMarkdownTripleDash {
		return orig
	}
	rest := s[firstNL+1:]
	restLines := strings.Split(rest, "\n")
	for i, line := range restLines {
		if strings.TrimSpace(line) == literalMarkdownTripleDash {
			body := strings.Join(restLines[i+1:], "\n")
			return []byte(body)
		}
	}
	return orig
}
