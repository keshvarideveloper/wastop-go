package entity

type User struct {
	ID       uint   `json:"id"`
	FullName string `json:"fullName"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Photo    string `json:"photo"`
	Bio      string `json:"bio"`
	Status   string `json:"status"`
}

// User Status
const (
	Completed  string = "completed"
	Suspended         = "suspended"
	Registered        = "registered"
)
