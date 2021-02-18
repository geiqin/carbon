package carbon

import (
	"testing"
)

func TestCarbon_DiffInYears(t *testing.T) {
	Tests := []struct {
		input1 string // 输入值1
		input2 string // 输入值2
		output int64  // 期望输出值
	}{
		{"2020-08-05 13:14:15", "2020-07-28 13:14:00", 0},
		{"2020-12-31 13:14:15", "2021-01-01 13:14:15", 0},
		{"2020-08-05 13:14:15", "2021-08-28 13:14:59", 1},
		{"2020-08-05 13:14:15", "2018-08-28 13:14:59", -2},
	}

	for _, v := range Tests {
		output := Parse(v.input1).DiffInYears(Parse(v.input2))

		if output != v.output {
			t.Errorf("Input start time %s and end time %s, expected %d, but got %d", v.input1, v.input2, v.output, output)
		}
	}
}

func Test1Carbon_DiffInYearsWithAbs(t *testing.T) {
	Tests := []struct {
		input1 string // 输入值1
		input2 string // 输入值2
		output int64  // 期望输出值
	}{
		{"2020-08-05 13:14:15", "2020-07-28 13:14:00", 0},
		{"2020-12-31 13:14:15", "2021-01-01 13:14:15", 0},
		{"2020-08-05 13:14:15", "2021-08-28 13:14:59", 1},
		{"2020-08-05 13:14:15", "2018-08-28 13:14:59", 2},
	}

	for _, v := range Tests {
		output := Parse(v.input1).DiffInYearsWithAbs(Parse(v.input2))

		if output != v.output {
			t.Errorf("Input start time %s and end time %s, expected %d, but got %d", v.input1, v.input2, v.output, output)
		}
	}
}

func TestCarbon_DiffInMonths(t *testing.T) {
	Tests := []struct {
		input1 string // 输入值1
		input2 string // 输入值2
		output int64  // 期望输出值
	}{
		{"2020-08-05 13:14:15", "2020-07-28 13:14:00", 0},
		{"2020-12-31 13:14:15", "2021-01-31 13:14:15", 1},
		{"2020-08-05 13:14:15", "2021-08-28 13:14:59", 12},
		{"2020-08-05 13:14:15", "2018-08-28 13:14:59", -24},
	}

	for _, v := range Tests {
		output := Parse(v.input1).DiffInMonths(Parse(v.input2))

		if output != v.output {
			t.Errorf("Input start time %s and end time %s, expected %d, but got %d", v.input1, v.input2, v.output, output)
		}
	}
}

func TestCarbon_DiffInMonthsWithAbs(t *testing.T) {
	Tests := []struct {
		input1 string // 输入值1
		input2 string // 输入值2
		output int64  // 期望输出值
	}{
		{"2020-08-05 13:14:15", "2020-07-28 13:14:00", 0},
		{"2020-12-31 13:14:15", "2021-01-01 13:14:15", 0},
		{"2020-08-05 13:14:15", "2021-08-28 13:14:59", 12},
		{"2020-08-05 13:14:15", "2018-08-28 13:14:59", 24},
	}

	for _, v := range Tests {
		output := Parse(v.input1).DiffInMonthsWithAbs(Parse(v.input2))

		if output != v.output {
			t.Errorf("Input start time %s and end time %s, expected %d, but got %d", v.input1, v.input2, v.output, output)
		}
	}
}

func TestCarbon_DiffInWeeks(t *testing.T) {
	Tests := []struct {
		input1 string // 输入值1
		input2 string // 输入值2
		output int64  // 期望输出值
	}{
		{"0000-00-00 00:00:00", "", 0},
		{"", "0000-00-00 00:00:00", 0},
		{"", "", 0},
		{"2020-08-05 13:14:15", "", -2639},
		{"", "2020-08-05 13:14:15", 2639},
		{"2020-08-05 13:14:15", "2020-07-28 13:14:00", -1},
		{"2020-08-05 13:14:15", "2020-07-28 13:14:15", -1},
		{"2020-08-05 13:14:15", "2020-07-28 13:14:59", -1},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-12 13:14:00", 0},
		{"2020-08-05 13:14:15", "2020-08-12 13:14:15", 1},
		{"2020-08-05 13:14:15", "2020-08-12 13:14:59", 1},
	}

	for _, v := range Tests {
		output := Parse(v.input1).DiffInWeeks(Parse(v.input2))

		if output != v.output {
			t.Errorf("Input start time %s and end time %s, expected %d, but got %d", v.input1, v.input2, v.output, output)
		}
	}
}

func TestCarbon_DiffInWeeksWithAbs(t *testing.T) {
	Tests := []struct {
		input1 string // 输入值1
		input2 string // 输入值2
		output int64  // 期望输出值
	}{
		{"0000-00-00 00:00:00", "", 0},
		{"", "0000-00-00 00:00:00", 0},
		{"", "", 0},
		{"2020-08-05 13:14:15", "", 2639},
		{"", "2020-08-05 13:14:15", 2639},
		{"2020-08-05 13:14:15", "2020-07-28 13:14:00", 1},
		{"2020-08-05 13:14:15", "2020-07-28 13:14:15", 1},
		{"2020-08-05 13:14:15", "2020-07-28 13:14:59", 1},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-12 13:14:00", 0},
		{"2020-08-05 13:14:15", "2020-08-12 13:14:15", 1},
		{"2020-08-05 13:14:15", "2020-08-12 13:14:59", 1},
	}

	for _, v := range Tests {
		output := Parse(v.input1).DiffInWeeksWithAbs(Parse(v.input2))

		if output != v.output {
			t.Errorf("Input start time %s and end time %s, expected %d, but got %d", v.input1, v.input2, v.output, output)
		}
	}
}

func TestCarbon_DiffInDays(t *testing.T) {
	Tests := []struct {
		input1 string // 输入值1
		input2 string // 输入值2
		output int64  // 期望输出值
	}{
		{"0000-00-00 00:00:00", "", 0},
		{"", "0000-00-00 00:00:00", 0},
		{"", "", 0},
		{"2020-08-05 13:14:15", "", -18479},
		{"", "2020-08-05 13:14:15", 18479},
		{"2020-08-05 13:14:15", "2020-08-04 13:00:00", -1},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:15", -1},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:00:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-06 13:00:00", 0},
		{"2020-08-05 13:14:15", "2020-08-06 13:14:15", 1},
		{"2020-08-05 13:14:15", "2020-08-06 13:14:59", 1},
	}

	for _, v := range Tests {
		output := Parse(v.input1).DiffInDays(Parse(v.input2))

		if output != v.output {
			t.Errorf("Input start time %s and end time %s, expected %d, but got %d", v.input1, v.input2, v.output, output)
		}
	}
}

func TestCarbon_DiffInDaysWithAbs(t *testing.T) {
	Tests := []struct {
		input1 string // 输入值1
		input2 string // 输入值2
		output int64  // 期望输出值
	}{
		{"0000-00-00 00:00:00", "", 0},
		{"", "0000-00-00 00:00:00", 0},
		{"", "", 0},
		{"2020-08-05 13:14:15", "", 18479},
		{"", "2020-08-05 13:14:15", 18479},
		{"2020-08-05 13:14:15", "2020-08-04 13:00:00", 1},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:15", 1},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:00:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-06 13:00:00", 0},
		{"2020-08-05 13:14:15", "2020-08-06 13:14:15", 1},
		{"2020-08-05 13:14:15", "2020-08-06 13:14:59", 1},
	}

	for _, v := range Tests {
		output := Parse(v.input1).DiffInDaysWithAbs(Parse(v.input2))

		if output != v.output {
			t.Errorf("Input start time %s and end time %s, expected %d, but got %d", v.input1, v.input2, v.output, output)
		}
	}
}

func TestCarbon_DiffInHours(t *testing.T) {
	Tests := []struct {
		input1 string // 输入值1
		input2 string // 输入值2
		output int64  // 期望输出值
	}{
		{"0000-00-00 00:00:00", "", 0},
		{"", "0000-00-00 00:00:00", 0},
		{"", "", 0},
		{"2020-08-05 13:14:15", "", -443501},
		{"", "2020-08-05 13:14:15", 443501},
		{"2020-08-05 13:14:15", "2020-08-05 12:14:00", -1},
		{"2020-08-05 13:14:15", "2020-08-05 12:14:15", -1},
		{"2020-08-05 13:14:15", "2020-08-05 12:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-05 14:14:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 14:14:15", 1},
		{"2020-08-05 13:14:15", "2020-08-05 14:14:59", 1},
	}

	for _, v := range Tests {
		output := Parse(v.input1).DiffInHours(Parse(v.input2))

		if output != v.output {
			t.Errorf("Input start time %s and end time %s, expected %d, but got %d", v.input1, v.input2, v.output, output)
		}
	}
}

func TestCarbon_DiffInHoursWithAbs(t *testing.T) {
	Tests := []struct {
		input1 string // 输入值1
		input2 string // 输入值2
		output int64  // 期望输出值
	}{
		{"0000-00-00 00:00:00", "", 0},
		{"", "0000-00-00 00:00:00", 0},
		{"", "", 0},
		{"2020-08-05 13:14:15", "", 443501},
		{"", "2020-08-05 13:14:15", 443501},
		{"2020-08-05 13:14:15", "2020-08-05 12:14:00", 1},
		{"2020-08-05 13:14:15", "2020-08-05 12:14:15", 1},
		{"2020-08-05 13:14:15", "2020-08-05 12:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-05 14:14:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 14:14:15", 1},
		{"2020-08-05 13:14:15", "2020-08-05 14:14:59", 1},
	}

	for _, v := range Tests {
		output := Parse(v.input1).DiffInHoursWithAbs(Parse(v.input2))

		if output != v.output {
			t.Errorf("Input start time %s and end time %s, expected %d, but got %d", v.input1, v.input2, v.output, output)
		}
	}
}

func TestCarbon_DiffInMinutes(t *testing.T) {
	Tests := []struct {
		input1 string // 输入值1
		input2 string // 输入值2
		output int64  // 期望输出值
	}{
		{"0000-00-00 00:00:00", "", 0},
		{"", "0000-00-00 00:00:00", 0},
		{"", "", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:13:00", -1},
		{"2020-08-05 13:14:15", "2020-08-05 13:13:15", -1},
		{"2020-08-05 13:14:15", "2020-08-05 13:13:59", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:15:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:15:15", 1},
		{"2020-08-05 13:14:15", "2020-08-05 13:15:59", 1},
		{"2020-08-05 13:14:15", "2020-08-05 13:16:00", 1},
		{"2020-08-05 13:14:15", "2020-08-05 13:16:15", 2},
		{"2020-08-05 13:14:15", "2020-08-05 13:16:59", 2},
		{"2020-08-05 13:14:15", "", -26610074},
		{"", "2010-08-05 13:14:15", 21349754},
	}

	for _, v := range Tests {
		output := Parse(v.input1).DiffInMinutes(Parse(v.input2))

		if output != v.output {
			t.Errorf("Input start time %s and end time %s, expected %d, but got %d", v.input1, v.input2, v.output, output)
		}
	}
}

func TestCarbon_DiffInMinutesWithAbs(t *testing.T) {
	Tests := []struct {
		input1 string // 输入值1
		input2 string // 输入值2
		output int64  // 期望输出值
	}{
		{"0000-00-00 00:00:00", "", 0},
		{"", "0000-00-00 00:00:00", 0},
		{"", "", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:13:00", 1},
		{"2020-08-05 13:14:15", "2020-08-05 13:13:15", 1},
		{"2020-08-05 13:14:15", "2020-08-05 13:13:59", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:15:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:15:15", 1},
		{"2020-08-05 13:14:15", "2020-08-05 13:15:59", 1},
		{"2020-08-05 13:14:15", "2020-08-05 13:16:00", 1},
		{"2020-08-05 13:14:15", "2020-08-05 13:16:15", 2},
		{"2020-08-05 13:14:15", "2020-08-05 13:16:59", 2},
		{"2020-08-05 13:14:15", "", 26610074},
		{"", "2010-08-05 13:14:15", 21349754},
	}

	for _, v := range Tests {
		output := Parse(v.input1).DiffInMinutesWithAbs(Parse(v.input2))

		if output != v.output {
			t.Errorf("Input start time %s and end time %s, expected %d, but got %d", v.input1, v.input2, v.output, output)
		}
	}
}

func TestCarbon_DiffInSeconds(t *testing.T) {
	Tests := []struct {
		input1 string // 输入值1
		input2 string // 输入值2
		output int64  // 期望输出值
	}{
		{"0000-00-00 00:00:00", "", 0},
		{"", "0000-00-00 00:00:00", 0},
		{"", "", 0},
		{"2020-08-05 13:14:15", "", -1596604455},
		{"", "2010-08-05 13:14:15", 1280985255},
		{"2020-08-05 13:14:15", "2010-08-05 13:14:15", -315619200},
	}

	for _, v := range Tests {
		output := Parse(v.input1).DiffInSeconds(Parse(v.input2))

		if output != v.output {
			t.Errorf("Input start time %s and end time %s, expected %d, but got %d", v.input1, v.input2, v.output, output)
		}
	}
}

func TestCarbon_DiffInSecondsWithAbs(t *testing.T) {
	Tests := []struct {
		input1 string // 输入值1
		input2 string // 输入值2
		output int64  // 期望输出值
	}{
		{"0000-00-00 00:00:00", "", 0},
		{"", "0000-00-00 00:00:00", 0},
		{"", "", 0},
		{"2020-08-05 13:14:15", "", 1596604455},
		{"", "2010-08-05 13:14:15", 1280985255},
		{"2020-08-05 13:14:15", "2010-08-05 13:14:15", 315619200},
	}

	for _, v := range Tests {
		output := Parse(v.input1).DiffInSecondsWithAbs(Parse(v.input2))

		if output != v.output {
			t.Errorf("Input start time %s and end time %s, expected %d, but got %d", v.input1, v.input2, v.output, output)
		}
	}
}

func TestCarbon_DiffForHumans1(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output string // 期望输出值
	}{
		{Now(), "just now"},
		{Now().AddYears(1), "1 year from now"},
		{Now().SubYears(1), "1 year ago"},
		{Now().AddYears(10), "10 years from now"},
		{Now().SubYears(10), "10 years ago"},

		{Now().AddMonths(1), "1 month from now"},
		{Now().SubMonths(1), "1 month ago"},
		{Now().AddMonths(10), "10 months from now"},
		{Now().SubMonths(10), "10 months ago"},

		{Now().AddDays(1), "1 day from now"},
		{Now().SubDays(1), "1 day ago"},
		{Now().AddDays(10), "1 week from now"},
		{Now().SubDays(10), "1 week ago"},

		{Now().AddHours(1), "1 hour from now"},
		{Now().SubHours(1), "1 hour ago"},
		{Now().AddHours(10), "10 hours from now"},
		{Now().SubHours(10), "10 hours ago"},

		{Now().AddMinutes(1), "1 minute from now"},
		{Now().SubMinutes(1), "1 minute ago"},
		{Now().AddMinutes(10), "10 minutes from now"},
		{Now().SubMinutes(10), "10 minutes ago"},

		{Now().AddSeconds(1), "1 second from now"},
		{Now().SubSeconds(1), "1 second ago"},
		{Now().AddSeconds(10), "10 seconds from now"},
		{Now().SubSeconds(10), "10 seconds ago"},
	}

	for _, v := range Tests {
		output := (v.input).DiffForHumans()

		if output != v.output {
			t.Errorf("Input time %s, expected %s, but got %s", v.input.ToDateTimeString(), v.output, output)
		}
	}
}

func TestCarbon_DiffForHumans2(t *testing.T) {
	Tests := []struct {
		input1 Carbon // 输入值1
		input2 Carbon // 输入值2
		output string // 期望输出值
	}{
		{Now(), Now(), "刚刚"},
		{Now().AddYears(1), Now(), "1 年后"},
		{Now().SubYears(1), Now(), "1 年前"},
		{Now().AddYears(10), Now(), "10 年后"},
		{Now().SubYears(10), Now(), "10 年前"},

		{Now().AddMonths(1), Now(), "1 个月后"},
		{Now().SubMonths(1), Now(), "1 个月前"},
		{Now().AddMonths(10), Now(), "10 个月后"},
		{Now().SubMonths(10), Now(), "10 个月前"},
	}

	for _, v := range Tests {
		output := (v.input1).SetLocale("zh-CN").DiffForHumans(v.input2)

		if output != v.output {
			t.Errorf("Input time %s, expected %s, but got %s", v.input1.ToDateTimeString(), v.output, output)
		}
	}
}