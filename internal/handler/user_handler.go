package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"

	"solver/internal/model"

	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user *model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		h.logger.Logger.Error("parse request", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err = h.useCase.CreateUser(r.Context(), user)
	if err != nil {
		h.logger.Logger.Error("create user failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) IncreaseBalanceHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(r.Header.Get("user-id"))
	if err != nil {
		h.logger.Logger.Error("parse user-id header", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var moneyTransfer decimal.Decimal
	err = json.NewDecoder(r.Body).Decode(&moneyTransfer)
	if err != nil {
		h.logger.Logger.Error("parse request", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := h.useCase.IncreaseUserBalance(r.Context(), userID, moneyTransfer)
	if err != nil {
		h.logger.Logger.Error("increase user balance failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) SolveTheTaskHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(r.Header.Get("user-id"))
	if err != nil {
		h.logger.Logger.Error("parse user-id header", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var arrayToSort []uint32
	err = json.NewDecoder(r.Body).Decode(&arrayToSort)
	if err != nil {
		h.logger.Logger.Error("parse request", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sortedArray, err := h.useCase.SolveTask(r.Context(), userID, arrayToSort)
	if err != nil {
		h.logger.Logger.Error("solve task failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(sortedArray); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}
