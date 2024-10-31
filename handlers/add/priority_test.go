package adder

import (
	"fmt"
	"testing"
)

func TestGetPriority(t *testing.T) {
	tests := []struct {
		setPriority      int
		tagPriority      int
		expectedPriority int
		expectedError    error
	}{
		{
			setPriority:      -1,
			tagPriority:      -1,
			expectedPriority: 0,
			expectedError:    nil,
		},
		{
			setPriority:      -10,
			tagPriority:      5,
			expectedPriority: 0,
			expectedError:    fmt.Errorf("Priority must be in the range of 0-10"),
		},
		{
			setPriority:      7,
			tagPriority:      5,
			expectedPriority: 7,
			expectedError:    nil,
		},
		{
			setPriority:      20,
			tagPriority:      -1,
			expectedPriority: 0,
			expectedError:    fmt.Errorf("Priority must be in the range of 0-10"),
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("setPriority=%d tagPriority=%d", tt.setPriority, tt.tagPriority), func(t *testing.T) {
			priority, err := getPriority(tt.setPriority, tt.tagPriority)

			if err != nil && tt.expectedError == nil {
				t.Fatalf("expected no error, but got %v", err)
			}

			if err == nil && tt.expectedError != nil {
				t.Fatalf("expected error %v, but got nil", tt.expectedError)
			}

			if err != nil && tt.expectedError != nil && err.Error() != tt.expectedError.Error() {
				t.Fatalf("expected error %v, but got %v", tt.expectedError, err)
			}

			if priority != tt.expectedPriority {
				t.Fatalf("expected priority %d, but got %d", tt.expectedPriority, priority)
			}
		})
	}
}
