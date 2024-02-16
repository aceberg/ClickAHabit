package web

import (
	"embed"

	"github.com/aceberg/CheckList/internal/models"
)

var (
	// appConfig - config for Web Gui
	appConfig models.Conf

	// allPlans - checks from plan.yaml
	allPlans []models.Plan

	// allChecks - checks from DB
	allChecks []models.Check
)

// templFS - html templates
//
//go:embed templates/*
var templFS embed.FS

// pubFS - public folder
//
//go:embed public/*
var pubFS embed.FS
