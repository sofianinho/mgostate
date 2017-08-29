package commands

import (
 "log"
 "github.com/sofianinho/mgostate/mgodb"
 "github.com/spf13/cobra"
 "os"
)


// The database state. Answers to the question: "does this databse exist inside this mongodb server/cluster ?"
var cmdDB = &cobra.Command{
 Use:   "database [database to check]",
 Short: "Show the state of your database in mongodb",
 Long:  `database is the command that shows you your mongodb database related information.`,
 Run: showDB,
}
func showDB(cmd *cobra.Command, args []string) {
 //log.Printf("Your command: %s, ARGS: %s, %s ",strings.Join(args, " "),url,db)
 log.SetOutput(os.Stdout)
 session,err := mgodb.ServerExists(url)
 if err != nil {
  log.Printf("[INFO][111] Server url \"%s\" is unreachable (means: server down or url malformed)", url)
 } else {
  defer session.Close()
  dbexists := mgodb.DbExists(session,db)
  if dbexists {
   log.Printf("[INFO][001] Database \"%s\" exists at \"%s\"",db,url)
  } else {
   log.Printf("[INFO][011] Database \"%s\" does not exist at \"%s\"",db,url)
  }
 }
}

