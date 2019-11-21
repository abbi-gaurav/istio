package passport_service

import (
	"istio.io/istio/mixer/adapter/mygrpcadapter/internal/db"
	"istio.io/istio/mixer/adapter/mygrpcadapter/internal/model"
	"istio.io/istio/mixer/template/metric"
	"log"
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

func (ps *PassportService) ExtractAndStorePassportAttributes(msgs []*metric.InstanceMsg) {
	storageKey := model.GetStorageKey(msgs)
	if storageKey == "" {
		return
	}

	passportJson, isValid := ps.marshaller.ParseAndValidate(msgs)
	if !isValid {
		return
	}

	log.Printf("Saving the passport json %s with key %s\n", passportJson, storageKey)
	_ = ps.db.Store(storageKey, passportJson)
}
