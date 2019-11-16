package main

import (
	"log"
	"regexp"
)

func FileClassifier(fileName string) int {
	// Matching Sourcefiles
	r := regexp.MustCompile(`.*\.go$`)
	if (r.MatchString(fileName)) {
		return 0
	}
	// Matching Shellscripts
	r = regexp.MustCompile(`.*\.sh$`)
	if (r.MatchString(fileName)) {
		return 1
	}
	// Matching Makefiles
	r = regexp.MustCompile(`^Makefile$`)
	if (r.MatchString(fileName)) {
		return 2
	}
	// Matching Dockerfiles
	r = regexp.MustCompile(`^Dockerfile$`)
	if (r.MatchString(fileName)) {
		return 3
	}
	// Matching Configfiles
	r = regexp.MustCompile(`.*\.(env|cfg)$`)
	if (r.MatchString(fileName)) {
		return 4
	}
	// Matching Staticfiles
	r = regexp.MustCompile(`.*\.(html|css|scss)$`)
	if (r.MatchString(fileName)) {
		return 5
	}
	// Matching Documents
	r = regexp.MustCompile(`.*\.(md|txt)$`)
	if (r.MatchString(fileName)) {
		return 6
	}
	// Matching Images
	r = regexp.MustCompile(`.*\.(jpeg|jpg|png|svc)$`)
	if (r.MatchString(fileName)) {
		return 7
	}
	// Others
	log.Println(fileName + " is classified as a others.")
	return 8
}
