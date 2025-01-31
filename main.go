package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
)

var (
	begin_vcalendar = `BEGIN:VCALENDAR
VERSION:2.0
PRODID:-//icalgen//NONSGML//EN`
	end_vcalendar = `END:VCALENDAR`
	vevent        = `BEGIN:VEVENT
DTSTART;VALUE=DATE:{{.Date}}
DTEND;VALUE=DATE:{{.Date}}
SUMMARY:{{.Summary}}
LOCATION:n/a
DESCRIPTION:n/a
BEGIN:VALARM
TRIGGER:-P0DT7H10M0S
ACTION:DISPLAY
DESCRIPTION:{{.Summary}}
END:VALARM
BEGIN:VALARM
TRIGGER:-P1DT6H00M0S
ACTION:DISPLAY
DESCRIPTION:{{.Summary}}
END:VALARM
END:VEVENT
`
)

func print_vevent(year int, month int, day int, title string) {
	templ, err := template.New("vevent").Parse(vevent)
	if err != nil {
		panic(err)
	}
	templ.Execute(os.Stdout, map[string]string{
		"Date":    fmt.Sprintf("%04d%02d%02d", year, month, day),
		"Summary": title,
	})
}

func main() {

	year := flag.Int("year", time.Now().Year(), "year for events")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	title := ""
	vcal_printed := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 && line[0] != '#' {
			if strings.HasPrefix(line, "title:") {
				new_title, _ := strings.CutPrefix(line, "title:")
				title = strings.TrimSpace(new_title)
			} else {
				param := strings.Split(line, " ")
				month, err := strconv.Atoi(param[0])
				if err != nil {
					fmt.Fprintf(os.Stderr, "# Ignore line: %v, error %v\n", line, err)
					continue
				}
				for i := 1; i < len(param); i++ {
					day, err := strconv.Atoi(param[i])
					if err != nil {
						fmt.Fprintf(os.Stderr, "# Ignore value %v, error %v\n", param[i], err)
						break
					}
					if !vcal_printed {
						fmt.Println(begin_vcalendar)
						vcal_printed = true
					}
					print_vevent(*year, month, day, title)
				}
			}
		}
	}

	if vcal_printed {
		fmt.Println(end_vcalendar)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "read stdin error: %v", err)
	}
}
