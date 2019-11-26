package passport_service

import (
	"log"

	"istio.io/istio/mixer/adapter/mygrpcadapter/internal/db"
	"istio.io/istio/mixer/adapter/mygrpcadapter/internal/model"
	"istio.io/istio/mixer/template/authorization"
)

type PassportService struct {
	marshaller *model.Marshaller
	db         *db.DB
}

func New() *PassportService {
	return &PassportService{
		marshaller: model.NewMarshaller(),
		db:         db.New(),
	}
}

func (ps *PassportService) ExtractAndStorePassportAttributes(msg *authorization.InstanceMsg) {
	log.Printf("Got msg %+v", msg)
}
