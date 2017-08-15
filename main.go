package main

import (
	"log"
	"os"
)

func main() {
	//db, err := sql.Open("sqlite3", `/Users/garethtan/Library/Application Support/Querious/Connections.sqlite`)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table'")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//var tableName string
	//for rows.Next() {
	//	err = rows.Scan(&tableName)
	//	fmt.Println(tableName)
	//}
	//
	//rows, err = db.Query("SELECT uuid FROM connection_settings")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//var uuid string
	//for rows.Next() {
	//	err = rows.Scan(&uuid)
	//	fmt.Println(uuid)
	//}
	//
	//statement, err := db.Prepare("UPDATE connection_settings SET user=? where uuid=?")
	//statement.Exec("newuser14", "88879A23-9708-4CB4-8C39-5D88735A9DE2")

	executeAmanar()
}

func executeAmanar() {
	configItems, err, resultErrors := LoadConfiguration(os.Getenv("CONFIG_FILEPATH"), "amanar_config_schema.json")

	if err != nil {
		log.Fatalf("[CONFIG] Could not load configuration file: %s", err)
		return
	}

	if resultErrors != nil {
		log.Println("[CONFIG SCHEMA] The provided configuration JSON did not conform to the structure required by the JSON Schema.")
		for _, resultErr := range resultErrors {
			log.Printf("[CONFIG SCHEMA] At JSON location %s: %s", resultErr.Context().String(), resultErr.Description())
		}
		return
	}

	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		log.Fatalln("[GITHUB AUTH] Please provide a valid GitHub token as the environment variable GITHUB_TOKEN so we can fetch new credentials.")
		return
	}

	ghc := &VaultGithubAuthClient{
		GithubToken: githubToken,
	}
	err = ghc.loginWithGithub()
	if err != nil {
		log.Fatalf("[GITHUB AUTH] Could not log in with Github: %s", err)
		return
	}

	for _, configItem := range configItems {
		secret, err := ghc.getCredential(configItem.VaultPath, configItem.VaultRole)
		if err != nil {
			log.Printf("[VAULT AUTH] Could not retrieve secret for vault path %s and vault role %s because %s. Skipping.", configItem.VaultPath, configItem.VaultRole, err)
			continue
		}

		credentials, err := CreateCredentialsFromSecret(secret)

		if err != nil {
			log.Printf("[VAULT AUTH] Could not convert Vault secret into Amanar credentials because %s. Skipping.", err)
			continue
		}

		ProcessConfigItem(&configItem.Configurables, credentials)
	}
}
