package sap_api_output_formatter

type MaintenanceOrderConfirmation struct {
	ConnectionKey                string `json:"connection_key"`
	Result                       bool   `json:"result"`
	RedisKey                     string `json:"redis_key"`
	Filepath                     string `json:"filepath"`
	Product                      string `json:"Product"`
	APISchema                    string `json:"api_schema"`
	MaintenanceOrderConfirmation string `json:"maintenance_order_confirmation"`
	Deleted                      string `json:"deleted"`
}

type Header struct {
	MaintOrderConf                 string `json:"MaintOrderConf"`
	MaintOrderConfCntrValue        string `json:"MaintOrderConfCntrValue"`
	MaintenanceOrder               string `json:"MaintenanceOrder"`
	MaintenanceOrderOperation      string `json:"MaintenanceOrderOperation"`
	MaintenanceOrderSubOperation   string `json:"MaintenanceOrderSubOperation"`
	PersonnelNumber                string `json:"PersonnelNumber"`
	ActualWorkQuantity             string `json:"ActualWorkQuantity"`
	ActualWorkQuantityUnit         string `json:"ActualWorkQuantityUnit"`
	ActualDuration                 string `json:"ActualDuration"`
	ActualDurationUnit             string `json:"ActualDurationUnit"`
	OperationConfirmedStartDate    string `json:"OperationConfirmedStartDate"`
	OperationConfirmedStartTime    string `json:"OperationConfirmedStartTime"`
	OperationConfirmedEndDate      string `json:"OperationConfirmedEndDate"`
	OperationConfirmedEndTime      string `json:"OperationConfirmedEndTime"`
	IsFinalConfirmation            bool   `json:"IsFinalConfirmation"`
	NoFurtherWorkQuantityIsExpd    bool   `json:"NoFurtherWorkQuantityIsExpd"`
	RemainingWorkQuantity          string `json:"RemainingWorkQuantity"`
	RemainingWorkQuantityUnit      string `json:"RemainingWorkQuantityUnit"`
	PostingDate                    string `json:"PostingDate"`
	ActivityType                   string `json:"ActivityType"`
	OpenReservationsIsCleared      bool   `json:"OpenReservationsIsCleared"`
	ConfirmationText               string `json:"ConfirmationText"`
	VarianceReasonCode             string `json:"VarianceReasonCode"`
	CapacityInternalID             string `json:"CapacityInternalID"`
	NmbrOfMaintTechnicianCapSplits int    `json:"NmbrOfMaintTechnicianCapSplits"`
	MaterialDocument               string `json:"MaterialDocument"`
	AccountingIndicatorCode        string `json:"AccountingIndicatorCode"`
	ActyConfFcstdEndDate           string `json:"ActyConfFcstdEndDate"`
	ActyConfFcstdEndTime           string `json:"ActyConfFcstdEndTime"`
}
