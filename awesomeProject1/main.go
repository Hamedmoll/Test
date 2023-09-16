package main

import (
	"awesomeProject1/entity"
	"awesomeProject1/memorystorage"
	"awesomeProject1/reader/consolereader"
	"awesomeProject1/repository/filerepository"
	"awesomeProject1/service/userservice"
	"awesomeProject1/writer/consolewriter"
	"flag"
	"fmt"
)

func main() {
	myrReader := consolereader.Reader{}
	myWriter := consolewriter.Writer{}

	var (
		myDataStruct = memorystorage.New(make(map[int]entity.Representation), make(map[string]int), make(map[string]int))
		myRepository = filerepository.New("Data.txt")
		uSrv         = userservice.New(&myDataStruct, myRepository, myrReader, myWriter)
	)

	representations, lErr := myRepository.Load()
	if lErr != nil {
		fmt.Println("cant load data ")
	}

	for _, rep := range representations {
		sErr := myDataStruct.Save(rep)
		if sErr != nil {
			fmt.Println("cant save data in memory")
		}
	}

	//fmt.Println(myDataStruct)

	pCommand := flag.String("command", "no command", "command to run")
	pRegion := flag.String("region", "no region", "select region")
	flag.Parse()

	command := *pCommand
	region := *pRegion

	for {
		//fmt.Println(command, region)

		rErr := runCommand(command, region, uSrv)
		if rErr != nil {
			fmt.Println(rErr)
		}

		command, region = uSrv.Reader.CommandRegion()
	}
}

func runCommand(command, region string, uSrv userservice.Service) error {
	//fmt.Println(command, region)
	//fmt.Println(uSrv)
	switch command {
	case "create":
		//fmt.Println("here1")
		req := userservice.CreateRequest{
			Region: region,
		}

		_, cErr := uSrv.CreateRepresentation(req)

		//fmt.Println("\n\n\n\n", uSrv.Memory)

		return cErr

	case "edit":
		req := userservice.EditRequest{
			Region: region,
		}

		_, edErr := uSrv.EditRepresentation(req)

		return edErr

	case "get":
		req := userservice.GetRequest{
			Region: region,
		}

		_, gErr := uSrv.GetByID(req)

		return gErr

	case "status":
		req := userservice.StatusRequest{
			Region: region,
		}

		_, sErr := uSrv.Status(req)

		return sErr

	case "list":
		req := userservice.ListRequest{Region: region}

		_, lErr := uSrv.ListByRegion(req)

		return lErr

	default:

		return fmt.Errorf("wrong command")
	}
}
