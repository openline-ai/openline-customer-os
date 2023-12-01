package organization

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	ai "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-ai/service"
	commonEntity "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/repository/postgres/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/config"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/constants"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/repository"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type WebScraper interface {
	Scrape(domainOrWebsite, tenant, organizationId string, directScrape bool) (*WebscrapeResponseV1, error)
}

type DomainScraperV1 struct {
	log          logger.Logger
	cfg          *config.Config
	repositories *repository.Repositories
	aiModel      ai.AiModel
}

func NewDomainScraper(log logger.Logger, cfg *config.Config, repositories *repository.Repositories, aiModel ai.AiModel) WebScraper {
	return &DomainScraperV1{
		log:          log,
		cfg:          cfg,
		repositories: repositories,
		aiModel:      aiModel,
	}
}

func (ds *DomainScraperV1) Scrape(domainOrWebsite, tenant, organizationId string, directScrape bool) (*WebscrapeResponseV1, error) {
	domainUrl := strings.TrimSpace(domainOrWebsite)
	if !strings.HasPrefix(domainUrl, "http") && !strings.HasPrefix(domainUrl, "www") {
		domainUrl = fmt.Sprintf("https://%s", domainUrl)
	}
	jsonStruct := jsonStructure()

	html, err := ds.getHtml(domainUrl, directScrape)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to getHtml domain: %s", domainUrl))
	}
	socialLinks, err := ds.extractSocialLinks(html)
	if err != nil {
		return nil, errors.Wrap(err, "failed to extract social links")
	}

	text, err := ds.extractRelevantText(html)
	if err != nil {
		return nil, errors.Wrap(err, "failed to extract relevant text")
	}

	companyAnalysis, err := ds.runCompanyPrompt(text, tenant, organizationId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to run company prompt")
	}

	r, err := ds.runDataPrompt(companyAnalysis, &domainUrl, socialLinks, jsonStruct, tenant, organizationId)

	if err != nil {
		return nil, errors.Wrap(err, "failed to run data prompt")
	}
	return r, nil
}

func (ds *DomainScraperV1) getHtml(domainUrl string, directGet bool) (*string, error) {
	var response *http.Response
	var err error
	if directGet {
		response, err = ds.getRequest(domainUrl)
		if err != nil {
			return nil, errors.Wrap(err, "failed to execute request")
		}
	} else {
		response, err = ds.proxyGetRequest(ds.cfg.Services.ScrapingBeeApiKey, domainUrl)
		if err != nil {
			return nil, errors.Wrap(err, "failed to execute request")
		}
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("StatusCode: %s", response.Status)
	}

	// Read Response Body
	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	s := string(respBody)

	return &s, nil
}

func (ds *DomainScraperV1) runCompanyPrompt(text *string, tenant, organizationId string) (*string, error) {
	p := strings.ReplaceAll(ds.cfg.Services.OpenAi.ScrapeCompanyPrompt, "{{jsonschema}}", ds.cfg.Services.PromptJsonSchema)
	prompt := strings.ReplaceAll(p, "{{text}}", *text)
	ctx := context.TODO()

	promptLog := commonEntity.AiPromptLog{
		CreatedAt:      utils.Now(),
		AppSource:      constants.AppSourceEventProcessingPlatform,
		Provider:       constants.OpenAI,
		Model:          "gpt-3.5-turbo",
		PromptType:     constants.PromptType_WebscrapeCompanyPrompt,
		Tenant:         &tenant,
		NodeId:         &organizationId,
		NodeLabel:      utils.StringPtr(constants.NodeLabel_Organization),
		PromptTemplate: &ds.cfg.Services.OpenAi.ScrapeCompanyPrompt,
		Prompt:         prompt,
	}

	// ignore error from storing prompt log, since it's not critical
	promptStoreLogId, _ := ds.repositories.CommonRepositories.AiPromptLogRepository.Store(promptLog)

	aiResult, err := ds.aiModel.Inference(ctx, prompt)
	if err != nil {
		_ = ds.repositories.CommonRepositories.AiPromptLogRepository.UpdateError(promptStoreLogId, err.Error())
		return nil, errors.Wrap(err, "unable to get openai result")
	}
	_ = ds.repositories.CommonRepositories.AiPromptLogRepository.UpdateResponse(promptStoreLogId, aiResult)

	return &aiResult, nil
}

func (ds *DomainScraperV1) getRequest(domainUrl string) (*http.Response, error) {
	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", domainUrl, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create direct request")
	}
	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
		return nil, errors.Wrap(parseFormErr, "failed to parse form")
	}

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch request")
	}
	return resp, nil // Return the response
}

func (ds *DomainScraperV1) proxyGetRequest(apiKey string, domainUrl string) (*http.Response, error) {
	// Create client
	client := &http.Client{}

	urlEscaped := url.QueryEscape(domainUrl) // Encoding the URL
	// Create request
	req, err := http.NewRequest("GET", "https://app.scrapingbee.com/api/v1/?api_key="+apiKey+"&url="+urlEscaped, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create scrappingbee request")
	}
	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
		return nil, errors.Wrap(parseFormErr, "failed to parse form")
	}

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch request")
	}
	return resp, nil // Return the response
}

func (ds *DomainScraperV1) extractRelevantText(html *string) (*string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(*html))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create document from reader")
	}

	converter := md.NewConverter("", true, nil)
	markdown := converter.Convert(doc.Selection)
	ds.log.Printf("text: %s", markdown)
	return &markdown, nil
}

func (ds *DomainScraperV1) extractSocialLinks(html *string) (*string, error) {
	socialSites := map[string]string{
		"linkedin":  "linkedin.com",
		"twitter":   "twitter.com",
		"instagram": "instagram.com",
		"facebook":  "facebook.com",
		"youtube":   "youtube.com",
		"github":    "github.com",
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(*html))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create document from reader")
	}
	links := make(map[string]string)

	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
			return
		}

		for name, site := range socialSites {
			if strings.Contains(href, site) {
				links[name] = href
				break
			}
		}
	})

	// Marshal the map into a JSON string
	jsonBytes, err := json.Marshal(links)
	if err != nil {
		ds.log.Printf("Error marshalling social links to JSON: %s", err)
	}

	// Convert JSON bytes to a string
	s := string(jsonBytes)
	return &s, nil
}

func (ds *DomainScraperV1) runDataPrompt(analysis, domainUrl, socials, jsonStructure *string, tenant, organizationId string) (*WebscrapeResponseV1, error) {

	replacements := map[string]string{
		"{{ANALYSIS}}":       *analysis,
		"{{DOMAIN_URL}}":     *domainUrl,
		"{{SOCIALS}}":        *socials,
		"{{JSON_STRUCTURE}}": *jsonStructure,
	}

	prompt := ds.cfg.Services.OpenAi.ScrapeDataPrompt
	for k, v := range replacements {
		prompt = strings.ReplaceAll(prompt, k, v)
	}

	promptLog := commonEntity.AiPromptLog{
		CreatedAt:      utils.Now(),
		AppSource:      constants.AppSourceEventProcessingPlatform,
		Provider:       constants.OpenAI,
		Model:          "gpt-3.5-turbo",
		PromptType:     constants.PromptType_WebscrapeExtractCompanyData,
		Tenant:         &tenant,
		NodeId:         &organizationId,
		NodeLabel:      utils.StringPtr(constants.NodeLabel_Organization),
		PromptTemplate: &ds.cfg.Services.OpenAi.ScrapeDataPrompt,
		Prompt:         prompt,
	}
	promptStoreLogId, _ := ds.repositories.CommonRepositories.AiPromptLogRepository.Store(promptLog)

	cleaned, err := ds.aiModel.Inference(context.TODO(), prompt)

	if err != nil {
		_ = ds.repositories.CommonRepositories.AiPromptLogRepository.UpdateError(promptStoreLogId, err.Error())
		return nil, err
	}
	_ = ds.repositories.CommonRepositories.AiPromptLogRepository.UpdateResponse(promptStoreLogId, cleaned)
	ds.log.Printf("scrapeResponse: %s", cleaned)
	scrapeResponse := WebscrapeResponseV1{}
	err = json.Unmarshal([]byte(cleaned), &scrapeResponse)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal response")
	}
	//Leaving the fields empty and not with AI generic error messages
	if strings.Contains(scrapeResponse.ValueProposition, "Unable error") {
		ds.log.Printf("Error to obtain value for ValueProposition: %s", scrapeResponse.ValueProposition)
		scrapeResponse.ValueProposition = ""
	}
	if strings.Contains(scrapeResponse.TargetAudience, "Unable error") {
		ds.log.Printf("Error to obtain value for TargetAudience: %s", scrapeResponse.TargetAudience)
		scrapeResponse.TargetAudience = ""
	}

	return &scrapeResponse, nil
}

func jsonStructure() *string {
	data := WebscrapeResponseV1{
		CompanyName:      "...",
		Website:          "...",
		Market:           "...",
		Industry:         "...",
		IndustryGroup:    "...",
		SubIndustry:      "...",
		TargetAudience:   "...",
		ValueProposition: "...",
		Github:           "...",
		Linkedin:         "...",
		Twitter:          "...",
		Youtube:          "...",
		Instagram:        "...",
		Facebook:         "...",
	}

	jsonStructure, _ := json.Marshal(data)

	s := string(jsonStructure)
	return &s
}
