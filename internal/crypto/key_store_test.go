package vaultcrypto

import (
	"reflect"
	"testing"
)

func Test_keyStore_Lock(t *testing.T) {
	ks := NewKeyStore()

	ks.Unlock("test", []byte("test"))

	type args struct {
		login string
	}
	tests := []struct {
		name    string
		v       *keyStore
		args    args
		wantErr bool
	}{
		{
			name: "valid login",
			v:    ks,
			args: args{
				login: "test",
			},
			wantErr: false,
		},
		{
			name: "login no found",
			v:    ks,
			args: args{
				login: "test1",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.v.Lock(tt.args.login); (err != nil) != tt.wantErr {
				t.Errorf("keyStore.Lock() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_keyStore_Unlock(t *testing.T) {
	ks := NewKeyStore()

	type args struct {
		login string
		key   []byte
	}
	tests := []struct {
		name    string
		v       *keyStore
		args    args
		wantErr bool
	}{
		{
			name: "unlock",
			v:    ks,
			args: args{
				login: "test",
				key:   []byte("test"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.v.Unlock(tt.args.login, tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("keyStore.Unlock() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_keyStore_GetKey(t *testing.T) {
	ks := NewKeyStore()

	login := "test"
	key := []byte("test")

	type args struct {
		login string
		key   []byte
	}
	tests := []struct {
		name    string
		v       *keyStore
		args    args
		op      func(string, []byte) error
		want    []byte
		wantErr bool
	}{
		{
			name: "valid login",
			v:    ks,
			args: args{
				login: login,
				key:   key,
			},
			op:      ks.Unlock,
			want:    key,
			wantErr: false,
		},
		{
			name: "not found login",
			v:    ks,
			args: args{
				login: "login",
				key:   key,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.op != nil {
				tt.op(tt.args.login, tt.args.key)
			}
			got, err := tt.v.GetKey(tt.args.login)
			if (err != nil) != tt.wantErr {
				t.Errorf("keyStore.GetKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("keyStore.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
