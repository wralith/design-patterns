package builder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuilder(t *testing.T) {
	tests := []struct {
		want string
		got  string
	}{
		{
			want: "SELECT * FROM users",
			got:  NewBuilder().Select("*").From("users").Build(),
		},
		{
			want: "SELECT * FROM users WHERE id = 1",
			got:  NewBuilder().Select("*").From("users").Where("id = 1").Build(),
		},
		{
			want: "SELECT username, email FROM users WHERE id = 1",
			got:  NewBuilder().Select("username", "email").From("users").Where("id = 1").Build(),
		},
	}
	for _, tc := range tests {
		assert.Equal(t, tc.want, tc.got)
	}
}
