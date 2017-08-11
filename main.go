package main

import (
	"os"
	"log"
)

func main() {
	ghc := &VaultGithubAuthClient{
		GithubToken: os.Getenv("GITHUB_TOKEN"),
	}
	err := ghc.loginWithGithub()
	if err != nil {
		log.Fatal(err)
		return
	}

	//"./example/backup.xml"
	filename := "/Users/garethtan/Library/Preferences/DataGrip2017.2/projects/default/.idea/dataSources.local.xml"

	source := DatagripStores{
		DatabaseUUID: "1f76a8f4-2328-4ee3-ad77-d95ba60c645b",
		DatagripFilepath: filename,
		VaultPath: "db-tzanalytic",
		VaultRole: "read-only",
		AuthedClient: ghc,
	}

	err = source.RefreshCredentials()
	if err != nil {
		log.Fatal(err)
		return
	}

	source = DatagripStores{
		DatabaseUUID: "73d00a82-3e4d-4978-ae34-2bb04755429c",
		DatagripFilepath: filename,
		VaultPath: "db-peakpass",
		VaultRole: "read-only",
		AuthedClient: ghc,
	}

	err = source.RefreshCredentials()
	if err != nil {
		log.Fatal(err)
		return
	}

	//gac := VaultGithubAuthClient{
	//	GithubToken: os.Getenv("GITHUB_TOKEN"),
	//}
	//
	//err := gac.loginWithGithub()
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//
	//secret, err := gac.getCredential("db-tzanalytic", "read-only")
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//
	//file, err := ioutil.ReadFile("./example/example.xml")
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//datagripConfig, err := NewDatagripConfig(file)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//
	//datagripConfig.UpdateUsername("12dbe053-2f7b-40d8-9319-46764e69cfee", secret)
	//
	//fmt.Println(datagripConfig.Document.WriteToString())

	// Where is the service, account is the user
	//where := "AirPort"
	//account := "SeatGeek"
	//
	//pw, err := keyring.Get(where, account)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//fmt.Println(pw)
}
