package commands

import (
 "log"
 "github.com/sofiane/mgostate/mgodb"
 "github.com/spf13/cobra"
 "os"
)


// The server state. Answers to the question: "is the mongodb server/cluster active at this url ?"
var cmdServer = &cobra.Command{
 Use:   "server [server to check]",
 Short: "Show the state of your mongodb server/cluster",
 Long:  `server is the command that shows you your mongodb server/cluster related information.`,
 Run: showServer,
}
func showServer(cmd *cobra.Command, args []string) {
 //log.Printf("Your command: %s, ARGS: %s",strings.Join(args, " "),url)
 log.SetOutput(os.Stdout)
 session,err := mgodb.ServerExists(url)
 if err != nil {
  log.Printf("[INFO][111] Server url \"%s\" is unreachable (means: server down or url malformed)", url)
 } else {
  defer session.Close()
  log.Printf("[INFO][011] The mongoDB server/cluser at \"%s\" is up",url)
 }
}

