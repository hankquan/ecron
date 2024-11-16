package crontab

import (
	"bufio"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

type STATE int

const (
	a = "1"
)

const (
	ALIVE STATE = iota
	PAUSED
)

func (state STATE) String() string {
	return [...]string{"Alive", "Paused"}[state]
}

type CronJobRow struct {
	Index int
	Cron  string
	Cmd   string
	Next  string
	State string
}

func GetCronElements() []CronJobRow {
	content, err := getCrontabList()
	if err != nil {
		log.Fatal("bad!", err)
	}
	cronLines, err := parseCrontab(content)
	if err != nil {
		log.Fatalf("Error reading crontab output: %v\n", err)
	}
	return cronLines
}

// get raw crontab -l output
func getCrontabList() (string, error) {
	cmd := exec.Command("crontab", "-l")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to execute crontab -l: %v", err)
	}
	return string(output), nil
}

// parse raw output to CronJobRow elements
func parseCrontab(crontabOutput string) ([]CronJobRow, error) {
	scanner := bufio.NewScanner(strings.NewReader(crontabOutput))
	lineNumber := 0
	var result []CronJobRow

	for scanner.Scan() {
		cronLine := &CronJobRow{
			State: ALIVE.String(),
		}

		line := strings.TrimSpace(scanner.Text())

		// 跳过空行
		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "#") {
			line = strings.TrimSpace(line[1:])
			parts := strings.Fields(line)
			if len(parts) >= 6 && !isValidCronExpression(strings.Join(parts[:5], " ")) {
				continue
			}
			cronLine.State = PAUSED.String()
		}

		// 分割 cronExpression 表达式和命令
		parts := strings.Fields(line)
		if len(parts) < 6 {
			fmt.Printf("Invalid crontab entry at line %d: %s\n", lineNumber, line)
			continue
		}

		// cronExpression 部分是前 5 个字段
		cronExpression := strings.Join(parts[:5], " ")
		cronLine.Cron = cronExpression

		c, err := cron.ParseStandard(cronExpression)
		if err != nil {
			log.Fatal("Error parsing cron expression: ", err)
		}

		// 获取当前时间
		now := time.Now()

		// 计算下次执行时间
		nextRun := c.Next(now)
		cronLine.Next = nextRun.Format("2006-01-02 15:04:05")

		// cmd 部分是第 6 个字段及其后续内容
		cmd := strings.Join(parts[5:], " ")
		cronLine.Cmd = cmd

		lineNumber++
		cronLine.Index = lineNumber
		result = append(result, *cronLine)
	}

	// 检查扫描时是否有错误
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func isValidCronExpression(cron string) bool {
	re := regexp.MustCompile(`^(\*|([0-5]?[0-9]))\s+(\*|([0-2]?[0-9]))\s+(\*|([1-9]|[1-2][0-9]|3[0-1]))\s+(\*|([1-9]|1[0-2]))\s+(\*|[0-6])`)
	return re.MatchString(cron)
}
