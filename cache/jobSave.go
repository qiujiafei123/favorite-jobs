package cache

import "favorite-jobs/types"

var (
	JobSave map[int64]*types.Job // 用于和详情页通信
)

func InitCache() {
	if JobSave == nil {
		JobSave = make(map[int64]*types.Job)
	}
}
