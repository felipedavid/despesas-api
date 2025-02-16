package service

import "testing"

func TestRegisterUserParams_Valid(t *testing.T) {
	tests := []struct {
		name     string
		params   RegisterUserParams
		expected bool
		errors   map[string]string
	}{
		{
			name: "Valid user",
			params: RegisterUserParams{
				Name:     strPtr("John"),
				Email:    strPtr("john@example.com"),
				Password: strPtr("password123"),
			},
			expected: true,
			errors:   nil,
		},
		{
			name: "Missing name",
			params: RegisterUserParams{
				Email:    strPtr("john@example.com"),
				Password: strPtr("password123"),
			},
			expected: false,
			errors: map[string]string{
				"name": "must be provided",
			},
		},
		{
			name: "Name too short",
			params: RegisterUserParams{
				Name:     strPtr("Jo"),
				Email:    strPtr("john@example.com"),
				Password: strPtr("password123"),
			},
			expected: false,
			errors: map[string]string{
				"name": "should be at least 3 characters long",
			},
		},
		{
			name: "Missing password",
			params: RegisterUserParams{
				Name:  strPtr("John"),
				Email: strPtr("john@example.com"),
			},
			expected: false,
			errors: map[string]string{
				"password": "must be provided",
			},
		},
		{
			name: "Password too short",
			params: RegisterUserParams{
				Name:     strPtr("John"),
				Email:    strPtr("john@example.com"),
				Password: strPtr("pass"),
			},
			expected: false,
			errors: map[string]string{
				"password": "should be at least 8 characters long",
			},
		},
		{
			name: "Missing email",
			params: RegisterUserParams{
				Name:     strPtr("John"),
				Password: strPtr("password123"),
			},
			expected: false,
			errors: map[string]string{
				"email": "must be provided",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid := tt.params.Valid()
			if valid != tt.expected {
				t.Errorf("expected valid to be %v, got %v", tt.expected, valid)
			}
			if len(tt.params.Errors) != len(tt.errors) {
				t.Errorf("expected %d errors, got %d", len(tt.errors), len(tt.params.Errors))
			}
			for field, message := range tt.errors {
				if tt.params.Errors[field] != message {
					t.Errorf("expected error for %s to be '%s', got '%s'", field, message, tt.params.Errors[field])
				}
			}
		})
	}
}

func strPtr(s string) *string {
	return &s
}
