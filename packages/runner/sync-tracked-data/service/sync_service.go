package service

import (
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-tracked-data/entity"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-tracked-data/repository"
	"log"
	"regexp"
	"strings"
	"sync"
)

type SyncService interface {
	Sync(runId string, bucketSize int) int
}

type syncService struct {
	repositories *repository.Repositories
	services     *Services
}

type tenantVisitor struct {
	tenant    string
	visitorId string
}

func NewSyncService(repositories *repository.Repositories, services *Services) SyncService {
	return &syncService{
		repositories: repositories,
		services:     services,
	}
}

func (s *syncService) Sync(runId string, bucketSize int) int {
	pageViewsToSync, err := s.repositories.PageViewRepository.GetPageViewsForSync(bucketSize)
	if err != nil {
		log.Printf("ERROR run id: %s failed to sync page views. error fetching page views: %v", runId, err.Error())
	}

	if len(pageViewsToSync) == 0 {
		return len(pageViewsToSync)
	}

	contactIds, err := s.prepareContactIds(pageViewsToSync)
	if err != nil {
		return 0
	}

	var wg sync.WaitGroup
	wg.Add(len(pageViewsToSync))

	for _, v := range pageViewsToSync {
		go s.syncPageView(&wg, runId, contactIds, v)
	}
	wg.Wait()

	return len(pageViewsToSync)
}

func (s *syncService) prepareContactIds(pageViews entity.PageViews) (map[tenantVisitor]string, error) {

	var contactIds = map[tenantVisitor]string{}

	for _, v := range pageViews {
		email := v.VisitorID.String
		tenantVisitor := tenantVisitor{
			tenant:    v.Tenant,
			visitorId: v.VisitorID.String,
		}
		if _, ok := contactIds[tenantVisitor]; !ok {
			firstName, lastName := s.prepareFirstAndLastNames(email)
			id, err := s.repositories.ContactRepository.GetOrCreateContactId(v.Tenant, email, firstName, lastName)
			if err != nil {
				return nil, err
			}
			contactIds[tenantVisitor] = id
		}
	}

	return contactIds, nil
}

func (s *syncService) syncPageView(wg *sync.WaitGroup, runId string, contactIds map[tenantVisitor]string, pv entity.PageView) string {
	defer wg.Done()

	contactId := contactIds[tenantVisitor{
		tenant:    pv.Tenant,
		visitorId: pv.VisitorID.String,
	}]
	if err := s.repositories.ActionRepository.CreatePageViewAction(contactId, pv); err != nil {
		log.Printf("ERROR run id: %s failed to create action item for page view %s error: %v", runId, pv.ID, err.Error())
	} else {
		if err = s.repositories.PageViewRepository.MarkSynced(pv, contactId); err != nil {
			log.Printf("ERROR run id: %s failed to mark as sycned page view %s error: %v", runId, pv.ID, err.Error())
		} else {
			log.Printf("run id: %s synced page view %s", runId, pv.ID)
		}
	}
	return pv.ID
}

func (s *syncService) parseEmail(email string) (string, string) {
	re := regexp.MustCompile("^\"{0,1}([^\"]*)\"{0,1}[ ]*<(.*)>$")
	matches := re.FindStringSubmatch(strings.Trim(email, " "))
	if matches != nil {
		return strings.Trim(matches[1], " "), matches[2]
	}
	return "", email
}

func (s *syncService) prepareFirstAndLastNames(email string) (string, string) {
	displayName, _ := s.parseEmail(email)
	firstName, lastName := "", ""
	if displayName != "" {
		parts := strings.SplitN(displayName, " ", 2)
		firstName = parts[0]
		if len(parts) > 1 {
			lastName = parts[1]
		}
	}
	return firstName, lastName
}
