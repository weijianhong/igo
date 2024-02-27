package model

type Model struct {
	ID int64 `json:"id"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
	DeletedAt int64 `json:"deleted_at"`
}

type Option struct {
	CreatorId int64 `json:"creator_id"`
	UpdaterId int64 `json:"updater_id"`
}
