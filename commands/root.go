package commands

import (
 "github.com/spf13/cobra"
)

var url string //the mongodb url
var db string //name of the database
var clt string //name of the collection

// RootCmd is the root command of the CLI
var RootCmd = &cobra.Command{
 Use:   "mgostate",
 Short: "Test mongodb related parameters from shell",
 Long: `mgostate. 
   Helps you test mongodb parameters (like the existence of a database or a collection) from the shell.
         `,
}

func init() {
 RootCmd.AddCommand(cmdState)
 RootCmd.AddCommand(cmdServer)
 RootCmd.AddCommand(cmdDB)
 RootCmd.AddCommand(cmdClt)
 RootCmd.PersistentFlags().StringVarP(&url, "url", "u", "mongodb://localhost:27017/", "URI to connect to mongoDB")
 RootCmd.PersistentFlags().StringVarP(&db, "database", "d", "myTestDB", "Name of the database to check in mongoDB")
 RootCmd.PersistentFlags().StringVarP(&clt, "collection", "c", "myTestCollection", "Name of the collection to check in the database")
}

