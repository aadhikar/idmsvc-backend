package model

import "gorm.io/gorm"

// See: https://gorm.io/docs/models.html

type Ipa struct {
	gorm.Model
	CaCerts []IpaCert `gorm:"foreignKey:id"`
	// TODO Do we want to create an ipa_server_list table
	//      related with this Ipa entry?
	// NOTE Thinking about this as a comma separated list
	//      of servers
	Servers      []IpaServer `gorm:"foreignKey:id"`
	RealmName    *string
	RealmDomains string

	Domain Domain `gorm:"foreignKey:id"`
}

func (i *Ipa) AfterCreate(tx *gorm.DB) (err error) {
	if i == nil {
		return nil
	}
	if i.CaCerts != nil {
		for _, cert := range i.CaCerts {
			cert.IpaID = i.ID
		}
	}
	if i.Servers != nil {
		for _, server := range i.Servers {
			server.IpaID = i.ID
		}
	}
	return nil
}
