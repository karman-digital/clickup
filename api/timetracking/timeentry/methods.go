package timeentry

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/karman-digital/clickup/api/shared"
	sharedmodels "github.com/karman-digital/clickup/models/shared"
	timetrackingmodels "github.com/karman-digital/clickup/models/timetracking"
)

func (t *TimeEntryService) GetTimeEntryHistory(id string) (timetrackingmodels.TimeTrackHistoryResponse, error) {
	var timeTrackHistory timetrackingmodels.TimeTrackHistoryResponse
	resp, err := t.SendTimeTrackingRequest(http.MethodGet, fmt.Sprintf("/team/%s/time_entries/%s/history", t.GetTeamId(), id), nil)
	if err != nil {
		return timetrackingmodels.TimeTrackHistoryResponse{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return timetrackingmodels.TimeTrackHistoryResponse{}, err
	}
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return timetrackingmodels.TimeTrackHistoryResponse{}, shared.ErrResourceNotFound
		}
		return timetrackingmodels.TimeTrackHistoryResponse{}, fmt.Errorf("error body: %s", string(body))
	}
	err = json.Unmarshal(body, &timeTrackHistory)
	if err != nil {
		return timetrackingmodels.TimeTrackHistoryResponse{}, err
	}
	return timeTrackHistory, nil
}

func (t *TimeEntryService) CreateTimeEntry(timeEntry timetrackingmodels.TimeEntry, opts ...sharedmodels.GetOptions) (timetrackingmodels.TimeEntryResponse, error) {
	responseTimeEntry := timetrackingmodels.TimeEntryResponse{}
	requestBody, err := json.Marshal(timeEntry)
	if err != nil {
		return timetrackingmodels.TimeEntryResponse{}, err
	}
	resp, err := t.SendTimeTrackingRequest(http.MethodPost, fmt.Sprintf("/team/%s/time_entries", t.GetTeamId()), requestBody, opts...)
	if err != nil {
		return timetrackingmodels.TimeEntryResponse{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return timetrackingmodels.TimeEntryResponse{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return timetrackingmodels.TimeEntryResponse{}, fmt.Errorf("error body: %s", string(body))
	}
	err = json.Unmarshal(body, &responseTimeEntry)
	if err != nil {
		return timetrackingmodels.TimeEntryResponse{}, err
	}
	return responseTimeEntry, nil
}

func (t *TimeEntryService) DeleteTimeEntry(id string) error {
	resp, err := t.SendTimeTrackingRequest(http.MethodDelete, fmt.Sprintf("/team/%s/time_entries/%s", t.GetTeamId(), id), nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error body: %s", string(body))
	}
	return nil
}
