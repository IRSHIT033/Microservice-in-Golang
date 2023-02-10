package domain_user

type ProductIdRequest struct {
	ProductId uint `json:"productid" binding:"required"`
}
