package crontab

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/robfig/cron/v3"
	"hankquan.top/ecron/pkg/store/history"
	"log"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

const (
	ALIVE  string = "Alive"
	PAUSED string = "Paused"
)

type CronEntry struct {
	Index int
	Cron  string
	Cmd   string
	Next  string
	State string
}

func GetCronEntries() ([]CronEntry, error) {
	content, err := readCronFile()
	if err != nil {
		return nil, err
	}
	cronLines, err := transformCronEntries(content)
	if err != nil {
		return nil, fmt.Errorf("failed to parse crontab output: %v", err)
	}
	return cronLines, nil
}

func AddCronEntry(cronExpr string, cmd string) error {
	content, err := readCronFile()
	if err != nil {
		return err
	}

	defer history.FlushHistoryCache()
	history.CacheCronFile(content)

	history.AddChangeLog("add " + cronExpr + " " + cmd)
	newEntry := fmt.Sprintf("%s %s", cronExpr, cmd)

	err = writeCronFile(appendEntry(content, newEntry))
	if err != nil {
		return fmt.Errorf("failed to update crontab: %v", err)
	}
	return nil
}

func EditCronEntry(lineNumber int, cronExpr string, cmd string) error {
	content, err := readCronFile()
	if err != nil {
		return err
	}

	defer history.FlushHistoryCache()
	history.CacheCronFile(content)

	lines := strings.Split(content, "\n")

	if lineNumber < 1 || lineNumber > len(lines) {
		return fmt.Errorf("invalid line number: %d", lineNumber)
	}

	existingCronExpr, existingCmd, err := extractCronEntry(lines, lineNumber)
	if err != nil {
		return err
	}

	if len(cronExpr) > 0 {
		cronExpr = existingCronExpr
	}

	if len(existingCmd) > 0 {
		cmd = existingCmd
	}

	lines[lineNumber-1] = fmt.Sprintf("%s %s", cronExpr, cmd)

	history.AddChangeLog(fmt.Sprintf("edit %d %s %s", lineNumber, cronExpr, cmd))

	updatedContent := strings.Join(lines, "\n")
	err = writeCronFile(updatedContent)
	if err != nil {
		return err
	}

	return nil
}

func DeleteCronEntry(lineNumber int) error {
	content, err := readCronFile()
	if err != nil {
		return err
	}
	defer history.FlushHistoryCache()
	history.CacheCronFile(content)

	lines := strings.Split(content, "\n")
	if lineNumber < 1 || lineNumber > len(lines) {
		return fmt.Errorf("index %d is invalid", lineNumber)
	}

	lineToDelete := lines[lineNumber-1]

	parts := strings.Fields(lineToDelete)
	if len(parts) < 6 {
		return fmt.Errorf("invalid crontab entry at line %d: %s", lineNumber, lineToDelete)
	}

	deletedCronExpr := strings.Join(parts[:5], " ")
	deletedCmd := strings.Join(parts[5:], " ")

	// Remove the specified line
	lines = append(lines[:lineNumber-1], lines[lineNumber:]...)

	updatedContent := strings.Join(lines, "\n")

	history.AddChangeLog("delete " + deletedCronExpr + " " + deletedCmd)

	// Write the updated crontab
	err = writeCronFile(updatedContent)
	if err != nil {
		return fmt.Errorf("failed to update crontab: %v", err)
	}

	return nil
}

func StopCronEntry(lineNumber int) error {
	content, err := readCronFile()
	if err != nil {
		return err
	}

	defer history.FlushHistoryCache()
	history.CacheCronFile(content)

	lines := strings.Split(content, "\n")
	if lineNumber < 1 || lineNumber > len(lines) {
		return fmt.Errorf("index %d is invalid", lineNumber)
	}

	// Add a # at the beginning of the specified line
	lines[lineNumber-1] = "#" + lines[lineNumber-1]
	updatedContent := strings.Join(lines, "\n")

	history.AddChangeLog("stop INDEX " + lines[lineNumber-1])

	// Write the updated crontab
	err = writeCronFile(updatedContent)
	if err != nil {
		return fmt.Errorf("failed to update crontab: %v", err)
	}
	return nil
}

func StartCronEntry(lineNumber int) error {
	content, err := readCronFile()
	if err != nil {
		return err
	}

	defer history.FlushHistoryCache()
	history.CacheCronFile(content)

	lines := strings.Split(content, "\n")
	if lineNumber < 1 || lineNumber > len(lines) {
		return fmt.Errorf("index %d is invalid", lineNumber)
	}

	// Remove all # characters from the beginning of the specified line
	lines[lineNumber-1] = strings.TrimLeft(lines[lineNumber-1], "#")
	updatedContent := strings.Join(lines, "\n")

	history.AddChangeLog("start INDEX " + lines[lineNumber-1])

	// Write the updated crontab
	err = writeCronFile(updatedContent)
	if err != nil {
		return fmt.Errorf("failed to update crontab: %v", err)
	}
	return nil
}

func appendEntry(content, newEntry string) string {
	if len(content) > 0 {
		return content + "\n" + newEntry
	}
	return newEntry
}

func extractCronEntry(lines []string, lineNumber int) (string, string, error) {
	if lineNumber < 1 || lineNumber > len(lines) {
		return "", "", errors.New("line number out of range")
	}

	line := lines[lineNumber-1]
	parts := strings.Fields(line)
	if len(parts) < 6 {
		return "", "", errors.New("invalid cron format")
	}

	cronExpr := strings.Join(parts[:5], " ")
	cmd := strings.Join(parts[5:], " ")
	fmt.Println("extractCronEntry: ", cronExpr, cmd)
	return cronExpr, cmd, nil
}

func readCronFile() (string, error) {
	cmd := exec.Command("crontab", "-l")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to execute crontab -l: %v", err)
	}
	content := string(output)
	history.CacheCronFile(content)
	return strings.TrimSpace(content), nil
}
func writeCronFile(content string) error {
	content = strings.TrimRight(content, "\n") + "\n"
	command := exec.Command("crontab", "-")
	command.Stdin = bytes.NewReader([]byte(content))
	err := command.Run()
	if err != nil {
		return fmt.Errorf("failed to update crontab: %v", err)
	}
	return nil
}

func transformCronEntries(crontabOutput string) ([]CronEntry, error) {
	scanner := bufio.NewScanner(strings.NewReader(crontabOutput))
	lineNumber := 0
	var result []CronEntry

	for scanner.Scan() {
		lineNumber++
		cronLine := &CronEntry{
			State: ALIVE,
		}

		line := strings.TrimSpace(scanner.Text())

		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "#") {
			line = strings.TrimSpace(line[1:])
			parts := strings.Fields(line)
			if len(parts) >= 6 && !isValidCronExpr(strings.Join(parts[:5], " ")) {
				continue
			}
			cronLine.State = PAUSED
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

		now := time.Now()
		nextRun := c.Next(now)
		cronLine.Next = nextRun.Format("2006-01-02 15:04:05")

		// cmd 部分是第 6 个字段及其后续内容
		cmd := strings.Join(parts[5:], " ")
		cronLine.Cmd = cmd

		cronLine.Index = lineNumber
		result = append(result, *cronLine)
	}

	// 检查扫描时是否有错误
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func isValidCronExpr(cron string) bool {
	re := regexp.MustCompile(`^(\*|([0-5]?[0-9]))\s+(\*|([0-2]?[0-9]))\s+(\*|([1-9]|[1-2][0-9]|3[0-1]))\s+(\*|([1-9]|1[0-2]))\s+(\*|[0-6])`)
	return re.MatchString(cron)
}
