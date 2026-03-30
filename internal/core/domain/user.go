package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/Mihail4531/golang-todo/internal/core/errors"
)

type User struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func NewUser(id int, version int, fullName string, phoneNumber *string) *User {
	return &User{
		ID:          id,
		Version:     version,
		FullName:    fullName,
		PhoneNumber: phoneNumber,
	}
}
func NewUserUninitialized(fullName string, phoneNumber *string) *User {
	return NewUser(
		UninitializedId,
		UninitializedVersion,
		fullName,
		phoneNumber,
	)
}
func (u *User) Validate() error {
	FullNameLen := len([]rune(u.FullName))
	if FullNameLen < 3 || FullNameLen > 100 {
		return fmt.Errorf("invalid 'FullName' len: %d: %w", FullNameLen, core_errors.ErrInvalidArgument)
	}
	if u.PhoneNumber != nil {
		phoneNumberLen := len([]rune(*u.PhoneNumber))
		if phoneNumberLen < 10 || phoneNumberLen > 15 {
			return fmt.Errorf("invalid 'PhoneNumber' len: %d: %w", phoneNumberLen, core_errors.ErrInvalidArgument)
		}
		re := regexp.MustCompile(`^\+[0-9]{10,15}$`)
		if !re.MatchString(*u.PhoneNumber) {
			return fmt.Errorf("invalid 'PhoneNumber' format: %d: %w", phoneNumberLen, core_errors.ErrInvalidArgument)
		}
	}
	return nil
}

type UserPatch struct {
	FullName    Nullable[string]
	PhoneNumber Nullable[string]
}
func NewUserPatch(fullName Nullable[string], phoneNumber Nullable[string]) *UserPatch{
	return &UserPatch{
		FullName: fullName,
		PhoneNumber: phoneNumber,
	}
}
func (p *UserPatch) Validate() error {
	if p.FullName.Set && p.FullName.Value == nil {
		return fmt.Errorf("'FullName' cant be patched to null : %w", core_errors.ErrInvalidArgument)
	}
	return nil
}
func (u *User) ApplayPatch(patch *UserPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate user path %w", err)
	}
	tmp := *u
	if patch.FullName.Set {
		tmp.FullName = *patch.FullName.Value
	}
	if patch.PhoneNumber.Set {
		tmp.PhoneNumber = patch.PhoneNumber.Value
	}
	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched user: %w", err)
	}
	*u = tmp
	return nil
}
