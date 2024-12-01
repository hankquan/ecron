package add

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type CronOptions struct {
	cmd      string
	minutely bool
	hourly   bool
	daily    bool
	weekly   bool
	at       string //12:01,9:05
	minute   string
	hour     string
	on       []string //monday, Monday, Mon, mon
	expr     string
	prompt   string
}

// validDays contains all valid weekday names.
var validDays = []string{
	"monday", "Monday", "Mon", "mon",
	"tuesday", "Tuesday", "Tue", "tue",
	"wednesday", "Wednesday", "Wed", "wed",
	"thursday", "Thursday", "Thu", "thu",
	"friday", "Friday", "Fri", "fri",
	"saturday", "Saturday", "Sat", "sat",
	"sunday", "Sunday", "Sun", "sun",
}

// validDaysMap is a map for quick lookup of valid weekdays.
var validDaysMap = func() map[string]struct{} {
	m := make(map[string]struct{})
	for _, day := range validDays {
		m[day] = struct{}{}
	}
	return m
}()

func (cronOptions *CronOptions) Resolve() error {
	if len(cronOptions.at) > 0 {
		parsedTime, err := time.Parse("15:04", cronOptions.at)
		if err != nil {
			return fmt.Errorf("invalid time format for 'at': %w, should be like 9:05,14:10", err)
		}
		cronOptions.minute = strconv.Itoa(parsedTime.Minute())
		cronOptions.hour = strconv.Itoa(parsedTime.Hour())
	} else {
		cronOptions.setDefaultMinuteAndHour()
	}
	return nil
}

func (cronOptions *CronOptions) setDefaultMinuteAndHour() {
	cronOptions.minute = "*"
	cronOptions.hour = "*"
}

func (cronOptions *CronOptions) Validate() error {
	trueFlagsCount := countTrueFlags(cronOptions.minutely, cronOptions.hourly, cronOptions.daily, cronOptions.weekly,
		len(cronOptions.expr) > 0, len(cronOptions.prompt) > 0)
	if trueFlagsCount > 1 {
		return fmt.Errorf("only one of --minutely,--hourly,--daily,--weekly,--expr or --prompt can be specified")
	}
	if cronOptions.weekly {
		for _, day := range cronOptions.on {
			if !isValidWeekDay(day) {
				return fmt.Errorf("invalid value of '--on': %s, week days can be: %s", day, strings.Join(validDays, ", "))
			}
		}
	}
	return nil
}

func (cronOptions *CronOptions) GenerateCronExpression() (string, error) {
	if err := cronOptions.Validate(); err != nil {
		return "", err
	}
	if err := cronOptions.Resolve(); err != nil {
		return "", err
	}
	switch {
	case cronOptions.minutely:
		return "* * * * *", nil
	case cronOptions.hourly:
		if cronOptions.minute == "*" {
			cronOptions.minute = "0"
		}
		return fmt.Sprintf("%s * * * *", cronOptions.minute), nil
	case cronOptions.daily:
		if cronOptions.minute == "*" {
			cronOptions.minute = "0"
		}
		if cronOptions.hour == "*" {
			cronOptions.hour = "0"
		}
		return fmt.Sprintf("%s %s * * *", cronOptions.minute, cronOptions.hour), nil
	case cronOptions.weekly:
		days := "Mon"
		if len(cronOptions.on) > 0 {
			days = strings.Join(cronOptions.on, ",")
		}
		return fmt.Sprintf("%s %s * * %s", cronOptions.minute, cronOptions.hour, days), nil
	case len(cronOptions.expr) > 0:
		return cronOptions.expr, nil
	case len(cronOptions.prompt) > 0:
		// Implement prompt-based cron expression generation
	}
	return "", fmt.Errorf("no valid cron expression options provided")
}

func countTrueFlags(vars ...bool) int {
	trueCount := 0
	for _, v := range vars {
		if v {
			trueCount++
		}
	}
	return trueCount
}

func isValidWeekDay(day string) bool {
	_, exists := validDaysMap[day]
	return exists
}

func (cronOptions *CronOptions) getValidWeekDays() []string {
	return validDays
}
