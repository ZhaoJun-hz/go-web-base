package api

import (
	"errors"
	"fmt"
	"github.com/ZhaoJun-hz/go-web-base/global"
	"github.com/ZhaoJun-hz/go-web-base/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"reflect"
)

type BaseApi struct {
	Ctx    *gin.Context
	Errors error
	Logger *zap.SugaredLogger
}

func NewBaseApi() BaseApi {
	return BaseApi{
		Logger: global.Logger,
	}

}

type BuildRequestOption struct {
	Ctx               *gin.Context
	DTO               any
	BindParamsFromUri bool
}

func (baseApi *BaseApi) BuildRequest(option BuildRequestOption) *BaseApi {
	var errResult error
	// 绑定请求上下文
	baseApi.Ctx = option.Ctx

	// 绑定请求数据
	if option.DTO != nil {
		if option.BindParamsFromUri {
			errResult = baseApi.Ctx.ShouldBindUri(option.DTO)
		} else {
			errResult = baseApi.Ctx.ShouldBind(option.DTO)
		}

		if errResult != nil {
			errResult = baseApi.ParseValidateErrors(errResult, option.DTO)
			baseApi.AddError(errResult)
			baseApi.Fail(Response{
				Msg: baseApi.GetError().Error(),
			})
		}
	}
	return baseApi
}

func (baseApi *BaseApi) AddError(errNew error) {
	baseApi.Errors = utils.AppendError(baseApi.Errors, errNew)
}

func (baseApi *BaseApi) GetError() error {
	return baseApi.Errors
}

func (baseApi *BaseApi) ParseValidateErrors(errs error, target any) error {
	var errResult error

	validationErrors, ok := errs.(validator.ValidationErrors)
	if !ok {
		return errs
	}

	fields := reflect.TypeOf(target).Elem()
	for _, fieldError := range validationErrors {
		field, _ := fields.FieldByName(fieldError.Field())
		errorMessageTag := fmt.Sprintf("%s_err", fieldError.Tag())
		errorMessage := field.Tag.Get(errorMessageTag)
		if errorMessage == "" {
			errorMessage = field.Tag.Get("message")
		}

		if errorMessage == "" {
			errorMessage = fmt.Sprintf("%s : %s Error", fieldError.Field(), fieldError.Tag())
		}
		errResult = utils.AppendError(errResult, errors.New(errorMessage))
	}
	return errResult
}

func (baseApi *BaseApi) Fail(resp Response) {
	Fail(baseApi.Ctx, resp)
}

func (baseApi *BaseApi) OK(resp Response) {
	OK(baseApi.Ctx, resp)
}

func (baseApi *BaseApi) ServerFail(resp Response) {
	ServerFail(baseApi.Ctx, resp)
}
