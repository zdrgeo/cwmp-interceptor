package services

import (
	"context"
	"errors"
	"time"

	bulkdatacollectorservices "github.com/zdrgeo/bulk-data-collector/pkg/services"
	"github.com/zdrgeo/cwmp-interceptor/pkg/models"
	"github.com/zdrgeo/cwmp-interceptor/pkg/models/soap"
)

var (
	ErrInvalidNoneMessage     = errors.New("invalid None message")
	ErrInvalidEnvelopeMessage = errors.New("invalid Envelope message")
)

type EavesdropperServiceOptions struct{}

type EavesdropperService struct {
	collectorService bulkdatacollectorservices.CollectorService
	options          *EavesdropperServiceOptions
}

func NewEavesdropperService(collectorService bulkdatacollectorservices.CollectorService, option *EavesdropperServiceOptions) *EavesdropperService {
	return &EavesdropperService{collectorService: collectorService, options: option}
}

func (s *EavesdropperService) Eavesdrop(ctx context.Context, requestMessage models.Message) error {
	actionName, err := getActionName(requestMessage)

	if err != nil {
		return err
	}

	switch actionName {
	case "Inform":
		return s.eavesdropInform(ctx, requestMessage)
	case "GetParameterValuesResponse":
		return s.eavesdropGetParameterValuesResponse(ctx, requestMessage)
	}

	return nil
}

func getActionName(message models.Message) (string, error) {
	var actionName string

	switch message.MessageDiscriminator() {
	case soap.EnvelopeMessageDiscriminator:
		envelopeMessage, ok := message.(*soap.EnvelopeMessage)

		if !ok {
			return "", ErrInvalidEnvelopeMessage
		}

		if envelopeMessage.Envelope.Body.GetRPCMethods != nil {
			actionName = "GetRPCMethods"
		}
		if envelopeMessage.Envelope.Body.Inform != nil {
			actionName = "Inform"
		}
		if envelopeMessage.Envelope.Body.TransferComplete != nil {
			actionName = "TransferComplete"
		}
		if envelopeMessage.Envelope.Body.AutonomousTransferComplete != nil {
			actionName = "AutonomousTransferComplete"
		}
		if envelopeMessage.Envelope.Body.GetRPCMethodsResponse != nil {
			actionName = "GetRPCMethodsResponse"
		}
		if envelopeMessage.Envelope.Body.GetParameterNamesResponse != nil {
			actionName = "GetParameterNamesResponse"
		}
		if envelopeMessage.Envelope.Body.GetParameterValuesResponse != nil {
			actionName = "GetParameterValuesResponse"
		}
		if envelopeMessage.Envelope.Body.SetParameterValuesResponse != nil {
			actionName = "SetParameterValuesResponse"
		}
		if envelopeMessage.Envelope.Body.GetParameterAttributesResponse != nil {
			actionName = "GetParameterAttributesResponse"
		}
		if envelopeMessage.Envelope.Body.SetParameterAttributesResponse != nil {
			actionName = "SetParameterAttributesResponse"
		}
		if envelopeMessage.Envelope.Body.AddObjectResponse != nil {
			actionName = "AddObjectResponse"
		}
		if envelopeMessage.Envelope.Body.DeleteObjectResponse != nil {
			actionName = "DeleteObjectResponse"
		}
		if envelopeMessage.Envelope.Body.RebootResponse != nil {
			actionName = "RebootResponse"
		}
		if envelopeMessage.Envelope.Body.DownloadResponse != nil {
			actionName = "DownloadResponse"
		}
		if envelopeMessage.Envelope.Body.Fault != nil {
			actionName = "Fault"
		}
	case models.NoneMessageDiscriminator:
		actionName = "None"
	}

	return actionName, nil
}

func (s *EavesdropperService) eavesdropInform(ctx context.Context, message models.Message) error {
	requestEnvelopeMessage, ok := message.(*soap.EnvelopeMessage)

	if !ok {
		return ErrInvalidEnvelopeMessage
	}

	requestBody := requestEnvelopeMessage.Envelope.Body

	inform := requestBody.Inform

	collectionTime, err := time.Parse(time.RFC3339, inform.CurrentTime)

	if err != nil {
		return err
	}

	reportModel := &bulkdatacollectorservices.ReportModel{
		CollectionTime: collectionTime,
		Parameters:     make(map[string]any, len(inform.ParameterList.ParameterValueStruct)),
	}

	dataModel := &bulkdatacollectorservices.DataModel{
		Reports: []*bulkdatacollectorservices.ReportModel{},
	}

	dataModel.Reports = append(dataModel.Reports, reportModel)

	for _, parameterValueStruct := range inform.ParameterList.ParameterValueStruct {
		reportModel.Parameters[parameterValueStruct.Name] = parameterValueStruct.Value
	}

	if err := s.collectorService.Collect(ctx, inform.DeviceId.OUI, inform.DeviceId.ProductClass, inform.DeviceId.SerialNumber, dataModel); err != nil {
		return err
	}

	return nil
}

func (s *EavesdropperService) eavesdropGetParameterValuesResponse(ctx context.Context, message models.Message) error {
	requestEnvelopeMessage, ok := message.(*soap.EnvelopeMessage)

	if !ok {
		return ErrInvalidEnvelopeMessage
	}

	requestBody := requestEnvelopeMessage.Envelope.Body

	getParameterValuesResponse := requestBody.GetParameterValuesResponse

	collectionTime, err := time.Parse(time.RFC3339, "<inform.CurrentTime>")

	if err != nil {
		return err
	}

	reportModel := &bulkdatacollectorservices.ReportModel{
		CollectionTime: collectionTime,
		Parameters:     make(map[string]any, len(getParameterValuesResponse.ParameterList.ParameterValueStruct)),
	}

	dataModel := &bulkdatacollectorservices.DataModel{
		Reports: []*bulkdatacollectorservices.ReportModel{},
	}

	dataModel.Reports = append(dataModel.Reports, reportModel)

	for _, parameterValueStruct := range getParameterValuesResponse.ParameterList.ParameterValueStruct {
		reportModel.Parameters[parameterValueStruct.Name] = parameterValueStruct.Value
	}

	if err := s.collectorService.Collect(ctx, "<inform.DeviceId.OUI>", "<inform.DeviceId.ProductClass>", "<inform.DeviceId.SerialNumber>", dataModel); err != nil {
		return err
	}

	return nil
}
