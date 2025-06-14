package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AvailabilitySlot struct {
    Date     string `bson:"date"`     // e.g., "2025-06-15"
    TimeFrom string `bson:"timeFrom"` // e.g., "14:00"
    TimeTo   string `bson:"timeTo"`   // e.g., "16:00"
}

type Task struct {
    ID             primitive.ObjectID `bson:"_id,omitempty"`
    Title          string             `bson:"title"`                // e.g., "Need help with math homework"
    Description    string             `bson:"description,omitempty"`// Optional task details
    Location       string             `bson:"location"`             // Could be address or general area
	Latitude       float64            `bson:"latitude,omitempty"`
	Longitude      float64            `bson:"longitude,omitempty"`
    LocationType   string             `bson:"locationType"`         // "online" or "in-person"
    Credits        int                `bson:"credits"`              // Points required for task
    CreatedBy      primitive.ObjectID `bson:"createdBy"`            // Reference to User ID
    Availability   []AvailabilitySlot `bson:"availability"`         // Dates and times the task can be done
    CreatedAt      int64              `bson:"createdAt,omitempty"`  // Unix timestamp
    Status         string             `bson:"status"`               // "open", "in progress", "completed"
	AcceptedBy     primitive.ObjectID `bson:"acceptedBy,omitempty"` // Who is helping
	CompletedAt    int64              `bson:"completedAt,omitempty"`// When the task was completed
}