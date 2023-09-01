package hash

func EncodeID(id uint) string {
	hd := hashids.NewData()
	hd.Salt = config.HashIDSalt()
	hd.MinLength = config.HashIDMinLength()
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{int(id)})
	return e
}

func DecodeID(id string) uint {
	hd := hashids.NewData()
	hd.Salt = config.HashIDSalt()
	hd.MinLength = config.HashIDMinLength()
	h, _ := hashids.NewWithData(hd)
	d, _ := h.DecodeWithError(id)
	return uint(d[0])
}
