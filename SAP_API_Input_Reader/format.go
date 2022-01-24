package sap_api_input_reader

import (
	"sap-api-integrations-maintenance-order-confirmation-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.MaintenanceOrderConfirmation
	return &requests.Header{
		MaintOrderConf:                 data.MaintOrderConf,
		MaintOrderConfCntrValue:        data.MaintOrderConfCntrValue,
		MaintenanceOrder:               data.MaintenanceOrder,
		MaintenanceOrderOperation:      data.MaintenanceOrderOperation,
		MaintenanceOrderSubOperation:   data.MaintenanceOrderSubOperation,
		PersonnelNumber:                data.PersonnelNumber,
		ActualWorkQuantity:             data.ActualWorkQuantity,
		ActualWorkQuantityUnit:         data.ActualWorkQuantityUnit,
		ActualDuration:                 data.ActualDuration,
		ActualDurationUnit:             data.ActualDurationUnit,
		OperationConfirmedStartDate:    data.OperationConfirmedStartDate,
		OperationConfirmedStartTime:    data.OperationConfirmedStartTime,
		OperationConfirmedEndDate:      data.OperationConfirmedEndDate,
		OperationConfirmedEndTime:      data.OperationConfirmedEndTime,
		IsFinalConfirmation:            data.IsFinalConfirmation,
		NoFurtherWorkQuantityIsExpd:    data.NoFurtherWorkQuantityIsExpd,
		RemainingWorkQuantity:          data.RemainingWorkQuantity,
		RemainingWorkQuantityUnit:      data.RemainingWorkQuantityUnit,
		PostingDate:                    data.PostingDate,
		ActivityType:                   data.ActivityType,
		OpenReservationsIsCleared:      data.OpenReservationsIsCleared,
		ConfirmationText:               data.ConfirmationText,
		VarianceReasonCode:             data.VarianceReasonCode,
		CapacityInternalID:             data.CapacityInternalID,
		NmbrOfMaintTechnicianCapSplits: data.NmbrOfMaintTechnicianCapSplits,
		MaterialDocument:               data.MaterialDocument,
		AccountingIndicatorCode:        data.AccountingIndicatorCode,
		ActyConfFcstdEndDate:           data.ActyConfFcstdEndDate,
		ActyConfFcstdEndTime:           data.ActyConfFcstdEndTime,
	}
}
