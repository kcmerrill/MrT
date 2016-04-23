package display

import (
	"fmt"
	"github.com/kcmerrill/MrT/entries"
	"github.com/kcmerrill/MrT/entry"
	"github.com/olekukonko/tablewriter"
	"os"
	"strings"
)

func Init() {
	fmt.Println("Tasks initialized.")
}

func Added() {
	added := entries.Added()
	for _, e := range added {
		fmt.Println(e.Description())
	}

	count := len(added)
	switch count {
	case 0:
		fmt.Println("No tasks were added.")
		break
	case 1:
		fmt.Println("---")
		fmt.Println("1 task added.")
		break
	default:
		fmt.Println("---")
		fmt.Println(count, "tasks added.")
	}
	entries.Save()
}

func Current() {
	e := entries.All()
	if len(e) >= 1 {
		fmt.Println(e[0].Description())
	} else {
		fmt.Println("No tasks.")
	}
	entries.Save()
}

func LS(fields []string) {
	ls := entries.All()
	if len(ls) == 0 {
		fmt.Println("No tasks.")
		return
	}
	default_fields := []string{"ID", "Description"}
	if fields == nil {
		fields = default_fields
	}
	table := tablewriter.NewWriter(os.Stdout)
	header := make([]string, 0, len(fields))
	for _, h := range fields {
		header = append(header, h)
	}
	table.SetHeader(header)
	table.SetBorder(false)
	table.SetColWidth(200)
	for id, e := range ls {
		row := make([]string, 0, len(fields))
		for _, m := range fields {
			switch strings.ToLower(m) {
			case "id":
				row = append(row, fmt.Sprintf("%d", id))
				break
			case "description":
				row = append(row, e.Description())
				break
			default:
				row = append(row, e.DisplayMeta(strings.ToLower(m), ""))
			}
		}
		table.Append(row)
	}
	table.Render()
	entries.Save()
}

func Complete(id int, e *entry.Entry) {
	fmt.Println(e.Description())
	fmt.Println("---")
	fmt.Println(fmt.Sprintf("Task #%d completed.", id))
	entries.Save()
}

func Start(id int, e *entry.Entry) {
	fmt.Println(e.Description())
	fmt.Println("---")
	fmt.Println(fmt.Sprintf("Task #%d started.", id))
	entries.Save()
}

func Error(msg string) {
	fmt.Println(msg)
}

func Undo() {
	fmt.Println("Undo succesful.")
}

func init() {
	fmt.Println("---")
}
