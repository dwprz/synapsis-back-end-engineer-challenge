package dto

import "user-service/src/model/entity"

type LoginRes struct {
	Data   *entity.SanitizedUser `json:"data"`
	Tokens *entity.Tokens        `json:"tokens"`
}