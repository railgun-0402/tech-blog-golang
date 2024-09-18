package handler

import (
	"github.com/labstack/echo/v4"
	"reflect"
	"testing"
)

func Test_htmlBlob(t *testing.T) {
	type args struct {
		file string
		data map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := htmlBlob(tt.args.file, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("htmlBlob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("htmlBlob() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_render(t *testing.T) {
	type args struct {
		c    echo.Context
		file string
		data map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := render(tt.args.c, tt.args.file, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("render() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
