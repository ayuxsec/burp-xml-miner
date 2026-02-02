package parser

import (
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
)

type Parser struct {
	XMLFile  string
	URLs     bool
	Request  bool
	Response bool
}

func New(
	xmlFile string,
	urls bool,
	all bool,
	request bool,
	response bool,
) *Parser {

	miner := &Parser{
		XMLFile:  xmlFile,
		URLs:     urls,
		Request:  request,
		Response: response,
	}

	if all {
		miner.URLs = true
		miner.Request = true
		miner.Response = true
	}

	return miner
}

// todo: verbose errors
func (p *Parser) Print() error {
	if err := p.validate(); err != nil {
		return fmt.Errorf("error validating args: %+v", err)
	}
	data, err := os.ReadFile(p.XMLFile)
	if err != nil {
		return err
	}
	items, err := parseXML(data)
	if err != nil {
		return err
	}

	for _, i := range items.Items {
		if p.URLs {
			fmt.Println(i.URL)
		}
		if p.Request {
			decodeBytes, _ := base64.StdEncoding.DecodeString(i.Request.Data)
			fmt.Println(string(decodeBytes))
		}
		if p.Response {
			decodeBytes, _ := base64.StdEncoding.DecodeString(i.Response.Data)
			fmt.Println(string(decodeBytes))
		}
	}
	return nil
}

func (p *Parser) validate() error {
	if p.XMLFile == "" {
		return errors.New("xml file is required")
	}
	return nil
}

func parseXML(data []byte) (burpXMLItems Items, err error) {
	err = xml.Unmarshal(data, &burpXMLItems)
	return
}
