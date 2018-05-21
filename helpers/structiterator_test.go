package helpers

import (
	"testing"
)

const (
	success = "\u2713"
	failed  = "\u2717"
)

func TestIsEmpty(t *testing.T) {
	t.Log("TestIsEmpty for empty struct")
	{
		tst := struct {
			Name string
		}{}

		if IsEmpty(tst) {
			t.Logf("%s expected empty", success)
		} else {
			t.Errorf("%s expected empty, got %b", failed, IsEmpty(tst))
		}
	}
	t.Log("TestIsEmpty for not empty struct")
	{
		tst := struct {
			Name string
		}{
			"Example Test",
		}

		if !IsEmpty(tst) {
			t.Logf("%s expected not empty", success)
		} else {
			t.Errorf("%s expected not empty, got empty", failed)
		}
	}
}

func TestStructIterator(t *testing.T) {
	t.Log("Test StructIterator")
	{
		tst1 := struct {
			FullName     string
			EmailAddress string
			Point        int
		}{
			"Example Test 1", "example@example.com", 151,
		}

		tst2 := struct {
			Name         string
			EmailAddress string
			Point        float64
		}{}

		StructIterator(tst1, &tst2)

		if tst2.Name == "" {
			t.Logf("%s expected tst2.Name is empty", success)
		} else {
			t.Errorf("%s expected tst2.Name is empty got %s", failed, tst2.Name)
		}

		if tst2.EmailAddress == "example@example.com" {
			t.Logf("%s expected tst2.EmailAddress = example@example.com", success)
		} else {
			t.Errorf("%s expected tst2.EmailAddress = example@example.com, got %s", failed, tst2.Name)
		}

		if tst2.Point != 0.0 {
			t.Errorf("%s expected tst2.Point != 0.0 got 0.0", failed)
		} else {
			t.Logf("%s expected tst2.Point = 0.0, got %f", success, tst2.Point)
		}
	}

	t.Log("Test StructIterator with tag field")
	{
		tst1 := struct {
			FullName     string `kmsg:"name"`
			EmailAddress string `kmsg:"email"`
			Point        int    `kmsg:"point"`
		}{
			"Example Test 1", "example@example.com", 151,
		}

		tst2 := struct {
			Name         string  `kmsg:"name"`
			EmailAddress string  `kmsg:"email"`
			Point        float64 `kmsg:"point"`
		}{}

		StructIterator(tst1, &tst2, "kmsg")

		if tst2.Name == "Example Test 1" {
			t.Logf("%s expected tst2.Name = Example Test 1", success)
		} else {
			t.Errorf("%s expected tst2.Name = Example Test 1, got %s", failed, tst2.Name)
		}

		if tst2.EmailAddress == "example@example.com" {
			t.Logf("%s expected tst2.EmailAddress = example@example.com", success)
		} else {
			t.Errorf("%s expected tst2.EmailAddress = example@example.com, got %s", failed, tst2.Name)
		}

		if tst2.Point != 0.0 {
			t.Errorf("%s expected tst2.Point != 0.0 got 0.0", failed)
		} else {
			t.Logf("%s expected tst2.Point = 0.0, got %f", success, tst2.Point)
		}
	}
}
