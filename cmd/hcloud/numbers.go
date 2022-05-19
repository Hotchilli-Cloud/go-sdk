package hcloud

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TimeOfDay struct {
	Enabled               bool   `json:"enabled"`
	StartTimeHours        string `json:"start_time_hours"`
	StartTimeMinutes      string `json:"start_time_minutes"`
	EndTimeHours          string `json:"end_time_hours"`
	EndTimeMinutes        string `json:"end_time_minutes"`
	OutOfHoursToVoicemail bool   `json:"out_of_hours_to_voicemail"`
	OutOfHoursToNumber    string `json:"out_of_hours_to_number"`
	EnabledMonday         bool   `json:"enabled_monday"`
	EnabledTuesday        bool   `json:"enabled_tuesday"`
	EnabledWednesday      bool   `json:"enabled_wednesday"`
	EnabledThursday       bool   `json:"enabled_thursday"`
	EnabledFriday         bool   `json:"enabled_friday"`
	EnabledSaturday       bool   `json:"enabled_saturday"`
	EnabledSunday         bool   `json:"enabled_sunday"`
}

type CallForwarding struct {
	Forward1      string `json:"forward_1"`
	Forward1Timer string `json:"forward_1_timer"`
	Forward2      string `json:"forward_2"`
	Forward2Timer string `json:"forward_2_timer"`
	Forward3      string `json:"forward_3"`
	Forward3Timer string `json:"forward_3_timer"`
	Forward4      string `json:"forward_4"`
	Forward4Timer string `json:"forward_4_timer"`
	Forward5      string `json:"forward_5"`
	Forward5Timer string `json:"forward_5_timer"`
	Routing       string `json:"routing"`
}

type PromptRecorder struct {
	UserID int32  `json:"user_id"`
	Pin    string `json:"pin"`
}

type Voicemail struct {
	Enabled          bool   `json:"enabled"`
	OutOfHoursEmail  string `json:"out_of_hours_email"`
	NoAnswerEmail    string `json:"no_answer_email"`
	BusyEmail        string `json:"busy_email"`
	OutOfHoursPrompt string `json:"out_of_hours_prompt"`
	NoAnswerPrompt   string `json:"no_answer_prompt"`
	BusyPrompt       string `json:"busy_prompt"`
}

type MissedCallAlert struct {
	Enabled bool   `json:"enabled"`
	SmsTo   string `json:"sms_to"`
}

type NumberConfig struct {
	PhoneNumber                string          `json:"phone_number"`
	RouteID                    int32           `json:"route_id"`
	TimeOfDay                  TimeOfDay       `json:"time_of_day"`
	CallForwarding             CallForwarding  `json:"call_forwarding"`
	CallRecording              bool            `json:"call_recording"`
	PlayWhisper                string          `json:"play_whisper"`
	NumberPresentation         bool            `json:"number_presentation"`
	PromptRecorder             PromptRecorder  `json:"prompt_recorder"`
	Voicemail                  Voicemail       `json:"voicemail"`
	MissedCallAlert            MissedCallAlert `json:"missed_call_alert"`
	ApplicationID              int32           `json:"application"`
	PlayIntro                  string          `json:"play_intro"`
	VirtualReceptionistEnabled bool            `json:"virtual_receptionist_enabled"`
	VirtualReceptionistPrompt  string          `json:"virtual_receptionist_prompt"`
}

type GetNumberRequest struct {
	Number string
}

func (c *Client) GetNumber(phone_number string, route_id string) (*NumberConfig, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/node-api/numbers/config/%s?phone_number=%s", c.BaseURL, route_id, phone_number), nil)
	if err != nil {
		return nil, err
	}

	res := NumberConfig{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) UpdateNumber(phone_number string, route_id string, params *NumberConfig) (string, error) {
	parameters, _ := json.Marshal(params)
	print(string(parameters))
	return "updated", nil
	//req, err := http.NewRequest("PUT", fmt.Sprintf("%s/node-api/numbers/config/%s?phone_number=%s", c.BaseURL, route_id, phone_number), parameters)
}
