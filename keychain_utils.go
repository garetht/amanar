package main

import "github.com/keybase/go-keychain"

// Deletes all keychain entries that belong to Service
func DeleteAllKeychainEntriesForService(service string) (err error) {
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService(service)

	for err != keychain.ErrorItemNotFound {
		err = keychain.DeleteItem(item)
	}

	if err == keychain.ErrorItemNotFound {
		return nil
	}
	return
}

// Creates a new keychain entry and
func CreateUniqueKeychainEntryForService(service, account, password string, trustedApplications []string) error {
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService(service)
	item.SetAccount(account)
	item.SetLabel(service)

	if len(trustedApplications) > 0 {
		access := keychain.Access{
			Label: "Amanar Trusted",
			TrustedApplications: trustedApplications,
		}
		item.SetAccess(&access)
	}

	item.SetData([]byte(password))

	// Remove all earlier entries for the service
	DeleteAllKeychainEntriesForService(service)
	err := keychain.AddItem(item)

	if err != nil {
		return err
	}

	return nil
}
