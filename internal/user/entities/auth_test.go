package entities

import "testing"

func TestAuthClaims_AuthorizeUser(t *testing.T) {
	key := []byte("key")
	ac := &AuthClaims{Login: "test"}
	token, _ := ac.AuthorizeUser(key)

	type args struct {
		jwtKey []byte
	}
	tests := []struct {
		name    string
		ac      *AuthClaims
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "correct jwtKey",
			ac: &AuthClaims{
				Login: "test",
			},
			args: args{
				jwtKey: key,
			},
			want:    token,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ac.AuthorizeUser(tt.args.jwtKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthClaims.AuthorizeUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AuthClaims.AuthorizeUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
