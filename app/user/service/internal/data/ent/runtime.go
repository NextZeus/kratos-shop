// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/nextzeus/kratos-shop/app/user/service/internal/data/ent/address"
	"github.com/nextzeus/kratos-shop/app/user/service/internal/data/ent/card"
	"github.com/nextzeus/kratos-shop/app/user/service/internal/data/ent/schema"
	"github.com/nextzeus/kratos-shop/app/user/service/internal/data/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	addressFields := schema.Address{}.Fields()
	_ = addressFields
	// addressDescCreatedAt is the schema descriptor for created_at field.
	addressDescCreatedAt := addressFields[5].Descriptor()
	// address.DefaultCreatedAt holds the default value on creation for the created_at field.
	address.DefaultCreatedAt = addressDescCreatedAt.Default.(func() time.Time)
	// addressDescUpdatedAt is the schema descriptor for updated_at field.
	addressDescUpdatedAt := addressFields[6].Descriptor()
	// address.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	address.DefaultUpdatedAt = addressDescUpdatedAt.Default.(func() time.Time)
	cardFields := schema.Card{}.Fields()
	_ = cardFields
	// cardDescCreatedAt is the schema descriptor for created_at field.
	cardDescCreatedAt := cardFields[5].Descriptor()
	// card.DefaultCreatedAt holds the default value on creation for the created_at field.
	card.DefaultCreatedAt = cardDescCreatedAt.Default.(func() time.Time)
	// cardDescUpdatedAt is the schema descriptor for updated_at field.
	cardDescUpdatedAt := cardFields[6].Descriptor()
	// card.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	card.DefaultUpdatedAt = cardDescUpdatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[3].Descriptor()
	// server.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[4].Descriptor()
	// server.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
}
