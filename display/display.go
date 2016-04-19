package display

import (
	"fmt"
	"github.com/kcmerrill/MrT/entries"
	"github.com/kcmerrill/MrT/entry"
	"github.com/olekukonko/tablewriter"
	"os"
)

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

func LS() {
	ls := entries.All()
	if len(ls) == 0 {
		fmt.Println("No tasks.")
		return
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Description"})
	table.SetBorder(false)
	table.SetColWidth(200)
	for id, e := range ls {
		table.Append([]string{fmt.Sprintf("%d", id), e.Description()})
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

func Error(msg string) {
	fmt.Println(msg)
}

func Undo() {
	fmt.Println("Undo succesful.")
}

func init() {
	fmt.Println("---")
}
