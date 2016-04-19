package entries

import (
	"bufio"
	"errors"
	"github.com/kcmerrill/MrT/entry"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var entries map[int][]*entry.Entry
var sorted_priorities []int
var added []*entry.Entry
var results []*entry.Entry

func Update() {
	Create()
	f, _ := os.Open(viper.GetString("tasks"))
	scanner := bufio.NewScanner(f)
	/* Go through and update our file */
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Trim(line, "\n") == "" {
			continue
		}
		e := entry.Parse(line)
		entries[e.Score()] = append(entries[e.Score()], e)
		if e.IsNew() {
			added = append(added, e)
		}
	}
	f.Close()
}

func List(show int) {
	for _, entry_score := range Sorted() {
		for _, entry := range entries[entry_score] {
			if entry.IsCompleted() {
				continue
			}
			results = append(results, entry)
			show--
			if show == 0 && show != -1 {
				return
			}
		}
	}
}

func Sorted() []int {
	if len(sorted_priorities) > 0 {
		return sorted_priorities
	}
	keys := make([]int, len(entries))
	i := 0
	for k, _ := range entries {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	sorted_priorities = keys
	return sorted_priorities
}

func All() []*entry.Entry {
	return results
}

func Added() []*entry.Entry {
	return added
}

func Undo() error {
	/* Make sure our backup file exists first */
	if _, err := os.Stat(viper.GetString("tasks_backup")); os.IsNotExist(err) {
		return errors.New("Unable to restore from backup.")
	}
	if err := os.Rename(viper.GetString("tasks_backup"), viper.GetString("tasks")); err != nil {
		return errors.New("Unable to restore from backup")
	}
	return nil
}

func Save() error {
	/* Lets first create the backup */
	if err := os.Rename(viper.GetString("tasks"), viper.GetString("tasks_backup")); err != nil {
		return errors.New("Unable to create backup of your tasks.")
	}

	if f, err := os.Create(viper.GetString("tasks")); err == nil {
		for _, entry_score := range Sorted() {
			for _, entry := range entries[entry_score] {
				f.WriteString(entry.ToString() + "\n")
			}
		}
	} else {
		return errors.New("Unable to save your tasks.")
	}
	return nil
}

func Create() {
	if _, err := os.Stat(viper.GetString("tasks")); os.IsNotExist(err) {
		/* Just need to make sure the directory exists */
		os.MkdirAll(filepath.Dir(viper.GetString("tasks")), 0644)
	}
}

func Get(index int) (*entry.Entry, error) {
	if _, exists := entries[index]; exists {
		return results[index], nil
	}
	return nil, errors.New("Unable to find task.")
}

func Complete(task_id int) (*entry.Entry, error) {
	List(-1)
	if len(results) > task_id {
		results[task_id].Complete()
		return results[task_id], nil
	} else {
		return &entry.Entry{}, errors.New("Unable to find task.")
	}

	return &entry.Entry{}, nil
}

func init() {
	entries = make(map[int][]*entry.Entry)
}
