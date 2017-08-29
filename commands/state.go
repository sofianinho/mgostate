package commands

import (
 "log"
 "github.com/sofianinho/mgostate/mgodb"
 "github.com/spf13/cobra"
 "os"
)


// The whole state. Answers to the question: "does this collection exist in this database at this url ?"
var cmdState = &cobra.Command{
 Use:   "state [state to check]",
 Short: "Show the state in your mongodb",
 Long:  `State is the command that shows you your mongodb related information.`,
 Run: showState,
}
func showState(cmd *cobra.Command, args []string){
 //log.Printf("Your command: %s, ARGS: %s, %s, %s ",strings.Join(args, " "),url,db,clt)
 log.SetOutput(os.Stdout)
 session,err := mgodb.ServerExists(url)
 if err != nil {
  log.Printf("[INFO][111] Server url \"%s\" is unreachable (means: server down or url malformed)", url)
 } else {
  defer session.Close()
  dbexists := mgodb.DbExists(session,db)
  if dbexists {
   sizeClt := mgodb.CollectionExists(session,db,clt)
   if sizeClt >= 0 {
    log.Printf("[INFO][000][%d] The collection \"%s\" exists with %d record(s), in the database \"%s\" at the url \"%s\"",sizeClt,clt, sizeClt,db,url)
   } else {
    log.Printf("[INFO][001] The collection \"%s\" does not exist in the database \"%s\" at the url \"%s\"",clt,db,url)
   }
  } else {
   log.Printf("[INFO][011] Database \"%s\" does not exist at \"%s\"",db,url)
  }
 }
}

