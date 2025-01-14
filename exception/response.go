package exception

import (
	"database/sql"
	"errors"
	"fmt"
	"runtime"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rizalarfiyan/be-tilik-jalan/constant"
	"github.com/rizalarfiyan/be-tilik-jalan/internal/response"
	"github.com/rizalarfiyan/be-tilik-jalan/logger"
	"github.com/rizalarfiyan/be-tilik-jalan/validation"
	"github.com/rs/zerolog"
)

type exception struct {
	log *zerolog.Logger
}

type Exception interface {
	Error(err error)
	ErrorSkipNotFound(err error)
	ManualValidation(key, message string, messages ...string)
	ManualValidationErr(err error, key, message string, messages ...string)
	ManualValidationBool(isError bool, key, message string, messages ...string)
	ManualValidations(errs map[string]string, messages ...string)
	ManualValidationsErr(err error, errs map[string]string, messages ...string)
	ManualValidationsBool(isError bool, errs map[string]string, messages ...string)
	Unauthorized(messages ...string)
	UnauthorizedErr(err error, messages ...string)
	UnauthorizedBool(isError bool, messages ...string)
	NotFound(messages ...string)
	NotFoundBool(isFound bool, messages ...string)
	SelectQuery(err error, messages ...string)
	BadRequest(messages ...string)
	BadRequestErr(err error, messages ...string)
	ValidateStruct(dataSet interface{}, fullPathPrefix ...bool)
	Forbidden(messages ...string)
	ForbiddenBool(isForbidden bool, messages ...string)
	UnprocessableEntity(messages ...string)
	UnprocessableEntityErr(err error, messages ...string)
	UnprocessableEntityBool(isUnprocessableEntity bool, messages ...string)
}

func NewException() Exception {
	return &exception{
		log: logger.GetWithoutCaller("exception"),
	}
}

func (e *exception) getCaller(skips ...int) string {
	skip := 2
	if len(skips) > 0 {
		skip = skips[0]
	}

	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return ""
	}

	return fmt.Sprintf("%s:%d", file, line)
}

func (e *exception) getMessage(defaultMessage string, messages ...string) string {
	if len(messages) > 0 {
		return messages[0]
	}

	return defaultMessage
}

func (e *exception) baseError(err error) {
	e.log.Error().Str("caller", e.getCaller(3)).Err(err).Msg("SERVER ERROR")
	panic(response.NewErrorMessage(fiber.StatusInternalServerError, constant.MsgInternalServerError, nil))
}

func (e *exception) Error(err error) {
	if err != nil {
		e.baseError(err)
	}
}

func (e *exception) ErrorSkipNotFound(err error) {
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		e.baseError(err)
	}
}

func (e *exception) baseErrorValidation(data map[string]string, messages ...string) {
	e.log.Error().Str("caller", e.getCaller(4)).Msg("CLIENT ERROR")
	panic(response.NewErrorMessage(fiber.StatusBadRequest, e.getMessage(constant.MsgErrorValidation, messages...), data))
}

func (e *exception) ManualValidation(key, message string, messages ...string) {
	e.baseErrorValidation(map[string]string{
		key: message,
	}, messages...)
}

func (e *exception) ManualValidationErr(err error, key, message string, messages ...string) {
	if err != nil {
		e.baseErrorValidation(map[string]string{
			key: message,
		}, messages...)
	}
}

func (e *exception) ManualValidationBool(isError bool, key, message string, messages ...string) {
	if isError {
		e.baseErrorValidation(map[string]string{
			key: message,
		}, messages...)
	}
}

func (e *exception) ManualValidations(errs map[string]string, messages ...string) {
	e.baseErrorValidation(errs, messages...)
}

func (e *exception) ManualValidationsErr(err error, errs map[string]string, messages ...string) {
	if err != nil {
		e.baseErrorValidation(errs, messages...)
	}
}

func (e *exception) ManualValidationsBool(isError bool, errs map[string]string, messages ...string) {
	if isError {
		e.baseErrorValidation(errs, messages...)
	}
}

func (e *exception) baseUnauthorized(messages ...string) {
	e.log.Warn().Str("caller", e.getCaller(3)).Msg("UNAUTHORIZED")
	panic(response.NewErrorMessage(fiber.StatusUnauthorized, e.getMessage(constant.MsgUnauthorized, messages...), nil))
}

func (e *exception) Unauthorized(messages ...string) {
	e.baseUnauthorized(messages...)
}

func (e *exception) UnauthorizedErr(err error, messages ...string) {
	if err != nil {
		e.baseUnauthorized(messages...)
	}
}

func (e *exception) UnauthorizedBool(isError bool, messages ...string) {
	if isError {
		e.baseUnauthorized(messages...)
	}
}

func (e *exception) baseNotFound(messages ...string) {
	e.log.Warn().Str("caller", e.getCaller(4)).Msg("NOT FOUND")
	panic(response.NewErrorMessage(fiber.StatusNotFound, e.getMessage(constant.MsgNotFound, messages...), nil))
}

func (e *exception) NotFound(messages ...string) {
	e.baseNotFound(messages...)
}

func (e *exception) NotFoundBool(isFound bool, messages ...string) {
	if !isFound {
		e.baseNotFound(messages...)
	}
}

func (e *exception) SelectQuery(err error, messages ...string) {
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			e.baseNotFound(messages...)
			return
		}

		e.baseError(err)
	}
}

func (e *exception) baseBadRequest(messages ...string) {
	e.log.Warn().Str("caller", e.getCaller(4)).Msg("BAD_REQUEST")
	panic(response.NewErrorMessage(fiber.StatusBadRequest, e.getMessage(constant.MsgBadRequest, messages...), nil))
}

func (e *exception) BadRequest(messages ...string) {
	e.baseBadRequest(messages...)
}

func (e *exception) BadRequestErr(err error, messages ...string) {
	if err != nil {
		e.baseBadRequest(messages...)
	}
}

func (e *exception) ValidateStruct(dataSet interface{}, fullPathPrefix ...bool) {
	err := validation.Get().Struct(dataSet)
	if err == nil {
		return
	}

	removePrefix := func(s string) string {
		for i := 0; i < len(s); i++ {
			if s[i] == '.' {
				return s[i+1:]
			}
		}
		return s
	}

	if len(fullPathPrefix) > 0 {
		removePrefix = func(s string) string {
			return s
		}
	}

	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		validates := make(map[string]string, len(ve))
		translate := validation.GetTranslator()
		for _, fe := range ve {
			key := removePrefix(fe.Namespace())
			validates[key] = fe.Translate(translate)
		}

		e.baseErrorValidation(validates)
		return
	}

	e.baseError(err)
}

func (e *exception) baseForbidden(messages ...string) {
	e.log.Warn().Str("caller", e.getCaller(4)).Msg("FORBIDDEN")
	panic(response.NewErrorMessage(fiber.StatusForbidden, e.getMessage(constant.MsgForbidden, messages...), nil))
}

func (e *exception) Forbidden(messages ...string) {
	e.baseForbidden(messages...)
}

func (e *exception) ForbiddenBool(isForbidden bool, messages ...string) {
	if isForbidden {
		e.baseForbidden(messages...)
	}
}

func (e *exception) baseUnprocessableEntity(messages ...string) {
	e.log.Warn().Str("caller", e.getCaller(4)).Msg("NOT PROCESS")
	panic(response.NewErrorMessage(fiber.StatusUnprocessableEntity, e.getMessage(constant.MsgUnprocessableEntity, messages...), nil))
}

func (e *exception) UnprocessableEntity(messages ...string) {
	e.baseUnprocessableEntity(messages...)
}

func (e *exception) UnprocessableEntityErr(err error, messages ...string) {
	if err != nil {
		e.baseUnprocessableEntity(messages...)
	}
}

func (e *exception) UnprocessableEntityBool(isUnprocessableEntity bool, messages ...string) {
	if isUnprocessableEntity {
		e.baseUnprocessableEntity(messages...)
	}
}
