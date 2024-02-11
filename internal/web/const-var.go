package web

import (
	"embed"

	"github.com/aceberg/CheckList/internal/models"
)

var (
	// appConfig - config for Web Gui
	appConfig models.Conf

	// plans - checks from plan.yaml
	plans []models.Plan

	// checks - checks from DB
	checks []models.Check
)

// templFS - html templates
//
//go:embed templates/*
var templFS embed.FS

// pubFS - public folder
//
//go:embed public/*
var pubFS embed.FS
