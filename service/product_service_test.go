package service

import (
	"chap3-challenge2/model"
	"chap3-challenge2/repository/mocks"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestProductService_GetAll(t *testing.T) {
	productRepository := mocks.NewIProductRepository(t)
	type args struct {
		userId string
		role   string
	}
	tests := []struct {
		name     string
		ps       *ProductService
		args     args
		want     []model.ProductResponse
		mockFunc func()
		wantErr  bool
	}{
		// TODO: Add test cases.
		// 1. Mock test service get all product found
		// 2. Mock test service get all product not found
		{
			name: "Case #1 - Get all product found",
			ps: &ProductService{
				ProductRepository: productRepository,
			},
			args: args{
				userId: "1",
				role:   "admin",
			},
			want: []model.ProductResponse{
				{
					ID:     "1",
					UserID: "1",
					Name:   "Banana",
					Price:  10000,
				},
				{
					ID:     "2",
					UserID: "1",
					Name:   "Apple",
					Price:  5000,
				},
			},
			mockFunc: func() {
				productRepository.
					On("Get").
					Return(
						[]model.Product{
							{
								ID:     "1",
								UserID: "1",
								Name:   "Banana",
								Price:  10000,
							},
							{
								ID:     "2",
								UserID: "1",
								Name:   "Apple",
								Price:  5000,
							},
						}, nil,
					).Once()
			},
			wantErr: false,
		},
		{
			name: "Case #2 - Get all product not found",
			ps: &ProductService{
				ProductRepository: productRepository,
			},
			args: args{
				userId: "1",
				role:   "admin",
			},
			want: []model.ProductResponse{},
			mockFunc: func() {
				productRepository.
					On("Get").
					Return(
						[]model.Product{}, nil,
					).Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			got, err := tt.ps.GetAll(tt.args.userId, tt.args.role)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductService_GetById(t *testing.T) {
	productRepository := mocks.NewIProductRepository(t)
	type args struct {
		id     string
		role   string
		userId string
	}
	tests := []struct {
		name     string
		ps       *ProductService
		args     args
		want     model.ProductResponse
		mockFunc func()
		wantErr  bool
	}{
		// TODO: Add test cases.
		// 1. Mock test service get one product found
		// 2. Mock test service get one product not found
		{
			name: "Case #1 - Get one product found",
			ps: &ProductService{
				ProductRepository: productRepository,
			},
			args: args{
				id:     "1",
				userId: "1",
				role:   "admin",
			},
			want: model.ProductResponse{
				ID:     "1",
				UserID: "1",
				Name:   "Banana",
				Price:  10000,
			},
			mockFunc: func() {
				productRepository.
					On("GetOne", mock.AnythingOfType("string")).
					Return(
						model.Product{
							ID:     "1",
							UserID: "1",
							Name:   "Banana",
							Price:  10000,
						}, nil,
					).Once()
			},
			wantErr: false,
		},
		{
			name: "Case #2 - Get one product not found",
			ps: &ProductService{
				ProductRepository: productRepository,
			},
			args: args{
				id:     "1",
				userId: "1",
				role:   "admin",
			},
			want: model.ProductResponse{},
			mockFunc: func() {
				productRepository.
					On("GetOne", mock.AnythingOfType("string")).
					Return(
						// Seharusnya berhasil, tapi errors dari mocking repo ketika di compare di service engga match. Mungkin salah di return errors disini.
						model.Product{}, errors.New("record not found"),
					).Once()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			got, err := tt.ps.GetById(tt.args.id, tt.args.role, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductService.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}
