package main

import (
	"fmt"
	"os"
	"strings"
)

type Preprocessor struct {
	content    string
	lineNumber int
}

func NewPreprocessor(content string) Preprocessor {
	return Preprocessor{content: content, lineNumber: 0}
}
func (self *Preprocessor) Preprocess() []string {
	out := []string{}
	self.content = RemoveBetween(self.content, `\/\*`, `\*\/`)
	for lineNum, line := range strings.Split(self.content, "\n") {
		if strings.HasPrefix(line, "@import") {
			//import call
			parts := strings.Split(line, " ")
			if len(parts) != 2 {
				Fatal(fmt.Sprintf("(preproc) Import call on line %d just requires path", lineNum))
			}
			path := parts[1]
			bytes, err := os.ReadFile(path)
			if err != nil {
				Fatal(fmt.Sprintf("(preproc) Error reading file at import statement on line %d", lineNum))
			}
			// Preprocess the file
			preprocessor := NewPreprocessor(string(bytes))
			lines := preprocessor.Preprocess()
			self.content = strings.Join(append(lines, self.content), "\n")
			// put the content from this file at the top of the file that's imported
		}
		self.lineNumber = lineNum + 1
		if strings.HasPrefix(line, "//") || strings.HasPrefix(line, "#") {
			continue
		}
		line = strings.Split(line, `//`)[0]
		line = strings.Trim(line, "\n")
		line = strings.TrimSpace(line)
		line = strings.ReplaceAll(line, ";", "")
		out = append(out, line)
	}
	out = DeleteEmptyFromSlice(out) // causes OOB errors
	return out
}
