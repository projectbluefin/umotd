# How to add new messages

(WIP)

First of all, you can edit messages wherever you like, from Github directly or from your favorite code editor.
You won't necessarily need go installed.

Then, you'll need to open `/internal/messages` - that's where the magic happens.
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

## Tags

The tags are defined simply by if statements, like this : 
```go
if slices.Contains(tags, "gnome") {
}
```

For this example, it's simply checking if there is "gnome" in the defined tags of the configuration.  

It then appends the strings related to the tag
```
messages = append(messages, []string{}[...])
```