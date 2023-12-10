package logger

import (
	"github.com/mailhedgehog/gounit"
	"testing"
)

func TestTruncateStart(t *testing.T) {
	truncated := truncateStart("Hello foo bar!", 6, "._.")
	(*gounit.T)(t).AssertEqualsString("._.o bar!", truncated)

	truncated = truncateStart("Hello foo bar!", 3, "foobar_")
	(*gounit.T)(t).AssertEqualsString("foobar_ar!", truncated)

	truncated = truncateStart("Hi", 3, "123456_")
	(*gounit.T)(t).AssertEqualsString("Hi", truncated)
}
