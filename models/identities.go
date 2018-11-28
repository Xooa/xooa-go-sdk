package models

import "time"


// Struct Identity Attribute Object
type Attrs struct {
	Name  string `json:"name"`
	Ecert bool `json:"ecert"`
	Value string `json:"value"`
}


// Struct Identity
type Identity struct {
	CanManageIdentities bool `json:"canManageIdentities"`
	CreatedAt           time.Time `json:"createdAt"`
	ApiToken            string `json:"ApiToken"`
	IdentityName        string `json:"IdentityName"`
	Access              string `json:"Access"`
	Id                  string `json:"Id"`
	Attrs               []Attrs `json:"Attrs"`
}


//Struct Delete Identity Response
type DeleteIdentityResponse struct {
	Deleted bool `json:"deleted"` 
}


//Struct Regenerate Identity Response
type RegenerateIdentityResponse struct {
	Success bool `json:"success"`
}

