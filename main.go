package main

import (
    "fmt"
    "os"
	"github.com/umisetokikaze/blog-idea-manager/database"
)

func addIdea(title, content string) error {
    _, err := database.DB.Exec("INSERT INTO ideas (title, content) VALUES (?, ?)", title, content)
    return err
}

func listIdeas() error {
    rows, err := database.DB.Query("SELECT id, title, created_at FROM ideas")
    if err != nil {
        return err
    }
    defer rows.Close()

    for rows.Next() {
        var id int
        var title string
        var createdAt string
        err = rows.Scan(&id, &title, &createdAt)
        if err != nil {
            return err
        }
        fmt.Printf("%d: %s (%s)\n", id, title, createdAt)
    }
    return nil
}


func main() {
    database.Init()

    if len(os.Args) < 2 {
        fmt.Println("コマンドが不足しています。'help'で使用可能なコマンドを確認してください。")
        return
    }

    switch os.Args[1] {
    case "help":
        fmt.Println("使用可能なコマンド: help, add, list, edit, delete")
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("タイトルを指定してください。")
			return
		}
		title := os.Args[2]
		content := ""
		if len(os.Args) > 3 {
			content = os.Args[3]
		}
		err := addIdea(title, content)
		if err != nil {
			fmt.Println("アイデアの追加に失敗しました:", err)
		} else {
			fmt.Println("アイデアを追加しました。")
		}
	case "list":
		err := listIdeas()
		if err != nil {
			fmt.Println("アイデアの取得に失敗しました:", err)
		}

    default:
        fmt.Println("不明なコマンドです。'help'で使用可能なコマンドを確認してください。")
    }
}
