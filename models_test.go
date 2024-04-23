package prac

import "testing"

func TestNewCountryCode(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name    string
		args    args
		want    CountryCode
		wantErr bool
	}{
		{name: "AIA should work.", args: args{"AIA"}, want: CountryCode("AIA")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCountryCode(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCountryCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewCountryCode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
