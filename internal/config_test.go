package internal

import (
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/stretchr/testify/assert"
)

func TestNewAppConfig(t *testing.T) {
	type nestedCfg struct {
		Name string `env:"TEST_CFG_NAME"`
	}
	// NOTE: zero need and counterproductive to allow pointer nested structs for config
	type cfg struct {
		NestedCfg       nestedCfg
		Length          int     `env:"TEST_CFG_LEN"`
		OptionalLength  *int    `env:"TEST_CFG_OPT_LEN"`
		OptionalString  *string `env:"TEST_CFG_STRING_PTR"`
		OptionalBool    *bool   `env:"TEST_CFG_BOOL_PTR"`
		BoolWithDefault bool    `env:"TEST_CFG_BOOL_DEFAULT,false"`
	}

	type params struct {
		name        string
		want        *cfg
		errContains string
		environ     map[string]string
	}

	tests := []params{
		{
			name:    "correct env",
			want:    &cfg{NestedCfg: nestedCfg{Name: "name"}, Length: 10, OptionalLength: pointers.New(40), BoolWithDefault: true},
			environ: map[string]string{"TEST_CFG_NAME": "name", "TEST_CFG_LEN": "10", "TEST_CFG_OPT_LEN": "40", "TEST_CFG_BOOL_DEFAULT": "true"},
		},
		{
			name:    "correct env with missing pointer fields",
			want:    &cfg{NestedCfg: nestedCfg{Name: "name"}, Length: 10, BoolWithDefault: false},
			environ: map[string]string{"TEST_CFG_NAME": "name", "TEST_CFG_LEN": "10"},
		},
		{
			name:    "empty but set envvar string pointer field is not nil",
			want:    &cfg{NestedCfg: nestedCfg{Name: "name"}, Length: 10, OptionalString: pointers.New("")},
			environ: map[string]string{"TEST_CFG_NAME": "name", "TEST_CFG_LEN": "10", "TEST_CFG_STRING_PTR": ""},
		},
		{
			name:    "unset envvar string pointer field is nil",
			want:    &cfg{NestedCfg: nestedCfg{Name: "name"}, Length: 10},
			environ: map[string]string{"TEST_CFG_NAME": "name", "TEST_CFG_LEN": "10"},
		},
		{
			name:        "bad env conversion",
			environ:     map[string]string{"TEST_CFG_NAME": "name", "TEST_CFG_LEN": "aaa"},
			errContains: `could not set "TEST_CFG_LEN" to "Length": could not convert TEST_CFG_LEN to int`,
		},
		{
			name:        "non pointer field corresponding envvar not set",
			environ:     map[string]string{"TEST_CFG_LEN": "10"},
			errContains: `could not set "TEST_CFG_NAME" to "Name": TEST_CFG_NAME is not set but required`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.environ {
				t.Setenv(k, v)
			}

			c := &cfg{}
			err := loadEnvToConfig(c)
			if err != nil && tt.errContains == "" {
				t.Errorf("unexpected error: %v", err)

				return
			}
			if tt.errContains != "" {
				if err == nil {
					t.Errorf("expected error but got nothing")

					return
				}
				assert.Contains(t, err.Error(), tt.errContains)

				return
			}

			assert.Equal(t, tt.want, c)
		})
	}
}
