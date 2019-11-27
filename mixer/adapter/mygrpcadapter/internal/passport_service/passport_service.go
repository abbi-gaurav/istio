package passport_service

import (
	"istio.io/istio/mixer/adapter/mygrpcadapter/internal/db"
	"istio.io/istio/mixer/adapter/mygrpcadapter/internal/model"
	"istio.io/istio/mixer/template/authorization"
	"log"
)

type PassportService struct {
	db *db.DB
}

func New() *PassportService {
	return &PassportService{
		db: db.New(),
	}
}

func (ps *PassportService) ExtractAndStorePassportAttributes(msg *authorization.InstanceMsg) {
	log.Printf("Got msg %+v", msg)

	storageKey := model.GetStorageKey(msg.Subject.Properties)
	if storageKey == "" {
		return
	}

	passportJson, isValid := model.ParseAndValidate(msg.Subject.Properties)
	if !isValid {
		return
	}

	log.Printf("Saving the passport json %s with key %s\n", passportJson, storageKey)
	_ = ps.db.Store(storageKey, passportJson)
}
