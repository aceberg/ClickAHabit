package web

import (
	"embed"

	"github.com/aceberg/ClickAHabit/internal/models"
)

var (
	// appConfig - config for Web Gui
	appConfig models.Conf

	// allPlans - checks from plan.yaml
	allPlans []models.Plan

	// allChecks - checks from DB
	allChecks []models.Check
	// allWeeks  []models.Check

	// lastToday - last date for which checks were set
	lastToday string
)

// templFS - html templates
//
//go:embed templates/*
var templFS embed.FS

// pubFS - public folder
//
//go:embed public/*
var pubFS embed.FS
