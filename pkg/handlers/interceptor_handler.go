package handlers

import (
	"bytes"
	"encoding/xml"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/zdrgeo/cwmp-interceptor/pkg/models"
	"github.com/zdrgeo/cwmp-interceptor/pkg/models/soap"
	"github.com/zdrgeo/cwmp-interceptor/pkg/services"
)

type InterceptorHandler struct {
	targetURL           *url.URL
	reverseProxy        *httputil.ReverseProxy
	eavesdropperService *services.EavesdropperService
}

func NewInterceptorHandler(targetURL *url.URL, reverseProxy *httputil.ReverseProxy, eavesdropperService *services.EavesdropperService) *InterceptorHandler {
	return &InterceptorHandler{targetURL, reverseProxy, eavesdropperService}
}

func (h *InterceptorHandler) Intercept(writer http.ResponseWriter, request *http.Request) {
	// requestModel, err := readRequestModel(request)

	// if err != nil {
	// 	http.Error(writer, "Bad Request", http.StatusBadRequest)

	// 	return
	// }

	// if err != h.eavesdropperService.Eavesdrop(request.Context(), requestModel) {
	// 	http.Error(writer, "Internal Server Error", http.StatusInternalServerError)

	// 	return
	// }

	requestMessage, err := peekRequestMessage(request)

	if err != nil {
		http.Error(writer, "Bad Request", http.StatusBadRequest)

		return
	}

	if err != h.eavesdropperService.Eavesdrop(request.Context(), requestMessage) {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)

		return
	}

	request.Host = h.targetURL.Host

	h.reverseProxy.ServeHTTP(writer, request)
}

func peekRequestModel(request *http.Request) (models.Model, error) {
	// xmlRequest, err := io.ReadAll(request.Body) // io.ReadAll(io.LimitReader(request.Body, readLimit))

	// if err != nil {
	// 	return nil, err
	// }

	// request.Body = io.NopCloser(bytes.NewBuffer(xmlRequest))

	buffer := &bytes.Buffer{}

	xmlRequest, err := io.ReadAll(io.TeeReader(request.Body, buffer))

	if err != nil {
		return nil, err
	}

	request.Body = io.NopCloser(buffer)

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

func peekRequestMessage(request *http.Request) (models.Message, error) {
	// xmlRequest, err := io.ReadAll(request.Body) // io.ReadAll(io.LimitReader(request.Body, readLimit))

	// if err != nil {
	// 	return nil, err
	// }

	// request.Body = io.NopCloser(bytes.NewBuffer(xmlRequest))

	buffer := &bytes.Buffer{}

	xmlRequest, err := io.ReadAll(io.TeeReader(request.Body, buffer))

	if err != nil {
		return nil, err
	}

	request.Body = io.NopCloser(buffer)

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
