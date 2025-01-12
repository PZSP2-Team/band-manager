package domain

type SystemStats struct {
	TotalUsers  int64 `json:"total_users"`
	TotalGroups int64 `json:"total_groups"`
}
