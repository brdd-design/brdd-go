package brdd

import (
	"testing"
)

type TestData struct {
	ID int
}

func TestDefaultExecutionContext_Initialization(t *testing.T) {
	ctx := NewExecutionContext(TestData{ID: 1})

	if ctx.GetData().ID != 1 {
		t.Errorf("Expected Data ID to be 1, got %v", ctx.GetData().ID)
	}
	if !ctx.IsValid() {
		t.Errorf("Expected IsValid to be true on initialization")
	}
	if ctx.GetStatus() != 200 {
		t.Errorf("Expected Status to be 200, got %v", ctx.GetStatus())
	}
	if len(ctx.GetErrors()) != 0 {
		t.Errorf("Expected 0 errors, got %v", len(ctx.GetErrors()))
	}
}

func TestDefaultExecutionContext_AddError(t *testing.T) {
	ctx := NewExecutionContext(TestData{})
	ctx.AddError("R001", "Invalid test")

	if ctx.IsValid() {
		t.Errorf("Expected IsValid to be false after adding error")
	}
	if ctx.GetStatus() != 400 {
		t.Errorf("Expected Status to be 400, got %v", ctx.GetStatus())
	}
	if len(ctx.GetErrors()) != 1 {
		t.Errorf("Expected 1 error, got %v", len(ctx.GetErrors()))
	}
	if ctx.GetErrors()[0].Code != "R001" {
		t.Errorf("Expected error code R001, got %v", ctx.GetErrors()[0].Code)
	}
}

func TestDefaultExecutionContext_AddEffectAndSetter(t *testing.T) {
	ctx := NewExecutionContext(TestData{})
	ctx.AddEffect("E001")
	ctx.AddSetter("S001")

	if len(ctx.GetEffects()) != 1 || ctx.GetEffects()[0] != "E001" {
		t.Errorf("Expected effect E001 to be added")
	}
	if len(ctx.GetSetters()) != 1 || ctx.GetSetters()[0] != "S001" {
		t.Errorf("Expected setter S001 to be added")
	}
}

func TestDefaultExecutionContext_SetData(t *testing.T) {
	ctx := NewExecutionContext(TestData{ID: 0})
	ctx.SetData(TestData{ID: 2})

	if ctx.GetData().ID != 2 {
		t.Errorf("Expected Data ID to be 2 after SetData")
	}
}
