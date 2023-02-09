package usecase

import (
	"context"
	"reflect"
	"testing"
)

type CurrencyConverterMock struct {
}

func NewCurrencyConverterMock() CurrencyConverter {
	return &CurrencyConverterMock{}
}

// mock соответствует интерфейсу
func (p *CurrencyConverterMock) GetValue() int {
	return 0
}

func TestNewHandler(t *testing.T) {
	type args struct {
		CurrencyConverter CurrencyConverter
		cache             CurrencyCache
	}
	tests := []struct {
		name string
		args args
		want Handler
	}{
		{
			args: args{
				CurrencyConverter: NewCurrencyConverterMock(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBudget(context.Background(), tt.args.CurrencyConverter, tt.args.cache); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBudget() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handler_Create(t *testing.T) {
	type fields struct {
		CurrencyConverter CurrencyConverter
		cache             CurrencyCache
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//h := handler{
			//	CurrencyConverter: tt.fields.CurrencyConverter,
			//	cache:            tt.fields.cache,
			//}
		})
	}
}

func Test_handler_List(t *testing.T) {
	type fields struct {
		CurrencyConverter CurrencyConverter
		cache             CurrencyCache
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//h := handler{
			//	CurrencyConverter: tt.fields.CurrencyConverter,
			//	cache:            tt.fields.cache,
			//}
		})
	}
}
