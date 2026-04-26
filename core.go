package brdd

import (
	"context"
)

// BRDDError represents a standardized business rule violation.
type BRDDError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ValidationContext is a subset of the context for pure validation logic.
type ValidationContext interface {
	AddError(code, message string)
	IsValid() bool
	GetErrors() []BRDDError
}

// ExecutionContext is the central state object interface returned by UseCases.
type ExecutionContext[T any] interface {
	ValidationContext
	GetData() T
	SetData(data T)
	AddEffect(code string)
	AddSetter(code string)
	GetEffects() []string
	GetSetters() []string
	GetStatus() int
}

// DefaultExecutionContext is the default implementation.
type DefaultExecutionContext[T any] struct {
	Data    T           `json:"data"`
	Errors  []BRDDError `json:"errors"`
	Setters []string    `json:"setters"`
	Effects []string    `json:"effects"`
	Status  int         `json:"status"`
}

// NewExecutionContext creates a new context with default successful status.
func NewExecutionContext[T any](initialData T) *DefaultExecutionContext[T] {
	return &DefaultExecutionContext[T]{
		Data:    initialData,
		Errors:  []BRDDError{},
		Setters: []string{},
		Effects: []string{},
		Status:  200,
	}
}

// AddError appends a business error and updates the status.
func (c *DefaultExecutionContext[T]) AddError(code, message string) {
	c.Errors = append(c.Errors, BRDDError{Code: code, Message: message})
	c.Status = 400
}

// AddEffect records an external effect.
func (c *DefaultExecutionContext[T]) AddEffect(code string) {
	c.Effects = append(c.Effects, code)
}

// AddSetter records a state change.
func (c *DefaultExecutionContext[T]) AddSetter(code string) {
	c.Setters = append(c.Setters, code)
}

// SetData updates the payload.
func (c *DefaultExecutionContext[T]) SetData(data T) {
	c.Data = data
}

// GetData retrieves the payload.
func (c *DefaultExecutionContext[T]) GetData() T {
	return c.Data
}

// GetErrors retrieves the errors.
func (c *DefaultExecutionContext[T]) GetErrors() []BRDDError {
	return c.Errors
}

// GetEffects retrieves the effects.
func (c *DefaultExecutionContext[T]) GetEffects() []string {
	return c.Effects
}

// GetSetters retrieves the setters.
func (c *DefaultExecutionContext[T]) GetSetters() []string {
	return c.Setters
}

// GetStatus retrieves the HTTP status.
func (c *DefaultExecutionContext[T]) GetStatus() int {
	return c.Status
}

// IsValid checks if there are any errors.
func (c *DefaultExecutionContext[T]) IsValid() bool {
	return len(c.Errors) == 0
}

// ValidateService is a protocol for pure business logic validation.
type ValidateService[I any] interface {
	Validate(ctx context.Context, validationCtx ValidationContext, input I)
}

// EnrichService is a protocol for fetching additional data.
type EnrichService[I any, E any, T any] interface {
	Enrich(ctx context.Context, execCtx ExecutionContext[T], input I) (E, error)
}

// ClientService is a protocol for external adapters.
type ClientService[I any, T any] interface {
	Execute(ctx context.Context, execCtx ExecutionContext[T], input I)
}

// UseCase is the orchestrator interface.
type UseCase[I any, O any] interface {
	Execute(ctx context.Context, input I) ExecutionContext[O]
}

