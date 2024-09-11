package helper

import (
	"user-service/src/common/errors"
	"user-service/src/model/dto"
	"user-service/src/model/entity"
	"fmt"
	"strings"
)

func BuildFindByFieldsQuery(data *entity.User) (query string, args []any, err error) {
	if data == nil {
		return "", nil, &errors.Response{HttpCode: 400, Message: "no fields to select"}
	}

	var whereClause []string
	argId := 1

	if data.UserId != "" {
		whereClause = append(whereClause, fmt.Sprintf("user_id = $%d", argId))
		args = append(args, data.UserId)
		argId++
	}

	if data.Email != "" {
		whereClause = append(whereClause, fmt.Sprintf("email = $%d", argId))
		args = append(args, data.Email)
		argId++
	}

	if data.FullName != "" {
		whereClause = append(whereClause, fmt.Sprintf("full_name = $%d", argId))
		args = append(args, data.FullName)
		argId++
	}

	if data.Role != "" {
		whereClause = append(whereClause, fmt.Sprintf("role = $%d", argId))
		args = append(args, data.Role)
		argId++
	}

	if data.Whatsapp != nil {
		whereClause = append(whereClause, fmt.Sprintf("whatsapp = $%d", argId))
		args = append(args, *data.Whatsapp)
		argId++
	}

	if data.RefreshToken != nil {
		whereClause = append(whereClause, fmt.Sprintf("refresh_token = $%d", argId))
		args = append(args, *data.RefreshToken)
		argId++
	}

	
	if len(whereClause) == 0 {
		return "", nil, &errors.Response{HttpCode: 400, Message: "no fields to select"}
	}

	query = fmt.Sprintf(`
	SELECT 
		user_id, email, full_name, whatsapp, role, password, refresh_token, created_at, updated_at 
	FROM 
		users WHERE %s LIMIT 1;
	`, strings.Join(whereClause, " AND "))

	return query, args, nil
}

func BuildUpdateUserQuery(data *dto.UpdateUserReq) (query string, args []any, err error) {
	if data == nil {
		return "", nil, &errors.Response{HttpCode: 400, Message: "no fields to select"}
	}

	var setClause []string
	argId := 1

	if data.Email != "" {
		setClause = append(setClause, fmt.Sprintf("email = $%d", argId))
		args = append(args, data.Email)
		argId++
	}

	if data.FullName != "" {
		setClause = append(setClause, fmt.Sprintf("full_name = $%d", argId))
		args = append(args, data.FullName)
		argId++
	}

	if data.Whatsapp != "" {
		setClause = append(setClause, fmt.Sprintf("whatsapp = $%d", argId))
		args = append(args, data.Whatsapp)
		argId++
	}

	if data.Password != "" {
		setClause = append(setClause, fmt.Sprintf("password = $%d", argId))
		args = append(args, data.Password)
		argId++
	}

	if data.RefreshToken != "" {
		setClause = append(setClause, fmt.Sprintf("refresh_token = $%d", argId))
		args = append(args, data.RefreshToken)
		argId++
	}

	if len(setClause) == 0 {
		return "", nil, &errors.Response{HttpCode: 400, Message: "no fields to update"}
	}

	query = fmt.Sprintf(`
	UPDATE 
		users 
	SET
		 %s, updated_at = now() 
	WHERE 
		user_id = $%d 
	RETURNING 
		user_id, email, full_name, whatsapp, role, password, refresh_token, created_at, updated_at;
	`, strings.Join(setClause, ", "), argId)

	args = append(args, data.UserId)

	return query, args, nil
}