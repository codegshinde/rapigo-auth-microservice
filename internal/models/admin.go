package models

import (
	"time"
)

// Address represents the Address schema
type Address struct {
	District    string `json:"district" bson:"district"`
	SubDistrict string `json:"subDistrict" bson:"subDistrict"`
	Village     string `json:"village,omitempty" bson:"village,omitempty"`
	PostalCode  string `json:"postalCode" bson:"postalCode"`
}

// Admin represents the Admin schema
type AdminInput struct {
	FirstName     string    `json:"firstName" bson:"firstName"`
	LastName      string    `json:"lastName" bson:"lastName"`
	Access        []string  `json:"access" bson:"access"`
	AdminRole     string    `json:"adminRole" bson:"adminRole"`
	AdminPosition string    `json:"adminPosition" bson:"adminPosition"`
	Email         string    `json:"email" bson:"email"`
	Mobile        string    `json:"mobile" bson:"mobile"`
	AdminId       string    `json:"adminId" bson:"adminId"`
	Password      string    `json:"password" bson:"password"`
	Address       []Address `json:"address" bson:"address"`
	CreatedAt     time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}

type AdminResponse struct {
	Token         string    `json:"token,omitempty"`
	FirstName     string    `json:"firstName" bson:"firstName"`
	LastName      string    `json:"lastName" bson:"lastName"`
	Access        []string  `json:"access" bson:"access"`
	AdminRole     string    `json:"adminRole" bson:"adminRole"`
	AdminPosition string    `json:"adminPosition" bson:"adminPosition"`
	Email         string    `json:"email" bson:"email"`
	Mobile        string    `json:"mobile" bson:"mobile"`
	AdminId       string    `json:"adminId" bson:"adminId"`
	Password      string    `json:"-"`
	Address       []Address `json:"address" bson:"address"`
	AvatarKey     string    `json:"avatarKey,omitempty" bson:"avatarKey,omitempty"`
	CreatedAt     time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt     time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

// NewAdminWithoutPassword creates a new Admin without the Password field.
func (a AdminResponse) ConvertToAdminResponse(token string) AdminResponse {
	return AdminResponse{
		Token:         token,
		FirstName:     a.FirstName,
		LastName:      a.LastName,
		Access:        a.Access,
		AdminRole:     a.AdminRole,
		AdminPosition: a.AdminPosition,
		Email:         a.Email,
		Mobile:        a.Mobile,
		AdminId:       a.AdminId,
		Address:       a.Address,
		AvatarKey:     a.AvatarKey,
		CreatedAt:     a.CreatedAt,
		UpdatedAt:     a.UpdatedAt,
	}
}
