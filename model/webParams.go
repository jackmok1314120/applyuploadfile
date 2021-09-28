package model

type ApplyCompany struct {
	Name            string `json:"name"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
	CoinName        string `json:"coin_name" binding:"required" `
	Introduce       string `json:"introduce"`
	IdCardPicture   string `json:"id_card_picture" binding:"required" `
	BusinessPicture string `json:"business_picture" binding:"required" `
}
