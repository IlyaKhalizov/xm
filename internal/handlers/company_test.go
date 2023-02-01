package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"xm/internal/handlers"
	"xm/internal/handlers/dto"
	"xm/internal/repositories/entities"

	"github.com/go-chi/chi/v5"
	gomock "github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestGetCompany(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cases := map[string]struct {
		mocks     func(*MockcompanyRepository)
		WantCode  int
		CompanyID string
	}{
		"success": {
			mocks: func(mr *MockcompanyRepository) {
				mr.EXPECT().GetCompany(ctx, 1).Return(entities.Company{
					Id:              1,
					Name:            "Test Company",
					Description:     "Test",
					EmployeesAmount: 10,
					Registered:      true,
					Type:            "corporations",
				}, nil)
			},
			WantCode:  200,
			CompanyID: "1",
		},
	}

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			var companyBody dto.Company
			var logger zerolog.Logger

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			companyRepository := NewMockcompanyRepository(ctrl)

			if tc.mocks != nil {
				tc.mocks(companyRepository)
			}

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/v1/company/{id}", nil).WithContext(ctx)

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tc.CompanyID)

			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

			handlers := handlers.NewCompanyHandler(ctx, companyRepository, logger)
			handler := http.HandlerFunc(handlers.GetCompany)
			handler.ServeHTTP(w, r)

			res := w.Result()
			err := json.NewDecoder(res.Body).Decode(&companyBody)
			if err != nil {
				t.Error("error decode response body")
			}

			assert.Equal(t, tc.WantCode, w.Code)
			assert.Equal(t, companyBody.Name, "Test Company")
		})
	}
}

func TestDeleteCompany(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cases := map[string]struct {
		mocks     func(*MockcompanyRepository)
		WantCode  int
		CompanyID string
	}{
		"success": {
			mocks: func(mr *MockcompanyRepository) {
				mr.EXPECT().DeleteCompany(ctx, 1).Return(nil)
			},
			WantCode:  200,
			CompanyID: "1",
		},
	}

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			var logger zerolog.Logger

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			companyRepository := NewMockcompanyRepository(ctrl)

			if tc.mocks != nil {
				tc.mocks(companyRepository)
			}

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodDelete, "/v1/company/{id}", nil).WithContext(ctx)

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tc.CompanyID)

			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

			handlers := handlers.NewCompanyHandler(ctx, companyRepository, logger)
			handler := http.HandlerFunc(handlers.DeleteCompany)
			handler.ServeHTTP(w, r)

			assert.Equal(t, tc.WantCode, w.Code)
		})
	}
}

func TestUpdateCompany(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	dtoCompany := dto.Company{
		Name:            "Test Company2",
		Description:     "Test",
		EmployeesAmount: 10,
		Registered:      true,
		Type:            "corporations",
	}

	cases := map[string]struct {
		mocks     func(*MockcompanyRepository)
		WantCode  int
		CompanyID string
	}{
		"success": {
			mocks: func(mr *MockcompanyRepository) {
				mr.EXPECT().UpdateCompany(ctx, 1, dtoCompany).Return(entities.Company{
					Id:              1,
					Name:            "Test Company2",
					Description:     "Test",
					EmployeesAmount: 10,
					Registered:      true,
					Type:            "corporations",
				}, nil)
			},
			WantCode:  200,
			CompanyID: "1",
		},
	}

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			var companyBody dto.Company
			var logger zerolog.Logger

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			companyRepository := NewMockcompanyRepository(ctrl)

			if tc.mocks != nil {
				tc.mocks(companyRepository)
			}

			bodyReq := map[string]interface{}{
				"name":             "Test Company2",
				"description":      "Test",
				"employees_amount": 10,
				"registered":       true,
				"type":             "corporations",
			}
			bodySend, err := json.Marshal(bodyReq)
			if err != nil {
				t.Error("error marshal request body")
			}

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPatch, "/v1/company/{id}", bytes.NewReader(bodySend)).WithContext(ctx)
			r.Header.Set("Content-Type", "application/json;charset=utf-8")

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tc.CompanyID)

			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

			handlers := handlers.NewCompanyHandler(ctx, companyRepository, logger)
			handler := http.HandlerFunc(handlers.UpdateCompany)
			handler.ServeHTTP(w, r)

			res := w.Result()
			err = json.NewDecoder(res.Body).Decode(&companyBody)
			if err != nil {
				t.Error("error decode response body")
			}

			assert.Equal(t, tc.WantCode, w.Code)
			assert.Equal(t, companyBody.Name, "Test Company2")
		})
	}
}

func TestCreateCompany(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	dtoCompany := dto.Company{
		Name:            "Test Company3",
		Description:     "Test",
		EmployeesAmount: 10,
		Registered:      true,
		Type:            "corporations",
	}

	cases := map[string]struct {
		mocks    func(*MockcompanyRepository)
		WantCode int
	}{
		"success": {
			mocks: func(mr *MockcompanyRepository) {
				mr.EXPECT().CreateCompany(ctx, dtoCompany).Return(entities.Company{
					Id:              1,
					Name:            "Test Company3",
					Description:     "Test",
					EmployeesAmount: 10,
					Registered:      true,
					Type:            "corporations",
				}, nil)
			},
			WantCode: 201,
		},
	}

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			var companyBody dto.Company
			var logger zerolog.Logger

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			companyRepository := NewMockcompanyRepository(ctrl)

			if tc.mocks != nil {
				tc.mocks(companyRepository)
			}

			bodyReq := map[string]interface{}{
				"name":             "Test Company3",
				"description":      "Test",
				"employees_amount": 10,
				"registered":       true,
				"type":             "corporations",
			}
			bodySend, err := json.Marshal(bodyReq)
			if err != nil {
				t.Error("error marshal request body")
			}

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/v1/company", bytes.NewReader(bodySend)).WithContext(ctx)
			r.Header.Set("Content-Type", "application/json;charset=utf-8")

			rctx := chi.NewRouteContext()

			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

			handlers := handlers.NewCompanyHandler(ctx, companyRepository, logger)
			handler := http.HandlerFunc(handlers.CreateCompany)
			handler.ServeHTTP(w, r)

			res := w.Result()
			err = json.NewDecoder(res.Body).Decode(&companyBody)
			if err != nil {
				t.Error("error decode response body")
			}

			assert.Equal(t, tc.WantCode, w.Code)
			assert.Equal(t, companyBody.Name, "Test Company3")
		})
	}
}
