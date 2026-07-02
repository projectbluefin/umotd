
# How to translate

## Prerequisites

To translate and test Umotd, you'll need to have the following tools installed on your system:

- [`go`](https://repology.org/project/go/versions)
- [`gettext`](https://repology.org/project/gettext/versions)
- [`xgotext`](https://pkg.go.dev/github.com/leonelquinteros/gotext/cli/xgotext) (it should be installed automatically on the first run of the `translators.sh` script)

## Usage

You can simply run the translators.sh script to extract the translatable strings and update the translations files.

```
./translators.sh <language code>
```

Your translation files are located in the `locales/<language code>/LC_MESSAGES/default.po` directory.

> If your language already exists, it will be updated automatically. If not, a new language file will be created for you.

Finally, use your favorite po editor to translate the strings in the `.po` file - like [Poedit](appstream://net.poedit.poedit), [Gtranslator](appstream://org.gnome.Gtranslator) or [Lokalize](appstream://org.kde.lokalize).

## Testing your translation

You can then use `LANGUAGE=<language code>` in front of the usual command to test your translation, like this:

```sh
# Run with the compiled binary - needs to be rebuilt after translation changes
LANGUAGE=fr ./umotd
```

```sh
# Run with the source code - not compiled, so no need to rebuild after translation changes
LANGUAGE=fr go run .
```
