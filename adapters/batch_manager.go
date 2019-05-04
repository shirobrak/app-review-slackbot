package adapters

// UpdateLogRepositoryIF is an interface for accessing to the updated_log repositoy.
type UpdateLogRepositoryIF interface {
	ReadLatestUpdatedDate() (string, error)
	WriteLatestUpdatedDate(lastUpdateDate string) error
}

// BatchManager is an adapter for getting log information.
type BatchManager struct {
	updateLogRepository UpdateLogRepositoryIF
}

// NewBatchManager is a method to create an instance of BatchManager
func NewBatchManager(updateLogRepository UpdateLogRepositoryIF) *BatchManager {
	return &BatchManager{updateLogRepository: updateLogRepository}
}

// GetLastUpdated is a method to get the latest review update date.
func (m *BatchManager) GetLastUpdated() (string, error) {
	lastUpdated, err := m.updateLogRepository.ReadLatestUpdatedDate()
	if err != nil {
		return "", err
	}
	return lastUpdated, nil
}

// SetLastUpdated is a method to update the latest review update date.
func (m *BatchManager) SetLastUpdated(lastUpdated string) error {
	return m.updateLogRepository.WriteLatestUpdatedDate(lastUpdated)
}
