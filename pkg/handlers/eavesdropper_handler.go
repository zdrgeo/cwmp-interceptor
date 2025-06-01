package handlers

import (
	"encoding/xml"
	"io"
	"net/http"

	"github.com/zdrgeo/cwmp-interceptor/pkg/models"
	"github.com/zdrgeo/cwmp-interceptor/pkg/models/soap"
	"github.com/zdrgeo/cwmp-interceptor/pkg/services"
)

type EavesdropperHandler struct {
	eavesdropperService *services.EavesdropperService
}

func NewEavesdropperHandler(eavesdropperService *services.EavesdropperService) *EavesdropperHandler {
	return &EavesdropperHandler{eavesdropperService}
}

func (h *EavesdropperHandler) Eavesdrop(writer http.ResponseWriter, request *http.Request) {
	// requestModel, err := readRequestModel(request)

	// if err != nil {
	// 	http.Error(writer, "Bad Request", http.StatusBadRequest)

	// 	return
	// }

	// if err != h.eavesdropperService.Eavesdrop(request.Context(), requestModel) {
	// 	http.Error(writer, "Internal Server Error", http.StatusInternalServerError)

	// 	return
	// }

	requestMessage, err := readRequestMessage(request)

	if err != nil {
		http.Error(writer, "Bad Request", http.StatusBadRequest)

		return
	}

	if err != h.eavesdropperService.Eavesdrop(request.Context(), requestMessage) {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)

		return
	}
}

func readRequestModel(request *http.Request) (models.Model, error) {
	xmlRequest, err := io.ReadAll(request.Body) // io.ReadAll(io.LimitReader(request.Body, readLimit))

	if err != nil {
		return nil, err
	}

	envelopeModel := &soap.EnvelopeModel{
		XMLName:  xml.Name{Local: "soap:Envelope"},
		Envelope: soap.Envelope{XMLName: xml.Name{Local: "soap:Envelope"}, XSD: soap.XSD, XSI: soap.XSI, SOAP: soap.SOAP, CWMP: soap.CWMP},
	}

	if err := xml.Unmarshal([]byte(xmlRequest), &envelopeModel); err != nil {
		if err == io.EOF {
			return &models.NoneModel{}, nil
		} else {
			return nil, err
		}
	}

	return envelopeModel, nil
}

func readRequestMessage(request *http.Request) (models.Message, error) {
	xmlRequest, err := io.ReadAll(request.Body) // io.ReadAll(io.LimitReader(request.Body, readLimit))

	if err != nil {
		return nil, err
	}

	envelopeMessage := &soap.EnvelopeMessage{
		XMLName:  xml.Name{Local: "soap:Envelope"},
		Envelope: soap.Envelope{XMLName: xml.Name{Local: "soap:Envelope"}, XSD: soap.XSD, XSI: soap.XSI, SOAP: soap.SOAP, CWMP: soap.CWMP},
	}

	if err := xml.Unmarshal([]byte(xmlRequest), &envelopeMessage); err != nil {
		if err == io.EOF {
			return &models.NoneMessage{}, nil
		} else {
			return nil, err
		}
	}

	return envelopeMessage, nil
}
