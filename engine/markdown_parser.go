package engine

import (
	"bufio"
	"os"
	"strings"
)

import (
	estrings "github.com/guangzhou-meta/go-lib/strings"
)

type markdownParser struct {
	dir          string
	templateName string
	filePath     string
	methods      map[string]string

	reader     *bufio.Reader
	lineNumber int64
	isEnd      bool
	inBody     bool

	penultimateLine string
	lastLine        string
}

type sqlSource struct {
	Id       string
	Template string
	Line     int64
}

func NewMarkdownParser(dir string, templateName string) *markdownParser {
	return &markdownParser{
		dir:          dir,
		templateName: templateName,
		filePath:     dir + templateName + ".md",
		methods:      make(map[string]string),
	}
}

func (parser *markdownParser) Parse() error {
	mdFile, err := os.Open(parser.filePath)
	if err != nil {
		return err
	}
	defer func(mdFile *os.File) {
		err := mdFile.Close()
		if err != nil {

		}
	}(mdFile)

	parser.reader = bufio.NewReader(mdFile)
	err = parser.skipHeader()
	if err != nil {
		return err
	}
	for {
		nextLine := parser.next()
		if nextLine == nil {
			break
		}
		parser.methods[nextLine.Id] = nextLine.Template
	}
	return nil
}

func (parser *markdownParser) skipHeader() error {
	for {
		line := parser.nextLine()
		if parser.isEnd ||
			strings.HasPrefix(*line, "===") {
			return nil
		}
	}
}

func (parser *markdownParser) next() *sqlSource {
	if parser.isEnd {
		return nil
	}
	sqlId := strings.Trim(parser.penultimateLine, " ")
	parser.skipComment()
	if parser.isEnd {
		return nil
	}
	sqlLine := parser.lineNumber
	sql := parser.readSql()
	source := &sqlSource{
		Id:       sqlId,
		Template: sql,
		Line:     sqlLine,
	}
	parser.inBody = false
	return source
}

func (parser *markdownParser) nextLine() *string {
	line, _, err := parser.reader.ReadLine()
	parser.lineNumber++
	if err != nil { // last line or read error
		parser.isEnd = true
		return nil
	}
	parser.penultimateLine = parser.lastLine
	lineStr := string(line)
	parser.lastLine = lineStr
	return &lineStr
}

func (parser *markdownParser) skipComment() {
	var findComment bool
	for {
		line := parser.nextLine()
		if parser.isEnd {
			return
		}
		l := strings.Trim(*line, " ")
		isEmptyLine := len(l) == 0
		if !findComment &&
			isEmptyLine {
			continue
		}
		if !parser.inBody &&
			strings.HasPrefix(l, "*") {
			// 注释
			findComment = true
			continue
		}
		if isEmptyLine {
			continue
		}
		if strings.HasPrefix(l, "```") ||
			strings.HasPrefix(l, "~~~") {
			parser.inBody = true
			continue
		} else {
			return
		}
	}
}

func (parser *markdownParser) readSql() string {
	list := []string{parser.lastLine}
	for {
		line := parser.nextLine()
		if parser.isEnd {
			return getBuildSql(list)
		}
		lineStr := *line
		if strings.HasPrefix(lineStr, "===") {
			list = list[0 : len(list)-1] // 删除下一个sqlId表示
			return getBuildSql(list)
		}
		list = append(list, lineStr)
	}
}

func getBuildSql(list []string) string {
	var elems []string
	sep := "\n"
	for _, str := range list {
		str = strings.Trim(str, " ")
		if strings.HasPrefix(str, "```") ||
			strings.HasPrefix(str, "~~~") { // 忽略以code block开头的符号
			continue
		}
		elems = append(elems, str)
	}
	if elems == nil {
		return ""
	}
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return elems[0]
	}
	// copy from strings.Join
	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += len(elems[i])
	}

	b := estrings.NewStringBuilder()
	b.Append(elems[0])
	for _, s := range elems[1:] {
		b.Append(sep).Append(s)
	}
	return b.String()
}

func ParseMarkdown(templateDir string, templateName string) (map[string]string, error) {
	markdownParser := NewMarkdownParser(templateDir, templateName)
	err := markdownParser.Parse()
	return markdownParser.methods, err
}
