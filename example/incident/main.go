package main

import (
	"bufio"
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"syscall"

	"github.com/stpabhi/go-servicenow/servicenow"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	callerID         = "Foo Bar"
	shortDescription = "test alert"
	assignmentGroup  = "Foo Bar Service"
	description      = "This is a test alert generated by go-servicenow library"
	category         = "foo"
	subCategory      = "bar"
	cmdbCi           = "foobar"
	assignedTo       = "foo"
	location         = "foo bar"
	impact           = "2"
	urgency          = "2"
)

func main() {
	ctx := context.Background()

	// load tls certs
	cert, err := tls.LoadX509KeyPair("tls.crt", "tls.key")
	if err != nil {
		log.Fatal(err)
	}

	// basic auth. read username and password from stdin
	r := bufio.NewReader(os.Stdin)
	fmt.Print("ServiceNow Username: ")
	username, _ := r.ReadString('\n')

	fmt.Print("ServiceNow Password: ")
	bytePassword, _ := terminal.ReadPassword(syscall.Stdin)
	password := string(bytePassword)

	// basic auth with custom transport
	tp := servicenow.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{cert},
			},
			Proxy: http.ProxyFromEnvironment, // remove if not using proxy.
		},
	}

	// replace instance with real instance name
	client, err := servicenow.NewClient("https://instance.service-now.com", tp.Client())
	if err != nil {
		log.Fatal(err)
	}

	// Get an existing Incident
	inc, _, err := client.Incidents.Get(ctx, "INC12345678", servicenow.GetOptions{DisplayValue: "true"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(inc.GetAssignmentGroup())

	// Create a new Incident.
	i := &servicenow.Incident{
		CallerID:         &callerID,
		ShortDescription: &shortDescription,
		AssignmentGroup:  &assignmentGroup,
		Description:      &description,
		Category:         &category,
		Subcategory:      &subCategory,
		CmdbCi:           &cmdbCi,
		AssignedTo:       &assignedTo,
		Location:         &location,
		Impact:           &impact,
		Urgency:          &urgency,
	}

	inc, _, err = client.Incidents.Create(ctx, i, servicenow.CreateOptions{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Incident %s has been created with desc %s\n", inc.GetNumber(), inc.GetDescription())

	// List incidents
	incs, _, err := client.Incidents.List(ctx, servicenow.ListOptions{Limit: "10"})
	if err != nil {
		log.Fatal(err)
	}

	for _, inc := range incs {
		fmt.Println(inc.GetNumber())
	}

	// Update and resolve an existing incident.
	newCallerId := "Bar Foo"
	state := "6" // resolved
	i = &servicenow.Incident{
		CallerID:         &newCallerId,
		ShortDescription: &shortDescription,
		AssignmentGroup:  &assignmentGroup,
		Description:      &description,
		Category:         &category,
		Subcategory:      &subCategory,
		CmdbCi:           &cmdbCi,
		AssignedTo:       &assignedTo,
		Location:         &location,
		Impact:           &impact,
		Urgency:          &urgency,
		State:            &state,
	}
	inc, _, err = client.Incidents.Update(ctx, "INC12345678", i, servicenow.UpdateOptions{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Incident %s has been updated with desc %s\n", inc.GetNumber(), inc.GetCallerID())

}