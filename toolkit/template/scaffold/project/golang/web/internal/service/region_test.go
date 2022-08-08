package service

import "testing"

func TestGetForwardGeocodingInformation(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name         string
		args         args
		wantLocation string
		wantErr      bool
	}{
		{
			name:         "test1",
			args:         args{address: "北京市海淀区燕园街道北京大学"},
			wantLocation: "116.310003,39.991957",
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLocation, err := GetForwardGeocodingInformation(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetForwardGeocodingInformation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotLocation != tt.wantLocation {
				t.Errorf("GetForwardGeocodingInformation() gotLocation = %v, want %v", gotLocation, tt.wantLocation)
			}
		})
	}
}

func TestGetReverseGeocodingInformation(t *testing.T) {
	type args struct {
		location string
	}
	tests := []struct {
		name        string
		args        args
		wantAddress string
		wantErr     bool
	}{
		{
			name:        "test1",
			args:        args{location: "116.310003,39.991957"},
			wantAddress: "北京市海淀区燕园街道北京大学",
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAddress, err := GetReverseGeocodingInformation(tt.args.location)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetReverseGeocodingInformation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAddress != tt.wantAddress {
				t.Errorf("GetReverseGeocodingInformation() gotAddress = %v, want %v", gotAddress, tt.wantAddress)
			}
		})
	}
}
