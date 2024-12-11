package handlers

import (
	"band-manager-backend/internal/usecases"
)

type GroupHandler struct {
	groupUsecase *usecases.GroupUsecase
}

func NewGroupHandler() *GroupHandler {
	groupUsecase := usecases.NewGroupUsecase()
	return &GroupHandler{
		groupUsecase: groupUsecase,
	}
}
