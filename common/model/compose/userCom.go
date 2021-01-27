package compose

import (
	"IndustrialInternetApi/model"
	"IndustrialInternetApi/model/paginate"
)

type UserComposer struct {
	Users []model.User
	paginate.Pagination
}
