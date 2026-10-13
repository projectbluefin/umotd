# How to add new messages

First of all, you can edit messages wherever you like, from Github directly or from your favorite code editor.
You won't necessarily need go installed.

Then, you'll need to open `/internal/messages.go` - that's where the magic happens.
There will be a lot, but you don't need to understand everything.

Here what you need to know.

## Messages

A translatable message looks like like this :

```go
l.Get(" **Want to install Decky Loader?** There's a `ujust` command for that! `ujust setup-decky install`"),
```

Between the `(" ")` is our message we want to show and translate.
`l.Get` is the command used by go to get the translation of our message.

It's also possible to "inject" infos into it :

```go
l.Get(" **%s is your gateway to Cloud Native** — find your flock at [landscape.cncf.io](%s)", GetOSName(), "https://l.cncf.io"),
```

`%s` represents what we inject into it (a string for that matter) and they are handled in order.
Here the first `%s` adds the OS Name and the second `%s` adds the link.

## Conditions

What is printed out depends on conditions - it will check different values accross the system and will print the strings related to what it finds.

In practice, those are defined simply by `if` statements, like this one:

```go
if Desktop == "GNOME" {
    ...
}
```

For the following example, it's checking if the current desktop is "KDE", then it appends the strings related to that desktop :

```go
if Desktop == "KDE" {
    messages = append(messages, []string{
        l.Get("Hi, if you see this, you use Kde!"),
    }[...])
}

```
