package main

import (
	"time"
)

type Lecture struct {
	EntityType string    `json:"entityType"`
	Date       time.Time `json:"date"`
	Site       string    `json:"site"`
	StartTime  time.Time `json:"startTime"`
	EndTime    time.Time `json:"endTime"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Lecturer   string    `json:"lecturer"`
	Rooms      []string  `json:"rooms"`
	Course     string    `json:"course"`
	ID         int       `json:"id"`
}

type byEndTime []Lecture

func (a byEndTime) Len() int           { return len(a) }
func (a byEndTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byEndTime) Less(i, j int) bool { return a[i].EndTime.Before(a[j].EndTime) }

type CachedLectures struct {
	LastRetrieved time.Time `json:"retrievedTime"`
	Lectures      []Lecture `json:"lectures"`
}
