package entry

import (
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	raw         string
	description []string
	meta        map[string][]string
	parsed      bool
	score       int
	is_new      bool
}

func New(raw string) *Entry {
	e := &Entry{
		raw: strings.Trim(raw, " \n"),
	}
	e.meta = make(map[string][]string)
	e.parse()
	return e
}

func Parse(raw string) *Entry {
	return New(raw)
}

func (e *Entry) parse() {
	for _, token := range strings.Split(e.raw, " ") {
		if token == "" {
			continue
		}
		switch {
		case string(token[0]) == "#":
			e.meta["lists"] = append(e.meta["lists"], token[1:])
			break
		case strings.Contains(token, ":"):
			meta := strings.SplitN(token, ":", 2)
			e.meta[meta[0]] = append(e.meta[meta[0]], meta[1])
			break
		default:
			e.description = append(e.description, token)
		}
	}
	e.Created()
	e.score = e.CalculateScore()
	e.Parsed()
}

func (e *Entry) CalculateScore() int {
	priority := e.Priority()
	/* Give higher priorities a solid default score */
	score := priority * 1000
	/* Depending if it's past due, add accordingly */
	if hours := time.Since(e.Due()).Hours(); hours >= 0 {
		/* Ugh. Past due ... */
		score = priority * 10
	} else {
		score += int(hours) * -1
	}
	return score
}

func (e *Entry) Meta() map[string][]string {
	return e.meta
}

func (e *Entry) Score() int {
	return e.score
}

func (e *Entry) Due() time.Time {
	due := time.Now().AddDate(0, 0, 30)
	if _, exists := e.meta["due"]; exists {
		if t, err := time.Parse(viper.GetString("date_format"), e.meta["due"][0]); err == nil {
			due = t
		} else {
			delete(e.meta, "due")
		}
	}
	return due
}

func (e *Entry) Created() string {
	created := time.Now()
	if _, exists := e.meta["created"]; exists {
		if t, err := time.Parse(viper.GetString("date_format"), e.meta["created"][0]); err == nil {
			created = t
		}
	} else {
		e.is_new = true
	}
	e.meta["created"] = []string{created.Format(viper.GetString("date_format"))}
	return e.meta["created"][0]
}

func (e *Entry) Start() string {
	started := time.Now().Format(viper.GetString("date_format"))
	e.SetMeta("started", started)
	return started
}

func (e *Entry) Complete() string {
	completed := time.Now().Format(viper.GetString("date_format"))
	e.SetMeta("completed", completed)
	return completed
}

func (e *Entry) SetMeta(key, value string) string {
	e.meta[key] = []string{value}
	return value
}

func (e *Entry) IsCompleted() bool {
	return e.HasMeta("completed")
}

func (e *Entry) HasStarted() bool {
	return e.HasMeta("started")
}

func (e *Entry) HasMeta(key string) bool {
	_, has := e.meta[key]
	return has
}

func (e *Entry) Parsed() {
	e.parsed = true
}
func (e *Entry) IsParsed() bool {
	return e.parsed
}

func (e *Entry) Priority() int {
	/* Are you completed? If so ... fake a high priority(bottom of list) */
	if e.IsCompleted() {
		return 10000
	}

	/* Clearly you're not done, but have you started? Shoot to the top! */
	if e.HasStarted() {
		return 0
	}

	if _, exists := e.meta["priority"]; exists {
		switch e.meta["priority"][0] {
		case "high":
			e.meta["priority"] = []string{"1"}
			break
		case "medium":
			e.meta["priority"] = []string{"5"}
			break
		case "low":
			e.meta["priority"] = []string{"10"}
			break
		}
	} else {
		e.meta["priority"] = []string{"10"}
	}
	if priority, err := strconv.Atoi(e.meta["priority"][0]); err == nil {
		return priority
	} else {
		/* Wut?! Broken priority ... */
		e.meta["priority"] = []string{"10"}
		return 10
	}
}

func (e *Entry) Description() string {
	return strings.Join(e.description, " ") + " "
}

func (e *Entry) IsNew() bool {
	return e.is_new
}
func (e *Entry) ToString() string {
	entry := e.Description() + " "
	for key, values := range e.meta {
		for _, value := range values {
			entry += key + ":" + value + " "
		}
	}
	return entry
}
