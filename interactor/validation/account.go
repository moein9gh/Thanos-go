package validation

import (
	"github.com/p3ym4n/re"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/thanos-go/param"
)

func (s *Validate) AddAccountValidation(req *param.AddAccountRequest) re.Error {
	const op = re.Op("store.AddAccountValidation")

	err := validation.ValidateStruct(req,
		validation.Field(&req.AppVersion, validation.Required),
		validation.Field(&req.Email, validation.Required, is.Email),
	)

	if err != nil {
		return re.New(op, err, re.KindInvalid)
	}

	return nil
}
