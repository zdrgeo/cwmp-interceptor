package soap

import (
	"encoding/xml"

	"github.com/zdrgeo/cwmp-interceptor/pkg/models"
	"github.com/zdrgeo/cwmp-interceptor/pkg/models/soap/cwmp"
)

const (
	XSD  = "http://www.w3.org/2001/XMLSchema"
	XSI  = "http://www.w3.org/2001/XMLSchema-instance"
	SOAP = "http://schemas.xmlsoap.org/soap/envelope/"
	CWMP = "urn:dslforum-org:cwmp-1-0"
)

const EnvelopeModelDiscriminator models.ModelDiscriminator = iota + 1
const EnvelopeMessageDiscriminator models.MessageDiscriminator = iota + 1

type (
	Header struct {
		XMLName               xml.Name
		ID                    *cwmp.ID                    `xml:"ID,omitempty"`
		HoldRequests          *cwmp.HoldRequests          `xml:"HoldRequests,omitempty"`
		SessionTimeout        *cwmp.SessionTimeout        `xml:"SessionTimeout,omitempty"`
		SupportedCWMPVersions *cwmp.SupportedCWMPVersions `xml:"SupportedCWMPVersions,omitempty"`
		UseCWMPVersion        *cwmp.UseCWMPVersion        `xml:"UseCWMPVersion,omitempty"`
	}

	Fault struct {
		XMLName     xml.Name
		Faultcode   string `xml:"faultcode"`
		Faultstring string `xml:"faultstring"`
		Faultactor  string `xml:"faultactor,omitempty"`
		Detail      string `xml:"detail,omitempty"`
	}

	Body struct {
		XMLName                            xml.Name
		GetRPCMethods                      *cwmp.GetRPCMethods                      `xml:"GetRPCMethods,omitempty"`
		GetRPCMethodsResponse              *cwmp.GetRPCMethodsResponse              `xml:"GetRPCMethodsResponse,omitempty"`
		Inform                             *cwmp.Inform                             `xml:"Inform,omitempty"`
		InformResponse                     *cwmp.InformResponse                     `xml:"InformResponse,omitempty"`
		TransferComplete                   *cwmp.TransferComplete                   `xml:"TransferComplete,omitempty"`
		TransferCompleteResponse           *cwmp.TransferCompleteResponse           `xml:"TransferCompleteResponse,omitempty"`
		AutonomousTransferComplete         *cwmp.AutonomousTransferComplete         `xml:"AutonomousTransferComplete,omitempty"`
		AutonomousTransferCompleteResponse *cwmp.AutonomousTransferCompleteResponse `xml:"AutonomousTransferCompleteResponse,omitempty"`
		GetParameterNames                  *cwmp.GetParameterNames                  `xml:"GetParameterNames,omitempty"`
		GetParameterNamesResponse          *cwmp.GetParameterNamesResponse          `xml:"GetParameterNamesResponse,omitempty"`
		GetParameterValues                 *cwmp.GetParameterValues                 `xml:"GetParameterValues,omitempty"`
		GetParameterValuesResponse         *cwmp.GetParameterValuesResponse         `xml:"GetParameterValuesResponse,omitempty"`
		SetParameterValues                 *cwmp.SetParameterValues                 `xml:"SetParameterValues,omitempty"`
		SetParameterValuesResponse         *cwmp.SetParameterValuesResponse         `xml:"SetParameterValuesResponse,omitempty"`
		GetParameterAttributes             *cwmp.GetParameterAttributes             `xml:"GetParameterAttributes,omitempty"`
		GetParameterAttributesResponse     *cwmp.GetParameterAttributesResponse     `xml:"GetParameterAttributesResponse,omitempty"`
		SetParameterAttributes             *cwmp.SetParameterAttributes             `xml:"SetParameterAttributes,omitempty"`
		SetParameterAttributesResponse     *cwmp.SetParameterAttributesResponse     `xml:"SetParameterAttributesResponse,omitempty"`
		AddObject                          *cwmp.AddObject                          `xml:"AddObject,omitempty"`
		AddObjectResponse                  *cwmp.AddObjectResponse                  `xml:"AddObjectResponse,omitempty"`
		DeleteObject                       *cwmp.DeleteObject                       `xml:"DeleteObject,omitempty"`
		DeleteObjectResponse               *cwmp.DeleteObjectResponse               `xml:"DeleteObjectResponse,omitempty"`
		Reboot                             *cwmp.Reboot                             `xml:"Reboot,omitempty"`
		RebootResponse                     *cwmp.RebootResponse                     `xml:"RebootResponse,omitempty"`
		Download                           *cwmp.Download                           `xml:"Download,omitempty"`
		DownloadResponse                   *cwmp.DownloadResponse                   `xml:"DownloadResponse,omitempty"`
		Fault                              *Fault                                   `xml:"Fault,omitempty"`
	}

	Envelope struct {
		XMLName xml.Name
		XSD     string  `xml:"xmlns:xsd,attr"`
		XSI     string  `xml:"xmlns:xsi,attr"`
		SOAP    string  `xml:"xmlns:soap,attr"`
		CWMP    string  `xml:"xmlns:cwmp,attr"`
		Header  *Header `xml:"Header"`
		Body    *Body   `xml:"Body"`
	}

	EnvelopeModel struct {
		XMLName xml.Name
		Envelope
	}

	EnvelopeMessage struct {
		XMLName xml.Name
		Envelope
	}
)

func (m *EnvelopeModel) ModelDiscriminator() models.ModelDiscriminator {
	return EnvelopeModelDiscriminator
}

func (m *EnvelopeMessage) MessageDiscriminator() models.MessageDiscriminator {
	return EnvelopeMessageDiscriminator
}
