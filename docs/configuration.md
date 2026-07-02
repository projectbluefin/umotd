# How to configure Umotd

Umotd has a very straightforward config file.

## Where to put your config

You can create and edit the config file at `/etc/ublue-os/config.json`.

## Translations ?!

That's the point ! 
It's made to support translation ! - [Translation Guide](translating.md)

## Breaking down the configuration file

Here's a small breakdown of the config file options - there's also the example folder if you want to see concrete use cases.

### Tags

The `tags` option defines which translated messages to show when running umotd.
They regroup thematic messages like the desktop envirronment or what kind of dev tools is used on your machine.

Currently, there are the following tags available:

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
