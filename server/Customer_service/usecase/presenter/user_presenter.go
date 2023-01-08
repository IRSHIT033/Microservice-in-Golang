package presenter

import "github.com/IRSHIT033/E-comm-GO-/server/Customer_service/domain/model"

type UserPresenter interface {
	ResponseUser(u *model.User) *model.User
}
