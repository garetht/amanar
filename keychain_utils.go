// +build darwin cgo

package main

import (
	"errors"
	"fmt"
	"github.com/keybase/go-keychain"
)

// Deletion poses problems because there does not seem to be a
// way to make deletion prompt for access, which means it will fail.
// trustedApplications is only set when creating a new keychain entry.
// Q: what happens when updating?
func CreateOrUpdateKeychainEntriesForService(service, account, password string, trustedApplications []string) error {
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService(service)
	item.SetMatchLimit(keychain.MatchLimitAll)
	item.SetReturnAttributes(true)

	results, err := keychain.QueryItem(item)

	if err != nil {
		return err
	}

	if len(results) == 0 {
		return CreateKeychainEntryForService(service, account, password, trustedApplications)
	}

	if len(results) > 1 {
		return errors.New(fmt.Sprintf("[KEYCHAIN] Found more than one entry for the service %s. Please delete duplicates and try again.", service))
	}

	for _, result := range results {
		originalItem := keychain.NewItem()
		originalItem.SetSecClass(keychain.SecClassGenericPassword)
		originalItem.SetService(result.Service)
		originalItem.SetAccount(result.Account)

		updateItem := keychain.NewItem()
		if account != "" {
			updateItem.SetAccount(account)
		}
		updateItem.SetData([]byte(password))

		// There should only ever be one result
		err = keychain.UpdateItem(originalItem, updateItem)
	}

	return err
}

// Creates a new keychain entry and
func CreateKeychainEntryForService(service, account, password string, trustedApplications []string) error {
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService(service)
	item.SetAccount(account)
	item.SetLabel(service)

	item.SetData([]byte(password))

	err := keychain.AddItem(item)

	if err != nil {
		return err
	}

	return nil
}
