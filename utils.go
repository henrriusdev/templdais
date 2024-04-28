package main

import (
	"fmt"

	"github.com/a-h/templ"
)

type Links struct {
	Text string
	Href templ.SafeURL
}

type OptsStruct struct {
	Text  string
	Value string
}

type OptsString []string

type Options interface{}

type FormAttrs struct {
	FieldName     string
	FieldID       string
	FieldBrand    string
	FieldSize     string
	Checked       string
	Class         string
	FieldBordered bool
	Options       Options
	Placeholder   string
}

func (f FormAttrs) GetClassName(prefix string) string {
	class := getClassName(f.FieldBrand, "", f.FieldSize, prefix)
	if f.FieldBordered {
		class += " " + prefix + "-bordered"
	}

	if f.Class != "" {
		class += " " + f.Class
	}

	return class
}

// func (f FormAttrs) GetOptionsType() string {
// 	switch f.Options.(type) {
// }

func GetBrandColorClass(prefix, brand string) string {
	switch brand {
	case "default":
		return ""
	case "primary":
		return fmt.Sprintf("%s-primary", prefix)
	case "secondary":
		return fmt.Sprintf("%s-secondary", prefix)
	case "accent":
		return fmt.Sprintf("%s-accent", prefix)
	case "ghost":
		return fmt.Sprintf("%s-ghost", prefix)
	case "link":
		return fmt.Sprintf("%s-link", prefix)
	case "info":
		return fmt.Sprintf("%s-info", prefix)
	case "success":
		return fmt.Sprintf("%s-success", prefix)
	case "warning":
		return fmt.Sprintf("%s-warning", prefix)
	case "error":
		return fmt.Sprintf("%s-error", prefix)
	default:
		return ""
	}
}

func GetSizeClass(prefix, size string) string {
	switch size {
	case "xsmall":
		return fmt.Sprintf("%s-xs", prefix)
	case "small":
		return fmt.Sprintf("%s-sm", prefix)
	case "medium":
		return ""
	case "large":
		return fmt.Sprintf("%s-lg", prefix)
	default:
		return ""
	}
}

func GetShapeClass(prefix, shape string) string {
	switch shape {
	case "square":
		return fmt.Sprintf("%s-square", prefix)
	case "circle":
		return fmt.Sprintf("%s-circle", prefix)
	default:
		return ""
	}
}

func getClassName(brand, figure, size, prefix string) string {
	return fmt.Sprintf("%s %s %s %s", prefix, GetBrandColorClass(prefix, brand), GetShapeClass(prefix, figure), GetSizeClass(prefix, size))
}
