package zeewImg


type Img struct {
	Token string
	Card Card
}

func NewImg(token string) *Img {
	return &Img{
		Token: token,
		Card:  NewCards(token),
	}
}
