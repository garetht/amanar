// +build !darwin !cgo

package amanar

// Compatibility functions for building with non-Darwin systems - assume that
// the keychain works as it should
func CreateOrUpdateKeychainEntriesForService(service, account, password string, trustedApplications []string) error {
	return nil
}

func CreateKeychainEntryForService(service, account, password string, trustedApplications []string) error {
	return nil
}
