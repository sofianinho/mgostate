package commands

import  "github.com/spf13/cobra"


// The collection state. Answers to the question: "does this collection exist inside this datbase in this mongodb server/cluster ?"
var cmdClt = &cobra.Command{
 Use:   "collection [collection to check]",
 Short: "Show the state of your collection in mongodb",
 Long:  `collection is the command that shows you your mongodb collection related information.`,
 Run: showState,
}


