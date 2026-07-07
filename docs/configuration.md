# How to configure uMotd

uMotd has a very straightforward config file.

## Where to put your config

You can create and edit the config file at `/etc/ublue-os/tags.json`.

## Translations ?

That's the point !
It's made to support translation ! - [Translation Guide](translating.md)

## Breaking down the configuration file

Here's a small breakdown of the config file options - there's also the example folder if you want to see concrete use cases.

### Tags

The `tags` option defines which translated messages to show when running uMotd.
They regroup thematic messages like the desktop environment or what kind of dev tools are used on your machine.

Currently, the following tags are available:

- `aurora`
- `bazzite`
- `bazzite-deck`
- `bluefin`
- `gnome`
- `kde`
- `vscode`
- `containers`

```json
{
  "tags": [
    "gnome",
    "bluefin",
    "vscode",
    "containers"
  ]
}
```

#### Commands

You can also edit the tags with the following commands :

```sh
umotd tags add <tag>...    # For adding tags
umotd tags remove <tag>... # For removing tags
umotd tags list            # For listing all of your tags
```
