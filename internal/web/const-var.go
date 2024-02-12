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

	// flatPlans - allPlans in one slice
	flatPlans []models.Check

	// groupList - list of groups from plan
	groupList []string

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
