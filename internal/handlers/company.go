package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"xm/internal/handlers/dto"
	"xm/internal/repositories/entities"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

type companyRepository interface {
	GetCompany(ctx context.Context, id int) (entities.Company, error)
	CreateCompany(ctx context.Context, company dto.Company) (entities.Company, error)
	DeleteCompany(ctx context.Context, id int) error
	UpdateCompany(ctx context.Context, id int, company dto.Company) (entities.Company, error)
}

type companyHandler struct {
	companyRepository companyRepository
	logger            zerolog.Logger
	ctx               context.Context
}

func NewCompanyHandler(ctx context.Context, companyRepository companyRepository, logger zerolog.Logger) companyHandler {
	return companyHandler{
		companyRepository: companyRepository,
		logger:            logger,
		ctx:               ctx,
	}
}

func (h companyHandler) GetCompany(w http.ResponseWriter, r *http.Request) {
	var err error
	var id int

	id, err = strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.errCompanyId(w, err)
	}

	company, err := h.companyRepository.GetCompany(h.ctx, id)

	if err != nil {
		h.logger.Info().Timestamp().Msg(err.Error())
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(errors.New("error was occurred while fetching the company").Error()))
		if err != nil {
			h.logger.Info().Timestamp().Msg(err.Error())
		}
		return
	}

	if company.Id == 0 {
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.WriteHeader(http.StatusNoContent)
		_, err := w.Write([]byte{})
		if err != nil {
			h.logger.Info().Timestamp().Msg(err.Error())
		}
		return
	}

	response, err := json.Marshal(company)
	if err != nil {
		h.logger.Info().Timestamp().Msg(err.Error())
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(errors.New("response marshal error").Error()))
		if err != nil {
			h.logger.Info().Timestamp().Msg(err.Error())
		}
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)
	if err != nil {
		h.logger.Info().Timestamp().Msg(err.Error())
	}
}

func (h companyHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	var err error
	var companyBody dto.Company
	var company entities.Company
	var response []byte

	err = json.NewDecoder(r.Body).Decode(&companyBody)
	if err != nil {
		h.logger.Info().Timestamp().Msg(err.Error())
	}

	err = dto.Validator.Struct(companyBody)
	if err != nil {
		h.errValidateCompanyStruct(w, err)
	}

	company, err = h.companyRepository.CreateCompany(h.ctx, companyBody)
	if err != nil {
		h.logger.Info().Timestamp().Msg(err.Error())
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(errors.New("response marshal error").Error()))
		if err != nil {
			h.logger.Info().Timestamp().Msg(err.Error())
		}
		return
	}

	response, err = json.Marshal(company)
	if err != nil {
		h.logger.Info().Timestamp().Msg(err.Error())
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(errors.New("response marshal error").Error()))
		if err != nil {
			h.logger.Info().Timestamp().Msg(err.Error())
		}
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(response)
	if err != nil {
		h.logger.Info().Timestamp().Msg(err.Error())
	}
}

func (h companyHandler) DeleteCompany(w http.ResponseWriter, r *http.Request) {
	var err error
	var id int

	id, err = strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.errCompanyId(w, err)
	}

	err = h.companyRepository.DeleteCompany(h.ctx, id)
	if err != nil {
		h.logger.Info().Timestamp().Msg(err.Error())
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(errors.New("can not delete the company").Error()))
		if err != nil {
			h.logger.Info().Timestamp().Msg(err.Error())
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte{})
	if err != nil {
		h.logger.Info().Timestamp().Msg(err.Error())
	}
}

func (h companyHandler) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	var id int
	var err error
	var companyBody dto.Company
	var company entities.Company

	id, err = strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.errCompanyId(w, err)
	}

	err = json.NewDecoder(r.Body).Decode(&companyBody)
	if err != nil {
		h.logger.Info().Timestamp().Msg(err.Error())
	}

	err = dto.Validator.Struct(companyBody)
	if err != nil {
		h.errValidateCompanyStruct(w, err)
	}

	company, err = h.companyRepository.UpdateCompany(h.ctx, id, companyBody)
	if err != nil {
		h.logger.Info().Timestamp().Msg(err.Error())
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(errors.New("error was occurred while updating the company").Error()))
		if err != nil {
			h.logger.Info().Timestamp().Msg(err.Error())
		}
		return
	}

	response, err := json.Marshal(company)
	if err != nil {
		h.logger.Info().Timestamp().Msg(err.Error())
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(errors.New("response marshal error").Error()))
		if err != nil {
			h.logger.Info().Timestamp().Msg(err.Error())
		}
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)
	if err != nil {
		h.logger.Info().Timestamp().Msg(err.Error())
	}
}

func (h companyHandler) errCompanyId(w http.ResponseWriter, err error) {
	h.logger.Info().Timestamp().Msg(err.Error())
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	_, err = w.Write([]byte(errors.New("invalid company ID").Error()))
	if err != nil {
		h.logger.Info().Timestamp().Msg(err.Error())
	}
}

func (h companyHandler) errValidateCompanyStruct(w http.ResponseWriter, err error) {
	h.logger.Info().Timestamp().Msg(err.Error())
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	_, err = w.Write([]byte(errors.New("invalid company request body").Error()))
	if err != nil {
		h.logger.Info().Timestamp().Msg(err.Error())
	}
}
