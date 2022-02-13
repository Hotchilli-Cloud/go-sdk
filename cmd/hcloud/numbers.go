package hcloud

import (
	"context"
	"fmt"
	"net/http"
)

type GetNumberResponse struct {
	NTSID                    string `json:"NTSID"`
	ClientID                 string `json:"ClientID"`
	CLI                      string `json:"CLI"`
	Number                   string `json:"Number"`
	EnableDivert             string `json:"EnableDivert"`
	Divert1To                string `json:"Divert1To"`
	Ring1Timeout             string `json:"Ring1Timeout"`
	Ring2Timeout             string `json:"Ring2Timeout"`
	Ring3Timeout             string `json:"Ring3Timeout"`
	Ring4Timeout             string `json:"Ring4Timeout"`
	Ring5Timeout             string `json:"Ring5Timeout"`
	np                       string `json:"np"`
	Divert2To                string `json:"Divert2To"`
	Divert3To                string `json:"Divert3To"`
	Divert4To                string `json:"Divert4To"`
	Divert5To                string `json:"Divert5To"`
	EnableVoicemail          string `json:"EnableVoicemail"`
	EnableWhisper            string `json:"EnableWhisper"`
	VmonBusy                 string `json:"VmonBusy"`
	RecordBusyMessage        string `json:"RecordBusyMessage"`
	BusyMessage              string `json:"BusyMessage"`
	VmonNoAns                string `json:"VmonNoAns"`
	RecordNoAnswerMessage    string `json:"RecordNoAnswerMessage"`
	NoAnswerMessage          string `json:"NoAnswerMessage"`
	EnableTimeOfDay          string `json:"EnableTimeOfDay"`
	VmonOutOfHours           string `json:"VmonOutOfHours"`
	OutOfHours               string `json:"OutOfHours"`
	SendToOutOfHours         string `json:"SendToOutOfHours"`
	StartTime                string `json:"StartTime"`
	endTime                  string `json:"endTime"`
	VmonTimeOfDay2           string `json:"VmonTimeOfDay2"`
	StartTime2               string `json:"StartTime2"`
	endtime2                 string `json:"endtime2"`
	Intro                    string `json:"Intro"`
	RecordIntro              string `json:"RecordIntro"`
	AdminPin                 string `json:"AdminPin"`
	AdminPinTemp             string `json:"AdminPinTemp"`
	AdminCLI                 string `json:"AdminCLI"`
	Mon                      string `json:"Mon"`
	Tue                      string `json:"Tue"`
	Wed                      string `json:"Wed"`
	Thu                      string `json:"Thu"`
	Fri                      string `json:"Fri"`
	Sat                      string `json:"Sat"`
	Sun                      string `json:"Sun"`
	NoAnswerSubject          string `json:"NoAnswerSubject"`
	NoAnswerBody             string `json:"NoAnswerBody"`
	BusySubject              string `json:"BusySubject"`
	BusyBody                 string `json:"BusyBody"`
	HSubject                 string `json:"HSubject"`
	HBody                    string `json:"HBody"`
	DivertOutOfHours         string `json:"DivertOutOfHours"`
	ParentClientID           string `json:"ParentClientID"`
	CheckQueue               string `json:"CheckQueue"`
	QueueLines               string `json:"QueueLines"`
	CallRecording            string `json:"CallRecording"`
	PlayIntro                string `json:"PlayIntro"`
	PlayWhisper              string `json:"PlayWhisper"`
	SendNP                   string `json:"SendNP"`
	Tandby                   string `json:"Tandby"`
	CallRecordingEmail       string `json:"CallRecordingEmail"`
	SendCallRecordingByEmail string `json:"SendCallRecordingByEmail"`
	EncryptCallRecording     string `json:"EncryptCallRecording"`
	EncryptKey               string `json:"EncryptKey"`
	TimeOfDayGroupID         string `json:"TimeOfDayGroupID"`
	AdvancedTimeOfDay        string `json:"AdvancedTimeOfDay"`
	MainRingGroup            string `json:"MainRingGroup"`
	SalesRingGroup           string `json:"SalesRingGroup"`
	SupportRingGroup         string `json:"SupportRingGroup"`
	AccountsRingGroup        string `json:"AccountsRingGroup"`
	CallCenterRingGroup      string `json:"CallCenterRingGroup"`
	OverFlowRingGroup        string `json:"OverFlowRingGroup"`
	OutOfHoursRingGroup      string `json:"OutOfHoursRingGroup"`
	QuickProfile             string `json:"QuickProfile"`
	QuickProfileStatus       string `json:"QuickProfileStatus"`
	Key1                     string `json:"Key1"`
	Key2                     string `json:"Key2"`
	Key3                     string `json:"Key3"`
	PlayOutOfHours           string `json:"PlayOutOfHours"`
	PlayNoAnswer             string `json:"PlayNoAnswer"`
	PlayBusy                 string `json:"PlayBusy"`
	PromptFolder             string `json:"PromptFolder"`
	PlayIntroID              string `json:"PlayIntroID"`
	PlayWhisperID            string `json:"PlayWhisperID"`
	PlayBusyID               string `json:"PlayBusyID"`
	PlayNoAnswerID           string `json:"PlayNoAnswerID"`
	PlayOutOfHoursID         string `json:"PlayOutOfHoursID"`
	VirtualReceptionist      string `json:"VirtualReceptionist"`
	VRPrompt                 string `json:"VRPrompt"`
	MissedCall               string `json:"MissedCall"`
	MissedCallDestination    string `json:"MissedCallDestination"`
	function                 string `json:"function"`
}

type GetNumberRequest struct {
	Number string `json:"number"`
}

func (c *Client) GetNumber(ctx context.Context, options *GetNumberRequest) (*SuccessResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("&s/node-api/numbers/data/summary?client_id=%d&year=%d", c.BaseURL, 2645, 2021), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := SuccessResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
