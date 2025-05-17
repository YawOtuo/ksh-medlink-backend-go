package interventionrecord

import (
	"ksh-medlink-backend/utils"
	"time"

	"gorm.io/gorm"
)

// Medication represents a single medication entry
type Medication struct {
	MedicationName string `json:"medicationName"`
	Dosage         string `json:"dosage"`
	Frequency      string `json:"frequency"`
	Route          string `json:"route"`
}

// PharmaceuticalIssue represents a single pharmaceutical care issue
type PharmaceuticalIssue struct {
	Description string `json:"description"`
	Severity    string `json:"severity"`
	ActionTaken string `json:"actionTaken"`
	Outcome     string `json:"outcome"`
}

// InterventionRecord represents the patient intervention data model
type InterventionRecord struct {
	gorm.Model
	PatientName              string                                `json:"patientName"`
	Age                      string                                `json:"age"`
	InpatientNo              string                                `json:"inpatientNo"`
	Sex                      string                                `json:"sex"`
	Ward                     string                                `json:"ward"`
	Weight                   string                                `json:"weight"`
	Date                     time.Time                             `json:"date"`
	DOA                      time.Time                             `json:"doa"`
	Allergy                  string                                `json:"allergy"`
	DrugHistory              string                                `json:"drugHistory"`
	PresentingComplaint      string                                `json:"presentingComplaint"`
	CurrentDiagnoses         string                                `json:"currentDiagnoses"`
	CurrentMedications       utils.JSONBArray[Medication]          `json:"currentMedications" gorm:"type:jsonb"`
	PharmaceuticalCareIssues utils.JSONBArray[PharmaceuticalIssue] `json:"pharmaceuticalCareIssues" gorm:"type:jsonb"`
	GeneralComments          string                                `json:"generalComments"`
	PharmacistName           string                                `json:"pharmacistName"`
	DocumentedBy             string                                `json:"documentedBy"`
}
