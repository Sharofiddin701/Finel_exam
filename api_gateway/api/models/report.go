package models

import "api_gateway/genproto/schedule_service"

type TeachersReport struct {
	FullName     string  `json:"full_name"`
	Phone        string  `json:"phone"`
	Salary       float64 `json:"salary"`
	MonthsWorked int64   `json:"month_worked"`
	TotalSum     float64 `json:"total_sum"`
	IeltsScore   string  `json:"ielts_score"`
}

type StudentPanelResponse struct {
	Scores             []*schedule_service.Score
	Schedules          []*schedule_service.Schedule
	EventsRegistration []*schedule_service.AssignStudent
}

type SupportTeacherPanelResponse struct {
	Groups []*schedule_service.Group
}

type TeacherPanelResponse struct {
	Groups []*schedule_service.Group
}
