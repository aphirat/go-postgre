package api

type (
	// Response struct for REST API
	Response struct {
		Response Rest        `json:"response"`
		Data     interface{} `json:"data,omitempty"`
	}
	// Rest struct for response Res API
	Rest struct {
		Message string `json:"message,omitempty"`
		Code    string `json:"code,omitempty"`
	}
	// Datas struct for data
	Datas struct {
		RmID      string `json:"id,omitempty"`
		RmName    string `json:"name,omitempty"`
		RmNameEng string `json:"name_eng,omitempty"`
	}