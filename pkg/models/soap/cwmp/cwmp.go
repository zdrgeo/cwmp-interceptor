package cwmp

import (
	"encoding/xml"

	"github.com/zdrgeo/cwmp-interceptor/pkg/models"

	"time"
)

const (
	ACS_GetRPCMethods              = "GetRPCMethods"
	ACS_Inform                     = "Inform"
	ACS_TransferComplete           = "TransferComplete"
	ACS_AutonomousTransferComplete = "AutonomousTransferComplete"

	CPE_GetRPCMethods          = "GetRPCMethods"
	CPE_GetParameterNames      = "GetParameterNames"
	CPE_GetParameterValues     = "GetParameterValues"
	CPE_SetParameterValues     = "SetParameterValues"
	CPE_GetParameterAttributes = "GetParameterAttributes"
	CPE_SetParameterAttributes = "SetParameterAttributes"
	CPE_AddObject              = "AddObject"
	CPE_DeleteObject           = "DeleteObject"
	CPE_Reboot                 = "Reboot"
	CPE_Download               = "Download"
)

const (
	ACS_Inform_EventCode_0_Bootstrap                            = "0 BOOTSTRAP"
	ACS_Inform_EventCode_1_Boot                                 = "1 BOOT"
	ACS_Inform_EventCode_2_Periodic                             = "2 PERIODIC"
	ACS_Inform_EventCode_3_Scheduled                            = "3 SCHEDULED"
	ACS_Inform_EventCode_4_Value_Changed                        = "4 VALUE CHANGED"
	ACS_Inform_EventCode_5_Kicked                               = "5 KICKED"
	ACS_Inform_EventCode_6_Connection_Request                   = "6 CONNECTION REQUEST"
	ACS_Inform_EventCode_7_Transfer_Complete                    = "7 TRANSFER COMPLETE"
	ACS_Inform_EventCode_8_Diagnostics_Complete                 = "8 DIAGNOSTICS COMPLETE"
	ACS_Inform_EventCode_9_Request_Download                     = "9 REQUEST DOWNLOAD"
	ACS_Inform_EventCode_10_Autonomous_Transfer_Complete        = "10 AUTONOMOUS TRANSFER COMPLETE"
	ACS_Inform_EventCode_11_DU_State_Change_Complete            = "11 DU STATE CHANGE COMPLETE"
	ACS_Inform_EventCode_12_Autonomous_DU_State_Change_Complete = "12 AUTONOMOUS DU STATE CHANGE COMPLETE"
	ACS_Inform_EventCode_13_Wakeup                              = "13 WAKEUP"
	ACS_Inform_EventCode_14_Heartbeat                           = "14 HEARTBEAT"

	ACS_Inform_EventCode_M_Reboot           = "M Reboot"
	ACS_Inform_EventCode_M_ScheduleInform   = "M ScheduleInform"
	ACS_Inform_EventCode_M_Download         = "M Download"
	ACS_Inform_EventCode_M_ScheduleDownload = "M ScheduleDownload"
	ACS_Inform_EventCode_M_Upload           = "M Upload"
	ACS_Inform_EventCode_M_ChangeDUState    = "M ChangeDUState"
)

const (
	GetRPCMethodsModelDiscriminator models.ModelDiscriminator = iota + 2
	GetRPCMethodsResponseModelDiscriminator
	InformModelDiscriminator
	InformResponseModelDiscriminator
	TransferCompleteModelDiscriminator
	TransferCompleteResponseModelDiscriminator
	AutonomousTransferCompleteModelDiscriminator
	AutonomousTransferCompleteResponseModelDiscriminator
	GetParameterNamesModelDiscriminator
	GetParameterNamesResponseModelDiscriminator
	GetParameterValuesModelDiscriminator
	GetParameterValuesResponseModelDiscriminator
	SetParameterValuesModelDiscriminator
	SetParameterValuesResponseModelDiscriminator
	GetParameterAttributesModelDiscriminator
	GetParameterAttributesResponseModelDiscriminator
	SetParameterAttributesModelDiscriminator
	SetParameterAttributesResponseModelDiscriminator
	AddObjectModelDiscriminator
	AddObjectResponseModelDiscriminator
	DeleteObjectModelDiscriminator
	DeleteObjectResponseModelDiscriminator
	RebootModelDiscriminator
	RebootResponseModelDiscriminator
	DownloadModelDiscriminator
	DownloadResponseModelDiscriminator
)

const (
	GetRPCMethodsMessageDiscriminator models.MessageDiscriminator = iota + 2
	GetRPCMethodsResponseMessageDiscriminator
	InformMessageDiscriminator
	InformResponseMessageDiscriminator
	TransferCompleteMessageDiscriminator
	TransferCompleteResponseMessageDiscriminator
	AutonomousTransferCompleteMessageDiscriminator
	AutonomousTransferCompleteResponseMessageDiscriminator
	GetParameterNamesMessageDiscriminator
	GetParameterNamesResponseMessageDiscriminator
	GetParameterValuesMessageDiscriminator
	GetParameterValuesResponseMessageDiscriminator
	SetParameterValuesMessageDiscriminator
	SetParameterValuesResponseMessageDiscriminator
	GetParameterAttributesMessageDiscriminator
	GetParameterAttributesResponseMessageDiscriminator
	SetParameterAttributesMessageDiscriminator
	SetParameterAttributesResponseMessageDiscriminator
	AddObjectMessageDiscriminator
	AddObjectResponseMessageDiscriminator
	DeleteObjectMessageDiscriminator
	DeleteObjectResponseMessageDiscriminator
	RebootMessageDiscriminator
	RebootResponseMessageDiscriminator
	DownloadMessageDiscriminator
	DownloadResponseMessageDiscriminator
)

type (
	SetParameterValuesFault struct {
		ParameterName string `xml:"ParameterName"`
		FaultCode     int    `xml:"FaultCode"`
		FaultString   string `xml:"FaultString"`
	}

	FaultStruct struct {
		FaultCode               int                      `xml:"FaultCode"`
		FaultString             string                   `xml:"FaultString"`
		SetParameterValuesFault *SetParameterValuesFault `xml:"SetParameterValuesFault"`
	}

	ID struct {
		XMLName        xml.Name `xml:"ID"`
		MustUnderstand int      `xml:"mustUnderstand,attr"`
		Value          string   `xml:",chardata"`
	}

	HoldRequests struct {
		MustUnderstand int  `xml:"mustUnderstand,attr"`
		Value          bool `xml:",chardata"`
	}

	SessionTimeout struct {
		MustUnderstand int `xml:"mustUnderstand,attr"`
		Value          int `xml:",chardata"`
	}

	SupportedCWMPVersions struct {
		MustUnderstand int    `xml:"mustUnderstand,attr"`
		Value          string `xml:",chardata"`
	}

	UseCWMPVersion struct {
		MustUnderstand int    `xml:"mustUnderstand,attr"`
		Value          string `xml:",chardata"`
	}

	DeviceIdStruct struct {
		Manufacturer string `xml:"Manufacturer"`
		OUI          string `xml:"OUI"`
		ProductClass string `xml:"ProductClass"`
		SerialNumber string `xml:"SerialNumber"`
	}

	EventStruct struct {
		EventCode  string `xml:"EventCode"`
		CommandKey string `xml:"CommandKey"`
	}

	EventList struct {
		EventStruct []EventStruct `xml:"EventStruct"`
	}

	ParameterInfoStruct struct {
		Name     string `xml:"Name"`
		Writable bool   `xml:"Writable"`
	}

	ParameterValueStruct struct {
		Name  string `xml:"Name"`
		Value string `xml:"Value"`
	}

	ParameterAttributeStruct struct {
		Name         string   `xml:"Name"`
		Notification int      `xml:"Notification"`
		AccessList   []string `xml:"AccessList>string"`
	}

	SetParameterAttributesStruct struct {
		Name               string   `xml:"Name"`
		NotificationChange bool     `xml:"NotificationChange"`
		Notification       int      `xml:"Notification"`
		AccessListChange   bool     `xml:"AccessListChange"`
		AccessList         []string `xml:"AccessList>string"`
	}

	ParameterInfoList struct {
		ParameterInfoStruct []*ParameterInfoStruct `xml:"ParameterInfoStruct"`
	}

	ParameterValueList struct {
		ParameterValueStruct []*ParameterValueStruct `xml:"ParameterValueStruct"`
	}

	ParameterAttributeList struct {
		ParameterAttributeStruct []*ParameterAttributeStruct `xml:"ParameterAttributeStruct"`
	}

	SetParameterAttributesList struct {
		SetParameterAttributesStruct []*SetParameterAttributesStruct `xml:"SetParameterAttributesStruct"`
	}

	// Models

	GetRPCMethods struct {
		XMLName xml.Name
	}

	GetRPCMethodsResponse struct {
		XMLName    xml.Name
		MethodList []string `xml:"MethodList>string"`
	}

	Inform struct {
		XMLName       xml.Name
		DeviceId      *DeviceIdStruct     `xml:"DeviceId"`
		Event         *EventList          `xml:"Event"`
		MaxEnvelopes  uint                `xml:"MaxEnvelopes"`
		CurrentTime   string              `xml:"CurrentTime"`
		RetryCount    uint                `xml:"RetryCount"`
		ParameterList *ParameterValueList `xml:"ParameterList"`
	}

	InformResponse struct {
		XMLName      xml.Name
		MaxEnvelopes uint `xml:"MaxEnvelopes"`
	}

	TransferComplete struct {
		XMLName      xml.Name
		CommandKey   string       `xml:"CommandKey"`
		FaultStruct  *FaultStruct `xml:"FaultStruct"`
		StartTime    time.Time    `xml:"StartTime"`
		CompleteTime time.Time    `xml:"CompleteTime"`
	}

	TransferCompleteResponse struct {
		XMLName xml.Name
	}

	AutonomousTransferComplete struct {
		XMLName        xml.Name
		AnnounceURL    string       `xml:"AnnounceURL"`
		TransferURL    string       `xml:"TransferURL"`
		IsDownload     bool         `xml:"IsDownload"`
		FileType       string       `xml:"FileType"`
		FileSize       uint         `xml:"FileSize"`
		TargetFileName string       `xml:"TargetFileName"`
		FaultStruct    *FaultStruct `xml:"FaultStruct"`
		StartTime      time.Time    `xml:"StartTime"`
		CompleteTime   time.Time    `xml:"CompleteTime"`
	}

	AutonomousTransferCompleteResponse struct {
		XMLName xml.Name
	}

	GetParameterNames struct {
		XMLName       xml.Name
		ParameterPath string `xml:"ParameterPath"`
		NextLevel     bool   `xml:"NextLevel"`
	}

	GetParameterNamesResponse struct {
		XMLName       xml.Name
		ParameterList *ParameterInfoList `xml:"ParameterList"`
	}

	GetParameterValues struct {
		XMLName        xml.Name
		ParameterNames []string `xml:"ParameterNames>string"`
	}

	GetParameterValuesResponse struct {
		XMLName       xml.Name
		ParameterList *ParameterValueList `xml:"ParameterList"`
	}

	SetParameterValues struct {
		XMLName       xml.Name
		ParameterList *ParameterValueList `xml:"ParameterList"`
		ParameterKey  string              `xml:"ParameterKey"`
	}

	SetParameterValuesResponse struct {
		XMLName xml.Name
		Status  int `xml:"Status"`
	}

	GetParameterAttributes struct {
		XMLName        xml.Name
		ParameterNames []string `xml:"ParameterNames>string"`
	}

	GetParameterAttributesResponse struct {
		XMLName       xml.Name
		ParameterList *ParameterAttributeList `xml:"ParameterAttributeList"`
	}

	SetParameterAttributes struct {
		XMLName       xml.Name
		ParameterList *SetParameterAttributesList `xml:"ParameterList"`
	}

	SetParameterAttributesResponse struct {
		XMLName xml.Name
	}

	AddObject struct {
		XMLName      xml.Name
		ObjectName   string `xml:"ObjectName"`
		ParameterKey string `xml:"ParameterKey"`
	}

	AddObjectResponse struct {
		XMLName        xml.Name
		InstanceNumber uint `xml:"InstanceNumber"`
		Status         int  `xml:"Status"`
	}

	DeleteObject struct {
		XMLName      xml.Name
		ObjectName   string `xml:"ObjectName"`
		ParameterKey string `xml:"ParameterKey"`
	}

	DeleteObjectResponse struct {
		XMLName xml.Name
		Status  int `xml:"Status"`
	}

	Reboot struct {
		XMLName    xml.Name
		CommandKey string `xml:"CommandKey"`
	}

	RebootResponse struct {
		XMLName xml.Name
	}

	Download struct {
		XMLName        xml.Name
		CommandKey     string `xml:"CommandKey"`
		FileType       string `xml:"FileType"`
		URL            string `xml:"URL"`
		Username       string `xml:"Username"`
		Password       string `xml:"Password"`
		FileSize       uint   `xml:"FileSize"`
		TargetFileName string `xml:"TargetFileName"`
		DelaySeconds   uint   `xml:"DelaySeconds"`
		SuccessURL     string `xml:"SuccessURL"`
		FailureURL     string `xml:"FailureURL"`
	}

	DownloadResponse struct {
		XMLName      xml.Name
		Status       int       `xml:"Status"`
		StartTime    time.Time `xml:"StartTime"`
		CompleteTime time.Time `xml:"CompleteTime"`
	}
)

// Model Discriminators

func (m *GetRPCMethods) ModelDiscriminator() models.ModelDiscriminator {
	return GetRPCMethodsModelDiscriminator
}

func (m *GetRPCMethodsResponse) ModelDiscriminator() models.ModelDiscriminator {
	return GetRPCMethodsResponseModelDiscriminator
}

func (m *Inform) ModelDiscriminator() models.ModelDiscriminator {
	return InformModelDiscriminator
}

func (m *InformResponse) ModelDiscriminator() models.ModelDiscriminator {
	return InformResponseModelDiscriminator
}

func (m *TransferComplete) ModelDiscriminator() models.ModelDiscriminator {
	return TransferCompleteModelDiscriminator
}

func (m *TransferCompleteResponse) ModelDiscriminator() models.ModelDiscriminator {
	return TransferCompleteResponseModelDiscriminator
}

func (m *AutonomousTransferComplete) ModelDiscriminator() models.ModelDiscriminator {
	return AutonomousTransferCompleteModelDiscriminator
}

func (m *AutonomousTransferCompleteResponse) ModelDiscriminator() models.ModelDiscriminator {
	return AutonomousTransferCompleteResponseModelDiscriminator
}

func (m *GetParameterNames) ModelDiscriminator() models.ModelDiscriminator {
	return GetParameterNamesModelDiscriminator
}

func (m *GetParameterNamesResponse) ModelDiscriminator() models.ModelDiscriminator {
	return GetParameterNamesResponseModelDiscriminator
}

func (m *GetParameterValues) ModelDiscriminator() models.ModelDiscriminator {
	return GetParameterValuesModelDiscriminator
}

func (m *GetParameterValuesResponse) ModelDiscriminator() models.ModelDiscriminator {
	return GetParameterValuesResponseModelDiscriminator
}

func (m *SetParameterValues) ModelDiscriminator() models.ModelDiscriminator {
	return SetParameterValuesModelDiscriminator
}

func (m *SetParameterValuesResponse) ModelDiscriminator() models.ModelDiscriminator {
	return SetParameterValuesResponseModelDiscriminator
}

func (m *GetParameterAttributes) ModelDiscriminator() models.ModelDiscriminator {
	return GetParameterAttributesModelDiscriminator
}

func (m *GetParameterAttributesResponse) ModelDiscriminator() models.ModelDiscriminator {
	return GetParameterAttributesResponseModelDiscriminator
}

func (m *SetParameterAttributes) ModelDiscriminator() models.ModelDiscriminator {
	return SetParameterAttributesModelDiscriminator
}

func (m *SetParameterAttributesResponse) ModelDiscriminator() models.ModelDiscriminator {
	return SetParameterAttributesResponseModelDiscriminator
}

func (m *AddObject) ModelDiscriminator() models.ModelDiscriminator {
	return AddObjectModelDiscriminator
}

func (m *AddObjectResponse) ModelDiscriminator() models.ModelDiscriminator {
	return AddObjectResponseModelDiscriminator
}

func (m *DeleteObject) ModelDiscriminator() models.ModelDiscriminator {
	return DeleteObjectModelDiscriminator
}

func (m *DeleteObjectResponse) ModelDiscriminator() models.ModelDiscriminator {
	return DeleteObjectResponseModelDiscriminator
}

func (m *Reboot) ModelDiscriminator() models.ModelDiscriminator {
	return RebootModelDiscriminator
}

func (m *RebootResponse) ModelDiscriminator() models.ModelDiscriminator {
	return RebootResponseModelDiscriminator
}

func (m *Download) ModelDiscriminator() models.ModelDiscriminator {
	return DownloadModelDiscriminator
}

func (m *DownloadResponse) ModelDiscriminator() models.ModelDiscriminator {
	return DownloadResponseModelDiscriminator
}

// Message Discriminators

func (m *GetRPCMethods) MessageDiscriminator() models.MessageDiscriminator {
	return GetRPCMethodsMessageDiscriminator
}

func (m *GetRPCMethodsResponse) MessageDiscriminator() models.MessageDiscriminator {
	return GetRPCMethodsResponseMessageDiscriminator
}

func (m *Inform) MessageDiscriminator() models.MessageDiscriminator {
	return InformMessageDiscriminator
}

func (m *InformResponse) MessageDiscriminator() models.MessageDiscriminator {
	return InformResponseMessageDiscriminator
}

func (m *TransferComplete) MessageDiscriminator() models.MessageDiscriminator {
	return TransferCompleteMessageDiscriminator
}

func (m *TransferCompleteResponse) MessageDiscriminator() models.MessageDiscriminator {
	return TransferCompleteResponseMessageDiscriminator
}

func (m *AutonomousTransferComplete) MessageDiscriminator() models.MessageDiscriminator {
	return AutonomousTransferCompleteMessageDiscriminator
}

func (m *AutonomousTransferCompleteResponse) MessageDiscriminator() models.MessageDiscriminator {
	return AutonomousTransferCompleteResponseMessageDiscriminator
}

func (m *GetParameterNames) MessageDiscriminator() models.MessageDiscriminator {
	return GetParameterNamesMessageDiscriminator
}

func (m *GetParameterNamesResponse) MessageDiscriminator() models.MessageDiscriminator {
	return GetParameterNamesResponseMessageDiscriminator
}

func (m *GetParameterValues) MessageDiscriminator() models.MessageDiscriminator {
	return GetParameterValuesMessageDiscriminator
}

func (m *GetParameterValuesResponse) MessageDiscriminator() models.MessageDiscriminator {
	return GetParameterValuesResponseMessageDiscriminator
}

func (m *SetParameterValues) MessageDiscriminator() models.MessageDiscriminator {
	return SetParameterValuesMessageDiscriminator
}

func (m *SetParameterValuesResponse) MessageDiscriminator() models.MessageDiscriminator {
	return SetParameterValuesResponseMessageDiscriminator
}

func (m *GetParameterAttributes) MessageDiscriminator() models.MessageDiscriminator {
	return GetParameterAttributesMessageDiscriminator
}

func (m *GetParameterAttributesResponse) MessageDiscriminator() models.MessageDiscriminator {
	return GetParameterAttributesResponseMessageDiscriminator
}

func (m *SetParameterAttributes) MessageDiscriminator() models.MessageDiscriminator {
	return SetParameterAttributesMessageDiscriminator
}

func (m *SetParameterAttributesResponse) MessageDiscriminator() models.MessageDiscriminator {
	return SetParameterAttributesResponseMessageDiscriminator
}

func (m *AddObject) MessageDiscriminator() models.MessageDiscriminator {
	return AddObjectMessageDiscriminator
}

func (m *AddObjectResponse) MessageDiscriminator() models.MessageDiscriminator {
	return AddObjectResponseMessageDiscriminator
}

func (m *DeleteObject) MessageDiscriminator() models.MessageDiscriminator {
	return DeleteObjectMessageDiscriminator
}

func (m *DeleteObjectResponse) MessageDiscriminator() models.MessageDiscriminator {
	return DeleteObjectResponseMessageDiscriminator
}

func (m *Reboot) MessageDiscriminator() models.MessageDiscriminator {
	return RebootMessageDiscriminator
}

func (m *RebootResponse) MessageDiscriminator() models.MessageDiscriminator {
	return RebootResponseMessageDiscriminator
}

func (m *Download) MessageDiscriminator() models.MessageDiscriminator {
	return DownloadMessageDiscriminator
}

func (m *DownloadResponse) MessageDiscriminator() models.MessageDiscriminator {
	return DownloadResponseMessageDiscriminator
}
