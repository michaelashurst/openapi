package main

import (
	"fmt"

	"github.com/michaelashurst/openapi/basicdocument"

	"github.com/thatisuday/commando"
)

func main() {
	commando.
		SetExecutableName("openapi").
		SetVersion("0.0.1").
		SetDescription("Test")

	// configure the root command
	commando.
		Register(nil).
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Printf("Help\n\n")
		})

	commando.
		Register("convert").
		AddArgument("type", "", "basic").
		AddArgument("path", "local input path", "./").
		SetShortDescription("Converts a basic api document to the standard openapi format").
		SetDescription("This command displays more information about the contents of the directory like size, permission and ownership, etc.").
		AddFlag("out", "local output path", commando.String, "./").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			if args["type"].Value == "basic" {
				fmt.Printf("Creating open api document\n")
				basic := basicdocument.NewBasicDocument(args["path"].Value)
				doc := basic.Document(flags["out"].Value.(string))
				doc.Save()
			}
		})

	// parse command-line arguments
	commando.Parse(nil)

}
