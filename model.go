package main

type Post struct {
	ID       int
	Title    string
	Content  string
	Comments []Comment
}

type Comment struct {
	ID      int
	PostID  int
	Content string
	Author  string
}
