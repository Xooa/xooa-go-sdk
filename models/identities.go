/**
 * Go SDK for Xooa
 *
 * Copyright 2018 Xooa
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at:
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License
 * for the specific language governing permissions and limitations under the License.
 *
 * Author: Rahul Kamboj
 */

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

