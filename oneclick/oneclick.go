package oneclick

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
func (e *Event) GetClickType() string {
	return e.DeviceEvent.ButtonClicked.ClickType
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
