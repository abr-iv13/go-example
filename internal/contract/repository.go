package contract

//Repository
type (
	OffersRepo interface {
		Create()
		Get()
		Update()
		Delete()
	}

	UserRepo interface {
		Create()
		Get()
		Update()
		Delete()
	}
)
