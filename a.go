package a

import c "github.com/AppliedGoCourses/C"

func A() string {
	return "a.A uses package c at version " + c.Version()
}
