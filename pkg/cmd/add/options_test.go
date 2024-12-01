package add

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateCronExpression(t *testing.T) {
	tests := []struct {
		name        string
		addOptions  CronOptions
		expected    string
		expectError bool
	}{
		{
			name: "Minutely",
			addOptions: CronOptions{
				minutely: true,
			},
			expected:    "* * * * *",
			expectError: false,
		},
		{
			name: "Hourly",
			addOptions: CronOptions{
				hourly: true,
			},
			expected:    "0 * * * *",
			expectError: false,
		},
		{
			name: "Daily",
			addOptions: CronOptions{
				daily: true,
			},
			expected:    "0 0 * * *",
			expectError: false,
		},
		{
			name: "Daily with at",
			addOptions: CronOptions{
				daily: true,
				at:    "14:05",
			},
			expected:    "5 14 * * *",
			expectError: false,
		},
		{
			name: "Weekly",
			addOptions: CronOptions{
				weekly: true,
			},
			expected:    "* * * * Mon",
			expectError: false,
		},
		{
			name: "Weekly with on",
			addOptions: CronOptions{
				weekly: true,
				on:     []string{"Mon", "Wed", "Fri"},
			},
			expected:    "* * * * Mon,Wed,Fri",
			expectError: false,
		},
		{
			name: "Weekly with on and at",
			addOptions: CronOptions{
				weekly: true,
				on:     []string{"Mon", "Wed", "Fri"},
				at:     "10:05",
			},
			expected:    "5 10 * * Mon,Wed,Fri",
			expectError: false,
		},
		{
			name: "Hourly with at",
			addOptions: CronOptions{
				hourly: true,
				at:     "00:05",
			},
			expected:    "5 * * * *",
			expectError: false,
		},
		{
			name: "Daily with only minute in at",
			addOptions: CronOptions{
				daily: true,
				at:    "00:05",
			},
			expected:    "5 0 * * *",
			expectError: false,
		},
		{
			name: "Daily with AM/PM in at",
			addOptions: CronOptions{
				daily: true,
				at:    "14:05",
			},
			expected:    "5 14 * * *",
			expectError: false,
		},
		{
			name: "Weekly with AM/PM in at",
			addOptions: CronOptions{
				weekly: true,
				at:     "14:05",
				on:     []string{"Mon", "Wed", "Fri"},
			},
			expected:    "5 14 * * Mon,Wed,Fri",
			expectError: false,
		},
		{
			name: "Invalid options",
			addOptions: CronOptions{
				cmd: "invalid",
			},
			expected:    "",
			expectError: true,
		},
		{
			name: "Multiple valid options",
			addOptions: CronOptions{
				minutely: true,
				hourly:   true,
			},
			expected:    "",
			expectError: true,
		},
		{
			name: "Weekly with empty on",
			addOptions: CronOptions{
				weekly: true,
				on:     []string{},
			},
			expected:    "* * * * Mon",
			expectError: false,
		},
		{
			name: "Weekly with invalid on",
			addOptions: CronOptions{
				weekly: true,
				on:     []string{"Mon", "Foo", "Fri"},
			},
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.addOptions.GenerateCronExpression()
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
