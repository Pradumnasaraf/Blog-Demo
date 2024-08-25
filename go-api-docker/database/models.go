package database

type Schedule struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Content string `json:"content"`
}

func (Schedule) TableName() string {
	return "schedule"
}

type ScheduleService interface {
	Get() []Schedule
	Create(schedule Schedule) Schedule
	Update(schedule Schedule) Schedule
	Delete(schedule Schedule)
}
