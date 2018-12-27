package iex

import (
	"net/http"
	"reflect"
	"testing"
)

func TestIEX_GetDEEP(t *testing.T) {
	type fields struct {
		client *http.Client
	}
	type args struct {
		symbol string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *DEEP
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &IEX{
				client: tt.fields.client,
			}
			got, err := i.GetDEEP(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("IEX.GetDEEP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IEX.GetDEEP() = %v, want %v", got, tt.want)
			}
		})
	}
}
