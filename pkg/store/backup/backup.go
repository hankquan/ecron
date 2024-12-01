package backup

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type HistoryRow struct {
	Index      int
	BackupFile string
	Changelog  string
	Timestamp  string
}

var _contentCache string
var _changelog string

func CacheCronFile(content string) {
	_contentCache = content
}

func AddChangeLog(changelog string) {
	_changelog = changelog
}

func listHistory() ([]string, error) {

	return nil, nil
}

// FlushHistoryCache flushes the history cache to a backup file and logs the operation.
// This function creates a backup directory under the user's home directory, writes the current cache content to a backup file,
// and logs the backup operation in the history.log file.
// It returns an error if any step fails.
func FlushHistoryCache() {
	// Define paths
	historyDir := filepath.Join(os.Getenv("HOME"), ".ecron")
	historyLogFile := filepath.Join(historyDir, "history.log")
	backupDir := filepath.Join(historyDir, "backup")
	timestamp := time.Now().Format("20060102_150405")
	backupFile := filepath.Join(backupDir, fmt.Sprintf("%s.bak", timestamp))

	// Create backup directory if it does not exist
	if err := os.MkdirAll(backupDir, 0700); err != nil {
		fmt.Errorf("failed to create backup directory: %v", err)
	}

	// Write current cache content to back up file
	if err := os.WriteFile(backupFile, []byte(_contentCache), 0600); err != nil {
		fmt.Errorf("failed to write backup file: %v", err)
	}

	// Open or create history log file
	f, err := os.OpenFile(historyLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Errorf("failed to open history log file: %v", err)
	}

	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	// Write log entry
	logEntry := fmt.Sprintf("%s,%s,%s\n", backupFile, _changelog, time.Now().Format("2006-01-02 15:04:05"))
	if _, err := f.WriteString(logEntry); err != nil {
		fmt.Errorf("failed to write to history log file: %v", err)
	}
}

func RollbackTo(lineNumber int) error {
	// Define paths
	historyDir := filepath.Join(os.Getenv("HOME"), ".ecron")
	historyLogFile := filepath.Join(historyDir, "history.log")

	// Read the history log file
	content, err := os.ReadFile(historyLogFile)
	if err != nil {
		return fmt.Errorf("failed to read history log file: %v", err)
	}

	// Split the content into lines
	lines := strings.Split(string(content), "\n")
	if lineNumber < 1 || lineNumber > len(lines) {
		return fmt.Errorf("index %d is invalid", lineNumber)
	}

	// Get the specified line
	line := lines[lineNumber-1]
	parts := strings.Split(line, ",")
	if len(parts) < 1 {
		return fmt.Errorf("invalid log entry at line %d: %s", lineNumber, line)
	}

	// Get the backup file path
	backupFile := parts[0]

	// Read the backup file content
	backupContent, err := os.ReadFile(backupFile)
	if err != nil {
		return fmt.Errorf("failed to read backup file: %v", err)
	}

	// Write the content to crontab
	command := exec.Command("crontab", "-")
	command.Stdin = bytes.NewReader(backupContent)
	if err := command.Run(); err != nil {
		return fmt.Errorf("failed to update crontab: %v", err)
	}

	return nil
}
