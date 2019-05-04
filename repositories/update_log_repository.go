package repositories

import (
	"bufio"
	"fmt"
	"os"
)

// UpdateLogRepository is repository for accessing to the updated log file.
type UpdateLogRepository struct {
	updateLogPath string
}

// NewUpdateLogRepository is a method to create an instance of UpdateLogRepository.
func NewUpdateLogRepository() *UpdateLogRepository {
	return &UpdateLogRepository{updateLogPath: os.Getenv("UPDATED_LOG_PATH")}
}

// ReadLatestUpdatedDate is a method to get the latest review update date from file.
func (r *UpdateLogRepository) ReadLatestUpdatedDate() (string, error) {
	fp, err := os.Open(r.updateLogPath)
	if err != nil {
		return "", err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	var targetText string
	for scanner.Scan() {
		// 最初の1行だけを読み込む
		targetText = scanner.Text()
		break
	}
	return targetText, nil
}

// WriteLatestUpdatedDate is a method to write the latest review update date to file.
func (r *UpdateLogRepository) WriteLatestUpdatedDate(lastUpdateDate string) error {
	fpw, err := os.Create(r.updateLogPath)
	if err != nil {
		return err
	}
	defer fpw.Close()

	w := bufio.NewWriter(fpw)
	_, err = fmt.Fprint(w, lastUpdateDate)
	if err != nil {
		return err
	}
	return w.Flush()
}
