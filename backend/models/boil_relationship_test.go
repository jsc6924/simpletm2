// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("PermissionToGameUsingGame", testPermissionToOneGameUsingGame)
	t.Run("PermissionToUserUsingUser", testPermissionToOneUserUsingUser)
	t.Run("TranslateToGameUsingGame", testTranslateToOneGameUsingGame)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("GameToGamePermissions", testGameToManyGamePermissions)
	t.Run("GameToGameTranslates", testGameToManyGameTranslates)
	t.Run("UserToUserPermissions", testUserToManyUserPermissions)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("PermissionToGameUsingGamePermissions", testPermissionToOneSetOpGameUsingGame)
	t.Run("PermissionToUserUsingUserPermissions", testPermissionToOneSetOpUserUsingUser)
	t.Run("TranslateToGameUsingGameTranslates", testTranslateToOneSetOpGameUsingGame)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("GameToGamePermissions", testGameToManyAddOpGamePermissions)
	t.Run("GameToGameTranslates", testGameToManyAddOpGameTranslates)
	t.Run("UserToUserPermissions", testUserToManyAddOpUserPermissions)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}