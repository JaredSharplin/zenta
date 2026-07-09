package breathing

import (
	"testing"
)

func TestNewSession(t *testing.T) {
	s := NewSession()

	if s.Cycles != 4 {
		t.Errorf("Expected default cycles to be 4, got %d", s.Cycles)
	}
	if s.ShowQuote {
		t.Errorf("Expected ShowQuote to be false by default, got %v", s.ShowQuote)
	}
	if s.InhaleDur != 4 {
		t.Errorf("Expected InhaleDur to be 4, got %d", s.InhaleDur)
	}
	if s.HoldDur != 7 {
		t.Errorf("Expected HoldDur to be 7, got %d", s.HoldDur)
	}
	if s.ExhaleDur != 8 {
		t.Errorf("Expected ExhaleDur to be 8, got %d", s.ExhaleDur)
	}
	if s.SimpleMode != shouldUseSimpleAnimation() {
		t.Errorf("Expected SimpleMode to match default from shouldUseSimpleAnimation()")
	}
}

func TestParseArgs(t *testing.T) {
	defaultSimple := shouldUseSimpleAnimation()

	testCases := []struct {
		name           string
		args           []string
		expectedCycles int
		expectedQuote  bool
		expectedSimple bool
	}{
		{"no args", []string{}, 4, false, defaultSimple},
		{"quick", []string{"--quick"}, 1, false, defaultSimple},
		{"extended", []string{"--extended"}, 5, false, defaultSimple},
		{"silent", []string{"--silent"}, 4, false, defaultSimple},
		{"simple", []string{"--simple"}, 4, false, true},
		{"complex", []string{"--complex"}, 4, false, false},
		{"-q", []string{"-q"}, 1, false, defaultSimple},
		{"-e", []string{"-e"}, 5, false, defaultSimple},
		{"-s", []string{"-s"}, 4, false, defaultSimple},
		{"combo", []string{"--quick", "--silent", "--simple"}, 1, false, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := NewSession()
			s.ParseArgs(tc.args)

			if s.Cycles != tc.expectedCycles {
				t.Errorf("For args %v, expected cycles to be %d, got %d", tc.args, tc.expectedCycles, s.Cycles)
			}
			if s.ShowQuote != tc.expectedQuote {
				t.Errorf("For args %v, expected ShowQuote to be %v, got %v", tc.args, tc.expectedQuote, s.ShowQuote)
			}
			if s.SimpleMode != tc.expectedSimple {
				t.Errorf("For args %v, expected SimpleMode to be %v, got %v", tc.args, tc.expectedSimple, s.SimpleMode)
			}
		})
	}
}

func TestShouldShowQuote(t *testing.T) {
	s := NewSession()

	s.ShowQuote = true
	if !s.ShouldShowQuote() {
		t.Error("ShouldShowQuote should return true when s.ShowQuote is true")
	}

	s.ShowQuote = false
	if s.ShouldShowQuote() {
		t.Error("ShouldShowQuote should return false when s.ShowQuote is false")
	}
}
