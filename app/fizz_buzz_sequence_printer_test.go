package app_test

import (
	"inoxe_e/OOP_fizzbuzz/app"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestFizzBuzzSequencePrinterNone(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	nc_mock := app.NewMockNumberConverter(ctrl)
	nc_mock.EXPECT().Convert(1).MaxTimes(0)

	op_mock := app.NewMockOutputWriter(ctrl)
	op_mock.EXPECT().Write("1").MaxTimes(0)

	p := app.FizzBuzzSequencePrinter{FizzBuzz: nc_mock, IOWriter: op_mock}
	p.PrintRange(0, -1)
}

func TestFizzBuzzSequencePrinter1to3(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	nc_mock := app.NewMockNumberConverter(ctrl)
	nc_mock.EXPECT().Convert(1).Return("1")
	nc_mock.EXPECT().Convert(2).Return("2")
	nc_mock.EXPECT().Convert(3).Return("Fizz")

	op_mock := app.NewMockOutputWriter(ctrl)
	op_mock.EXPECT().Write("1 1\n")
	op_mock.EXPECT().Write("2 2\n")
	op_mock.EXPECT().Write("3 Fizz\n")

	p := app.FizzBuzzSequencePrinter{FizzBuzz: nc_mock, IOWriter: op_mock}
	p.PrintRange(1, 3)
}
