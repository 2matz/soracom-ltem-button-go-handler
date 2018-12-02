package oneclick

import (
	"fmt"
)

const (
	// SingleClick is pressing a button once.
	SingleClick = "SINGLE"
	// DoubleClick is pressing a button twice quickly.
	DoubleClick = "DOUBLE"
	// LongClick is press and keep holding down a button over 1.2 secs.
	LongClick = "LONG"
)

type (
	// Event ...
	Event struct {
		DeviceInfo    DeviceInfo
		DeviceEvent   DeviceEvent
		PlacementInfo PlacementInfo
	}

	// DeviceInfo ...
	DeviceInfo struct {
		DeviceID      string
		Type          string
		RemainingLife float32
		Attributes    Attributes
	}

	// Attributes ...
	Attributes struct {
		ProjectRegion      string
		ProjectName        string
		PlacementName      string
		DeviceTemplateName string
	}

	// DeviceEvent ...
	DeviceEvent struct {
		ButtonClicked ButtonClicked
	}

	// ButtonClicked ...
	ButtonClicked struct {
		ClickType    string
		ReportedTime string
	}

	// PlacementInfo ...
	PlacementInfo struct {
		ProjectName   string
		PlacementName string
		Attributes    interface{}
		Devices       Devices
	}

	// Devices ...
	Devices struct {
		LambdaCaller string
	}
)

// GetClickType return click type of button click event that one of "SINGLE", "DOUBLE" or "LONG".
func (e *Event) GetClickType() (string, error) {
	switch clickType := e.DeviceEvent.ButtonClicked.ClickType; clickType {
	case SingleClick:
		return SingleClick, nil
	case DoubleClick:
		return DoubleClick, nil
	case LongClick:
		return LongClick, nil
	default:
		return "", fmt.Errorf("%s", "Click type is not defined.")
	}

}

// GetDeviceID return device serial number of button.
func (e *Event) GetDeviceID() (buttonName string) {
	buttonName = e.DeviceInfo.DeviceID
	return
}

// GetPlacementName return placement name if that related to button.
func (e *Event) GetPlacementName() (placementName string) {
	placementName = e.PlacementInfo.PlacementName
	return
}

// GetPlacementAttributes return attributes of placement that related to button.
func (e *Event) GetPlacementAttributes() (placementAttributes interface{}) {
	return e.PlacementInfo.Attributes
}

// GetProjectName return project name if that related to button.
func (e *Event) GetProjectName() (projectName string) {
	projectName = e.DeviceInfo.Attributes.PlacementName
	return
}
