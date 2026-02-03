package parser

import (
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type BurpParser struct {
	FilePath      string
	IncludeURLs   bool
	IncludeReq    bool
	IncludeResp   bool
	FilterPattern string
}

func New(
	filePath string,
	includeURLs, all, includeReq, includeResp bool,
	filterPattern string,
) *BurpParser {
	parser := &BurpParser{
		FilePath:      filePath,
		IncludeURLs:   includeURLs,
		IncludeReq:    includeReq,
		IncludeResp:   includeResp,
		FilterPattern: filterPattern,
	}

	if all {
		parser.IncludeURLs = true
		parser.IncludeReq = true
		parser.IncludeResp = true
	}
	return parser
}

func (p *BurpParser) Print() error {
	if p.FilePath == "" {
		return errors.New("xml file path is required")
	}

	file, err := os.Open(p.FilePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	items, err := decodeXML(file)
	if err != nil {
		return fmt.Errorf("failed to parse XML: %w", err)
	}

	for _, item := range items.Items {
		if p.FilterPattern != "" && !strings.Contains(item.URL, p.FilterPattern) {
			continue
		}

		if p.IncludeURLs {
			fmt.Println(item.URL)
		}

		if p.IncludeReq {
			p.decodeAndPrint("\nREQUEST:\n", item.Request.Data)
		}

		if p.IncludeResp {
			p.decodeAndPrint("\nRESPONSE:\n", item.Response.Data)
		}
	}
	return nil
}

func (p *BurpParser) decodeAndPrint(label string, rawData string) {
	decoded, err := base64.StdEncoding.DecodeString(rawData)
	if err != nil {
		fmt.Printf("Error decoding %s: %+v\n", label, err)
		return
	}
	fmt.Printf("%s\n%s\n", label, string(decoded))
}

func decodeXML(r io.Reader) (Items, error) {
	var burpXMLItems Items
	decoder := xml.NewDecoder(r)
	err := decoder.Decode(&burpXMLItems)
	return burpXMLItems, err
}
