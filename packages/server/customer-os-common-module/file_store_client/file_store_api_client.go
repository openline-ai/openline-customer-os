package fsc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

type FileStoreApiService interface {
	UploadSingleMultipartFile(tenantName string, multipartFileHeader *multipart.FileHeader) (*FileDTO, error)
	UploadSingleFileBytes(tenantName string, fileBytes []byte) (*FileDTO, error)
}

type fileStoreApiService struct {
	conf *FileStoreApiConfig
}

func (fsas *fileStoreApiService) UploadSingleMultipartFile(tenantName string, multipartFileHeader *multipart.FileHeader) (*FileDTO, error) {
	file, err := multipartFileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("UploadSingleMultipartFile: failed to open multipart file: %w", err)
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("UploadSingleMultipartFile: failed to read multipart file: %w", err)
	}

	return sendRequest(fsas.conf, tenantName, fileBytes)
}

func (fsas *fileStoreApiService) UploadSingleFileBytes(tenantName string, fileBytes []byte) (*FileDTO, error) {
	return sendRequest(fsas.conf, tenantName, fileBytes)
}

func sendRequest(conf *FileStoreApiConfig, tenantName string, fileBytes []byte) (*FileDTO, error) {
	// Create a new buffer to store the request body
	var requestBody bytes.Buffer

	// Create a new multipart writer
	writer := multipart.NewWriter(&requestBody)

	// Create a form file field for the file
	fileWriter, err := writer.CreateFormFile("file", "filename.txt")
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return nil, err
	}

	// Copy the file content (bytes) to the form file field
	_, err = fileWriter.Write(fileBytes)
	if err != nil {
		fmt.Println("Error writing file content:", err)
		return nil, err
	}

	// Close the multipart writer to finalize the request body
	writer.Close()

	url := fmt.Sprintf("%s/file", conf.ApiPath)
	log.Printf("UploadSingleMultipartFile: url: %s", url)
	req, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		return nil, fmt.Errorf("UploadSingleMultipartFile: failed to create new request: %w", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("accept", "application/json")
	req.Header.Add("X-Openline-API-KEY", conf.ApiKey)
	req.Header.Add("X-Openline-Tenant", tenantName)

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("UploadSingleMultipartFile: failed to perform request: %w", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		var fileResponse FileDTO
		if err := json.NewDecoder(resp.Body).Decode(&fileResponse); err != nil {
			return nil, fmt.Errorf("UploadSingleMultipartFile: failed to decode response: %w", err)
		}
		return &fileResponse, nil
	} else {
		var responseBody bytes.Buffer
		_, err = io.Copy(&responseBody, resp.Body)
		if err != nil {
			fmt.Println("Error reading response:", err)
			return nil, err
		}

		err = fmt.Errorf("Got error from File Store API: Status: %d Response: %s", resp.StatusCode, responseBody.String())
		return nil, err
	}
}

func NewFileStoreApiService(conf *FileStoreApiConfig) *fileStoreApiService {
	return &fileStoreApiService{
		conf: conf,
	}
}