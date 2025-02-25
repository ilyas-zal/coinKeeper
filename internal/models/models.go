// models содержит структуру данных пользователей и задач.
package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"user_id"`
	Name     string `gorm:"unique;not null" json:"user_name"`
	Password string
	Email    string `gorm:"unique;not null" json:"user_email"`
}

type Task struct {
	ID          uint   `gorm:"primaryKey" json:"task_id"`
	Description string `json:"task_description"`
	Reward      int    `json:"task_reward"`
}

type UserTask struct {
	ID        uint  `gorm:"primaryKey" json:"user_task_id"`
	UserID    uint  `json:"user_id" gorm:"foreignKey:UserID;references:ID"`
	TaskID    uint  `json:"task_id" gorm:"foreignKey:TaskID;references:ID"`
	Completed bool  `json:"task_completed"`
	User      *User `json:"user"`
	Task      *Task `json:"task"`
}
