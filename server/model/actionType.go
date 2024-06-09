package model

type ActionType string

const (
	Pending  ActionType = "click"
	Approved ActionType = "input"
)
