package services

import (
	"github.com/spiffe/sri/pkg/common"
	ds "github.com/spiffe/sri/pkg/server/datastore"
)

//Registration service interface.
type Registration interface {
	CreateEntry(entry *common.RegistrationEntry) (registeredEntryId string, err error)
	FetchEntry(registeredID string) (entry *common.RegistrationEntry, err error)
}

//RegistrationImpl is an implementation of the Registration interface.
type RegistrationImpl struct {
	dataStore ds.DataStore
}

//NewRegistrationImpl creastes a new RegistrationImpl.
func NewRegistrationImpl(dataStore ds.DataStore) *RegistrationImpl {
	return &RegistrationImpl{dataStore: dataStore}
}

//CreateEntry with the DataStore plugin.
func (r *RegistrationImpl) CreateEntry(entry *common.RegistrationEntry) (registeredEntryID string, err error) {
	var response *ds.CreateRegistrationEntryResponse
	if response, err = r.dataStore.CreateRegistrationEntry(&ds.CreateRegistrationEntryRequest{RegisteredEntry: entry}); err != nil {
		return registeredEntryID, err
	}
	return response.RegisteredEntryId, nil
}

//FetchEntry gets a RegisteredEntry based on a registeredID.
func (r *RegistrationImpl) FetchEntry(registeredID string) (*common.RegistrationEntry, error) {
	response, err := r.dataStore.FetchRegistrationEntry(&ds.FetchRegistrationEntryRequest{RegisteredEntryId: registeredID})
	if err != nil {
		return nil, err
	}
	return response.RegisteredEntry, nil
}
