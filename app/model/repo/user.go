package repo

import "app/app/model/dao"

type UserRepo interface {
	IBase[dao.User]
}
