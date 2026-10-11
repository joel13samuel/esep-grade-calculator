package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeHighA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 95, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 91, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
    gradeCalculator := NewGradeCalculator()

    gradeCalculator.AddGrade("assignment", 75, Assignment)
    gradeCalculator.AddGrade("exam", 75, Exam)
    gradeCalculator.AddGrade("essay", 75, Essay)

    actual := gradeCalculator.GetFinalGrade()

    if actual != "C" {
        t.Errorf("Expected C, got %s", actual)
    }
}

func TestGetGradeD(t *testing.T) {
    gradeCalculator := NewGradeCalculator()

    gradeCalculator.AddGrade("assignment", 65, Assignment)
    gradeCalculator.AddGrade("exam", 65, Exam)
    gradeCalculator.AddGrade("essay", 65, Essay)

    actual := gradeCalculator.GetFinalGrade()

    if actual != "D" {
        t.Errorf("Expected D, got %s", actual)
    }
}

func TestGetGradeF(t *testing.T) {
    gradeCalculator := NewGradeCalculator()

    gradeCalculator.AddGrade("assignment", 50, Assignment)
    gradeCalculator.AddGrade("exam", 50, Exam)
    gradeCalculator.AddGrade("essay", 50, Essay)

    actual := gradeCalculator.GetFinalGrade()

    if actual != "F" {
        t.Errorf("Expected F, got %s", actual)
    }
}

func TestComputeAverageEmpty(t *testing.T) {
    grades := []Grade{}

    actual := computeAverage(grades)

    if actual != 0 {
        t.Errorf("Expected 0 for empty grades, got %d", actual)
    }
}

func TestGradeTypeString(t *testing.T) {
    tests := []struct {
        gradeType GradeType
        expected  string
    }{
        {Assignment, "assignment"},
        {Exam, "exam"},
        {Essay, "essay"},
    }

    for _, test := range tests {
        actual := test.gradeType.String()

        if actual != test.expected {
            t.Errorf("Expected %s, got %s", test.expected, actual)
        }
    }
}
