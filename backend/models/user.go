package models

import (
	"github.com/go-webauthn/webauthn/webauthn"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID    `bson:"_id,omitempty"`
	Username    string                `bson:"username"`
	Password    string                `bson:"password"`
	DisplayName string                `bson:"display_name"`
	Credentials []webauthn.Credential `bson:"credentials"`
}

func (u *User) WebAuthnID() []byte {
	return []byte(u.ID.Hex())
}

func (u *User) WebAuthnName() string {
	return u.Username
}

func (u *User) WebAuthnDisplayName() string {
	return u.DisplayName
}

func (u *User) WebAuthnCredentials() []webauthn.Credential {
	return u.Credentials
}

func (u *User) WebAuthnIcon() string {
	return ""
}
