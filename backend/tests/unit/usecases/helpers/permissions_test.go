package helpers

import (
	"band-manager-backend/internal/usecases/helpers"
	"testing"
)

func TestIsManagerOrModeratorRole(t *testing.T) {
	tests := []struct {
		name     string
		role     string
		expected bool
	}{
		{
			name:     "should return true for manager role",
			role:     helpers.RoleManager,
			expected: true,
		},
		{
			name:     "should return true for moderator role",
			role:     helpers.RoleModerator,
			expected: true,
		},
		{
			name:     "should return false for member role",
			role:     helpers.RoleMember,
			expected: false,
		},
		{
			name:     "should return false for empty role",
			role:     "",
			expected: false,
		},
		{
			name:     "should return false for invalid role",
			role:     "invalid_role",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := helpers.IsManagerOrModeratorRole(tt.role)
			if result != tt.expected {
				t.Errorf("IsManagerOrModeratorRole(%s) = %v, want %v",
					tt.role, result, tt.expected)
			}
		})
	}
}
