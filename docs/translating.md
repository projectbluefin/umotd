# How to translate

## Prerequisites

To translate uMotd, you'll need to have [`gettext`](https://repology.org/project/gettext/versions) installed on your system.
If you're on a Universal Blue system, it is already installed on your system.

## Usage

You can simply run `just translate <language code>` to extract the translatable strings and update the translations file.

Example:

```sh
just translate fr
```

Your generated/updated translation file will be located in the `locales/<language code>` directory.

> If your language already exists, it will get updated automatically. If not, a new translation file will be created for you.

Finally, use your favorite po editor to translate the strings in the `.po` file - like [Poedit](https://flathub.org/en/apps/net.poedit.Poedit), [Gtranslator](https://flathub.org/en/apps/org.gnome.Gtranslator) or [Lokalize](https://flathub.org/en/apps/org.kde.lokalize).

You can also use a simple text editor if you wish.

## Testing your translation

To test umotd, you will need to have [`go`](https://repology.org/project/go/versions) installed.

Then, to build the program, open the terminal in the project folder.
After, you can run the project with the following:

```sh
go run .
```

It will run in the language used by your system (or english if umotd doesn't have any translation yet).

If you want to make sure the program run in a specific language, you can use `LANGUAGE=<language code>` in front of the usual command to test your translation, like this:

```sh
LANGUAGE=fr go run .
```
