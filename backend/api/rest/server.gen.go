// Package rest provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.4 DO NOT EDIT.
package rest

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get expense entries
	// (GET /v1/expenses)
	GetV1Expenses(w http.ResponseWriter, r *http.Request)
	// Create an expense entry
	// (POST /v1/expenses)
	PostV1Expenses(w http.ResponseWriter, r *http.Request)
	// Delete an expense entry by ID
	// (DELETE /v1/expenses/{expense_id})
	DeleteV1ExpensesExpenseId(w http.ResponseWriter, r *http.Request, expenseId int)
	// Get an expense entry by ID
	// (GET /v1/expenses/{expense_id})
	GetV1ExpensesExpenseId(w http.ResponseWriter, r *http.Request, expenseId int)
	// Update an expense entry by ID
	// (PUT /v1/expenses/{expense_id})
	PutV1ExpensesExpenseId(w http.ResponseWriter, r *http.Request, expenseId int)
	// Get user entries
	// (GET /v1/users/)
	GetV1Users(w http.ResponseWriter, r *http.Request)
	// Create an user entry
	// (POST /v1/users/)
	PostV1Users(w http.ResponseWriter, r *http.Request)
	// Delete an user entry by ID
	// (DELETE /v1/users/{user_id})
	DeleteV1UsersUserId(w http.ResponseWriter, r *http.Request, userId int)
	// Get an user entry by ID
	// (GET /v1/users/{user_id})
	GetV1UsersUserId(w http.ResponseWriter, r *http.Request, userId int)
	// Update an user entry by ID
	// (PUT /v1/users/{user_id})
	PutV1UsersUserId(w http.ResponseWriter, r *http.Request, userId int)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// Get expense entries
// (GET /v1/expenses)
func (_ Unimplemented) GetV1Expenses(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Create an expense entry
// (POST /v1/expenses)
func (_ Unimplemented) PostV1Expenses(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete an expense entry by ID
// (DELETE /v1/expenses/{expense_id})
func (_ Unimplemented) DeleteV1ExpensesExpenseId(w http.ResponseWriter, r *http.Request, expenseId int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get an expense entry by ID
// (GET /v1/expenses/{expense_id})
func (_ Unimplemented) GetV1ExpensesExpenseId(w http.ResponseWriter, r *http.Request, expenseId int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update an expense entry by ID
// (PUT /v1/expenses/{expense_id})
func (_ Unimplemented) PutV1ExpensesExpenseId(w http.ResponseWriter, r *http.Request, expenseId int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get user entries
// (GET /v1/users/)
func (_ Unimplemented) GetV1Users(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Create an user entry
// (POST /v1/users/)
func (_ Unimplemented) PostV1Users(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete an user entry by ID
// (DELETE /v1/users/{user_id})
func (_ Unimplemented) DeleteV1UsersUserId(w http.ResponseWriter, r *http.Request, userId int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get an user entry by ID
// (GET /v1/users/{user_id})
func (_ Unimplemented) GetV1UsersUserId(w http.ResponseWriter, r *http.Request, userId int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update an user entry by ID
// (PUT /v1/users/{user_id})
func (_ Unimplemented) PutV1UsersUserId(w http.ResponseWriter, r *http.Request, userId int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetV1Expenses operation middleware
func (siw *ServerInterfaceWrapper) GetV1Expenses(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetV1Expenses(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostV1Expenses operation middleware
func (siw *ServerInterfaceWrapper) PostV1Expenses(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostV1Expenses(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteV1ExpensesExpenseId operation middleware
func (siw *ServerInterfaceWrapper) DeleteV1ExpensesExpenseId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "expense_id" -------------
	var expenseId int

	err = runtime.BindStyledParameterWithLocation("simple", false, "expense_id", runtime.ParamLocationPath, chi.URLParam(r, "expense_id"), &expenseId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "expense_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteV1ExpensesExpenseId(w, r, expenseId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetV1ExpensesExpenseId operation middleware
func (siw *ServerInterfaceWrapper) GetV1ExpensesExpenseId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "expense_id" -------------
	var expenseId int

	err = runtime.BindStyledParameterWithLocation("simple", false, "expense_id", runtime.ParamLocationPath, chi.URLParam(r, "expense_id"), &expenseId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "expense_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetV1ExpensesExpenseId(w, r, expenseId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PutV1ExpensesExpenseId operation middleware
func (siw *ServerInterfaceWrapper) PutV1ExpensesExpenseId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "expense_id" -------------
	var expenseId int

	err = runtime.BindStyledParameterWithLocation("simple", false, "expense_id", runtime.ParamLocationPath, chi.URLParam(r, "expense_id"), &expenseId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "expense_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutV1ExpensesExpenseId(w, r, expenseId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetV1Users operation middleware
func (siw *ServerInterfaceWrapper) GetV1Users(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetV1Users(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostV1Users operation middleware
func (siw *ServerInterfaceWrapper) PostV1Users(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostV1Users(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteV1UsersUserId operation middleware
func (siw *ServerInterfaceWrapper) DeleteV1UsersUserId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId int

	err = runtime.BindStyledParameterWithLocation("simple", false, "user_id", runtime.ParamLocationPath, chi.URLParam(r, "user_id"), &userId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "user_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteV1UsersUserId(w, r, userId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetV1UsersUserId operation middleware
func (siw *ServerInterfaceWrapper) GetV1UsersUserId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId int

	err = runtime.BindStyledParameterWithLocation("simple", false, "user_id", runtime.ParamLocationPath, chi.URLParam(r, "user_id"), &userId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "user_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetV1UsersUserId(w, r, userId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PutV1UsersUserId operation middleware
func (siw *ServerInterfaceWrapper) PutV1UsersUserId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId int

	err = runtime.BindStyledParameterWithLocation("simple", false, "user_id", runtime.ParamLocationPath, chi.URLParam(r, "user_id"), &userId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "user_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutV1UsersUserId(w, r, userId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/v1/expenses", wrapper.GetV1Expenses)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/v1/expenses", wrapper.PostV1Expenses)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/v1/expenses/{expense_id}", wrapper.DeleteV1ExpensesExpenseId)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/v1/expenses/{expense_id}", wrapper.GetV1ExpensesExpenseId)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/v1/expenses/{expense_id}", wrapper.PutV1ExpensesExpenseId)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/v1/users/", wrapper.GetV1Users)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/v1/users/", wrapper.PostV1Users)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/v1/users/{user_id}", wrapper.DeleteV1UsersUserId)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/v1/users/{user_id}", wrapper.GetV1UsersUserId)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/v1/users/{user_id}", wrapper.PutV1UsersUserId)
	})

	return r
}
