package commands

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
	"github.com/joeymiller/trooper/requestcredentials"
	"github.com/mitchellh/cli"
)

type HeartbeatResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

type ServerRunCommand struct {
	Port string
	Host string
	Role string
	Ui   cli.Ui
}

func (s *ServerRunCommand) Run(args []string) int {

	cmdFlags := flag.NewFlagSet("server", flag.ContinueOnError)
	cmdFlags.StringVar(&s.Role, "role", "playground", "The role assumed by the credentials.")
	cmdFlags.StringVar(&s.Port, "port", "8080", "Server Port Number.")
	cmdFlags.StringVar(&s.Host, "host", "127.0.0.1", "Server Host.")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", s.Heartbeat)
	router.HandleFunc("/credentials", s.AWSTemporaryCredentials)

	s.Ui.Output(fmt.Sprintf("Running trooper in Server mode on %s:%s", s.Host, s.Port))

	log.Fatal(http.ListenAndServe(s.Host+":"+s.Port, router))

	return 0
}

func (s *ServerRunCommand) Heartbeat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(HeartbeatResponse{
		Status: "OK",
		Code:   200,
	})
}

func (s *ServerRunCommand) AWSTemporaryCredentials(w http.ResponseWriter, r *http.Request) {
	c, err := requestcredentials.Generate("playground")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	response, _ := json.Marshal(&struct {
		AccessKeyId     string `json:"accessKeyId"`
		SecretAccessKey string `json:"secretAccessKey"`
		SessionToken    string `json:"sessionToken"`
	}{
		*c.Credentials.AccessKeyId,
		*c.Credentials.SecretAccessKey,
		*c.Credentials.SessionToken,
	})

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}

func (s *ServerRunCommand) Help() string {
	return "Run trooper in Server Mode."
}

func (s *ServerRunCommand) Synopsis() string {
	return "Run trooper in Server Mode."
}
