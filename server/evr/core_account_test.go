// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    gameProfiles, err := UnmarshalGameProfiles(bytes)
//    bytes, err = gameProfiles.Marshal()

package evr

import (
	"testing"

	"github.com/gofrs/uuid"
)

func TestGUID_UnmarshalBytes(t *testing.T) {
	tests := []struct {
		name    string
		g       *GUID
		b       []byte
		wantErr bool
	}{
		{
			name:    "valid bytes",
			g:       &GUID{},
			b:       []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			wantErr: false,
		},
		{
			name:    "invalid bytes",
			g:       &GUID{},
			b:       []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			wantErr: true,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.g.UnmarshalBytes(tt.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("GUID.UnmarshalBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// Add additional assertions if needed
		})
	}
}

func TestGUID_UnmarshalJSON(t *testing.T) {

	// Test unmarshalling

	tests := []struct {
		name string
		data string
		want GUID
	}{
		{
			name: "valid GUID",
			data: `"01020304-0506-0708-090A-0A0B0C0D0E0F"`,
			want: GUID(uuid.FromStringOrNil("01020304-0506-0708-090A-0A0B0C0D0E0F")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got GUID
			if err := got.UnmarshalJSON([]byte(tt.data)); err != nil {
				t.Errorf("GUID.JSON() error = %v", err)
			}
			if got != tt.want {
				t.Errorf("GUID.JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGUID_MarshalJSON(t *testing.T) {
	// Test marshalling

	tests := []struct {
		name string
		g    GUID
		want string
	}{
		{
			name: "valid GUID",
			g:    GUID(uuid.FromStringOrNil("01020304-0506-0708-090a-0A0b0C0D0E0F")),
			want: `"01020304-0506-0708-090A-0A0B0C0D0E0F"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := tt.g.MarshalJSON(); string(got) != tt.want {
				t.Errorf("GUID.JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
