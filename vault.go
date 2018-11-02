package main

// import (
// 	"log"
// 	"regexp"
// 	"strings"

// 	"github.com/vinc3m1/opvault"
// )

// type vault struct {
// 	vaultSetting VaultSetting
// 	vault        *opvault.Vault
// }

// type vaultProfile struct {
// 	vaultProfile *opvault.Profile
// }

// // NewVault creates a new vault
// func NewVault(settings VaultSetting) *vault {
// 	v, err := opvault.Open(settings.GetPath())
// 	if err != nil {
// 		Panic(err, "Did you provide a valid path to the opvault folder?")
// 	}
// 	return &vault{
// 		vaultSetting: settings,
// 		vault:        v,
// 	}
// }

// func (v *vault) TryOpen(string profile, password string) error {
// 	return profi
// }

// // Open profile
// func (v *vault) Open(profileName string) *vaultProfile {
// 	_, secret := Environment.GetCredentials(v.vaultSetting.GetIdentifier())
// 	return v.Open(profileName, secret)
// }

// // Open profile
// func (v *vault) Open(profileName string, password string) *vaultProfile {
// 	profile, err := v.vault.Profile(profileName)
// 	if err != nil {
// 		Panic(err, "Looks like the profile does not exist")
// 	}
// 	err = profile.Unlock(password)
// 	if err != nil {
// 		Panic(err, "Looks like the password is wrong")
// 	}
// 	return &vaultProfile{
// 		vaultProfile: profile,
// 	}
// }

// func (v *vault) GetProfileNames() []string {
// 	names, err := v.vault.ProfileNames()
// 	if err != nil {
// 		Panic(err, "Unable to obtain profiles from vault")
// 	}
// 	return names
// }

// func (p *vaultProfile) GetItemCount() int {
// 	items, err := p.vaultProfile.Items()
// 	if err != nil {
// 		Panic(err, "Unable to obtain Items in profile")
// 	}
// 	return len(items)
// }

// func (v *vault) SearchProfile(userRegex string) []string {
// 	regex, _ := regexp.Compile(userRegex)
// 	result := make([]string, 0)
// 	profiles, _ := v.vault.ProfileNames()
// 	for _, profile := range profiles {
// 		if regex.MatchString(profile) {
// 			result = append(result, profile)
// 		}
// 	}
// 	return result
// }

// func (p *vaultProfile) GetPassword(itemName string) string {
// 	items, err := p.vaultProfile.Items()
// 	if err != nil {
// 		Panic(err, "Unable to obtain Items in profile")
// 	}
// 	for _, item := range items {
// 		if strings.EqualFold(item.Title(), itemName) {
// 			details, _ := item.Detail()
// 			for _, f := range details.Fields() {
// 				log.Println(f.Type())
// 				log.Println(f.Value())
// 			}
// 		}
// 	}
// 	return ""
// }
