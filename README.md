# MGHSIAC

MGHSIAC (an elegant abbreviation of My GitHub Status Is A Clock) is a utility to turn your GitHub status emoji into an almost functional clock.

It turned [my GitHub status into a clock](https://github.com/hjr265).

![](screen.png)

## Usage

Create a classic personal access token with the "user" scope for your GitHub account and run `mghsiac` with it.

```
GITHUB_TOKEN=... ./mghsiac
```

MGHSIAC, when run, will update the user status emoji and then sleep until the next update is necessary.

Leave it running in the background to turn your GitHub status emoji into a clock.

If you need to specify a timezone for your clock, you can use the `-tz` flag (e.g. `-tz=Asia/Dhaka`).

## Why?

What else would I set as my GitHub status emoji? :dart:?

Please.

## How It Works

I wrote a [brief post on how this works](https://hjr265.me/blog/my-github-status-is-a-clock/) on my blog.
