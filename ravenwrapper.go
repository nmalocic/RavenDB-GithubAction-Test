package ravendbtest

import (
	"fmt"
	ravendb "github.com/ravendb/ravendb-go-client"
)

var (
	dbName        = "testdapr"
	serverNodeURL = "http://127.0.0.1:8081"
	enableTTL     = true
	ttlFrequency  = 2
)

type RavenDB_Wrapper struct {
	documentStore *ravendb.DocumentStore
}

func (r *RavenDB_Wrapper) Init() (err error) {
	store, err := r.getDocumentStore(dbName)
	if err != nil {
		return err
	}
	r.documentStore = store

	if err := r.initTTL(store); err != nil {
		return fmt.Errorf("failed to initialize TTL: %w", err)
	}

	return nil
}

func (r *RavenDB_Wrapper) initTTL(store *ravendb.DocumentStore) error {
	ttl := int64(ttlFrequency)
	configurationExpiration := ravendb.ExpirationConfiguration{
		Disabled:             false,
		DeleteFrequencyInSec: &ttl,
	}
	operation, err := ravendb.NewConfigureExpirationOperationWithConfiguration(&configurationExpiration)
	if err != nil {
		return fmt.Errorf("failed to create expiration operation: %w", err)
	}
	if err := store.Maintenance().Send(operation); err != nil {
		return fmt.Errorf("failed to send expiration operation: %w", err)
	}
	return nil
}

func (r *RavenDB_Wrapper) openSession(databaseName string) (*ravendb.DocumentStore, *ravendb.DocumentSession, error) {
	store, err := r.getDocumentStore(databaseName)
	if err != nil {
		return nil, nil, fmt.Errorf("getDocumentStore() failed: %w", err)
	}

	session, err := store.OpenSession(databaseName)
	if err != nil {
		return nil, nil, fmt.Errorf("store.OpenSession() failed: %w", err)
	}
	return store, session, nil
}

func (r *RavenDB_Wrapper) getDocumentStore(databaseName string) (*ravendb.DocumentStore, error) {
	serverNodes := []string{serverNodeURL}
	store := ravendb.NewDocumentStore(serverNodes, databaseName)
	if err := store.Initialize(); err != nil {
		return nil, err
	}
	return store, nil
}
