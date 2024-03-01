package entities

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func TestUser_SetPassword(t *testing.T) {
	type args struct {
		password []byte
	}
	tests := []struct {
		name    string
		u       *User
		args    args
		wantErr bool
	}{
		{
			name: "returns no err",
			u:    getTestUser(),
			args: args{
				password: []byte("super secret password"),
			},
			wantErr: false,
		},
		{
			name: "returns bcrypt to long password err",
			u:    getTestUser(),
			args: args{
				password: []byte("super secret passwordsuper secret passwordsuper secret passwordsuper secret passwordsuper secret password"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.SetPassword(tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("User.SetPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_Authenticate(t *testing.T) {
	type args struct {
		password []byte
	}
	tests := []struct {
		name    string
		u       User
		args    args
		wantErr bool
	}{
		{
			name: "same password",
			u:    *getTestUserWithPassword([]byte("password")),
			args: args{
				password: []byte("password"),
			},
			wantErr: false,
		},
		{
			name: "diff password",
			u:    *getTestUserWithPassword([]byte("password")),
			args: args{
				password: []byte("password1"),
			},
			wantErr: true,
		},
		{
			name: "empty password",
			u:    *getTestUser(),
			args: args{
				password: []byte("password"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.Authenticate(tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("User.Authenticate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetLoginFromContext(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "ctx without claims",
			args: args{
				ctx: context.Background(),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "ctx with claims",
			args: args{
				ctx: context.WithValue(context.Background(), ClaimsKeyName, &AuthClaims{
					Login:            "test",
					RegisteredClaims: jwt.RegisteredClaims{},
				}),
			},
			want:    "test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetLoginFromContext(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLoginFromContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetLoginFromContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_SetVaultKey(t *testing.T) {
	type args struct {
		vaultPassword []byte
	}
	tests := []struct {
		name    string
		u       *User
		args    args
		wantErr bool
	}{
		{
			name: "correct password",
			u:    getTestUser(),
			args: args{
				vaultPassword: []byte("password"),
			},
			wantErr: false,
		},
		{
			name: "password > 256 bit",
			u:    getTestUser(),
			args: args{
				vaultPassword: make([]byte, 257),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.SetVaultKey(tt.args.vaultPassword); (err != nil) != tt.wantErr {
				t.Errorf("User.SetVaultKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_GetVaultKey(t *testing.T) {
	p := []byte("password")
	u := getTestUserWithPassword(p)

	key, _ := u.GetVaultKey(p)

	type args struct {
		vaultPassword []byte
	}
	tests := []struct {
		name    string
		u       *User
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "correct password",
			u:    u,
			args: args{
				vaultPassword: p,
			},
			want:    key,
			wantErr: false,
		},
		{
			name: "wrong password",
			u:    u,
			args: args{
				vaultPassword: []byte("wrong password"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.GetVaultKey(tt.args.vaultPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.GetVaultKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.GetVaultKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getTestUser() *User {
	return &User{
		Login: "test",
	}
}

func getTestUserWithPassword(p []byte) *User {
	u := &User{
		Login: "test",
	}
	u.SetPassword(p)
	u.SetVaultKey(p)
	return u
}
