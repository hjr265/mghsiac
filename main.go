package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/shurcooL/githubv4"
	"github.com/shurcooL/graphql"
	"golang.org/x/oauth2"
)

var (
	flagTZ = flag.String("tz", "", "timezone")
)

func main() {
	flag.Parse()

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpclient := oauth2.NewClient(context.Background(), src)
	client := githubv4.NewClient(httpclient)

	for {
		now := time.Now()
		if *flagTZ != "" {
			loc, err := time.LoadLocation(*flagTZ)
			if err != nil {
				log.Fatalf("Could not load timezone data: %s", err)
			}
			now = now.In(loc)
		}

		emoji := timeToEmoji(now)
		message := timeToMessage(now)
		log.Printf("Updating status to %s (%s)", emoji, message)

		m := struct {
			ChangeUserStatus struct {
				Status struct {
					Emoji   graphql.String
					Message graphql.String
				}
			} `graphql:"changeUserStatus(input: $input)"`
		}{}
		err := client.Mutate(context.Background(), &m, githubv4.ChangeUserStatusInput{
			Emoji:   githubv4.NewString(githubv4.String(emoji)),
			Message: githubv4.NewString(githubv4.String(message)),
		}, nil)
		if err != nil {
			log.Printf("Could not update user status: %s", err)
		}

		d := sleepDuration(now)
		log.Printf("Sleeping for %s", d)
		time.Sleep(d)
	}
}

func timeToEmoji(t time.Time) string {
	t = t.Add(10 * time.Minute)
	h := t.Hour()
	if h > 12 {
		h -= 12
	}
	if h == 0 {
		h = 12
	}
	m := t.Minute()
	m = (m / 30) * 30
	if m == 0 {
		return fmt.Sprintf(":clock%d:", h)
	}
	return fmt.Sprintf(":clock%d%d:", h, m)
}

var hourWord = map[int]string{
	1:  "One",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
	11: "Eleven",
	12: "Twelve",
}

func timeToMessage(t time.Time) string {
	t = t.Add(10 * time.Minute)
	h := t.Hour()
	if h > 12 {
		h -= 12
	}
	if h == 0 {
		h = 12
	}
	m := t.Minute()
	m = (m / 30) * 30
	if m == 0 {
		return fmt.Sprintf("%s o'clock", hourWord[h])
	}
	return fmt.Sprintf("Half past %s", strings.ToLower(hourWord[h]))
}

func sleepDuration(t time.Time) time.Duration {
	t = t.Add(10 * time.Minute)
	m := t.Minute()
	if m < 30 {
		return time.Duration(30-m) * time.Minute
	}
	return time.Duration(60-m) * time.Minute
}
