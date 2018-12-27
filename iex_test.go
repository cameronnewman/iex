package iex

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewIEX(t *testing.T) {
	type args struct {
		client *http.Client
	}
	tests := []struct {
		name string
		args args
		want *IEX
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIEX(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIEX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIEX_getJSON(t *testing.T) {
	type fields struct {
		client *http.Client
	}
	type args struct {
		url    string
		result interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &IEX{
				client: tt.fields.client,
			}
			if err := i.getJSON(tt.args.url, tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("IEX.getJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIEX_endpoint(t *testing.T) {
	type fields struct {
		client *http.Client
	}
	type args struct {
		route string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &IEX{
				client: tt.fields.client,
			}
			if got := i.endpoint(tt.args.route); got != tt.want {
				t.Errorf("IEX.endpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
