package cmd

import (
	"github.com/ayuxsec/burp-xml-miner/internal/parser"
	"github.com/spf13/cobra"
)

var (
	xmlFile         string
	extractUrls     bool
	extractAll      bool
	extractResponse bool
	extractRequest  bool
	urlFilter       string
)

func NewExtractCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "extract <xml_file>",
		Short: "Extract objects from Burp Suite XML output",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			xmlFile = args[0]
			return runExtract()
		},
	}

	cmd.Flags().BoolVar(&extractUrls, "urls", false, "Extract URLs")
	cmd.Flags().BoolVar(&extractAll, "all", false, "Extract all objects")
	cmd.Flags().BoolVar(&extractRequest, "request", false, "Extract HTTP requests")
	cmd.Flags().BoolVar(&extractResponse, "response", false, "Extract HTTP responses")
	cmd.Flags().StringVar(&urlFilter, "url-filter", urlFilter, "Only give output if any part in url has this string as it's substring")

	return cmd
}

func runExtract() error {
	parser := parser.New(
		xmlFile,
		extractUrls,
		extractAll,
		extractRequest,
		extractResponse,
		urlFilter,
	)

	return parser.Print()
}
