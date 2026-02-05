package huh

import (
    "strings"
    "testing"
)

func TestSelectorInit(t *testing.T) {
    s := NewSelect[string]()
    cmd := s.Init()

    if cmd != nil {
        t.Errorf("Init did not return an empty command")
    }
}

func TestSelectViewIncludesTitle(t *testing.T) {
    s := NewSelect[string]().Title("Test")
    view := s.View()

    if !strings.Contains(view, "Test") {
        t.Errorf("Expected view to contain title 'Test', got '%s'", view)
    }
}
