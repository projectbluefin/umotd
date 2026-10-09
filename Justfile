package := "umotd"
version := "0.3.3"

# This returns the current version of umotd. It is used in the build command to set the version of the binary.
version :
    @echo "Current version: {{version}}"

# This builds the umotd binary for Linux on x86_64 and arm64 architectures.
build :
    @echo "Building {{package}}..."

    GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o {{package}}_{{version}}_linux_amd64
    GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o {{package}}_{{version}}_linux_arm64

    @echo "Built {{package}} v{{version}} for Linux on x86_64 and arm64"

# This checks if the dependencies are on the user's system
check-deps: check-deps-go check-deps-gettext
    @echo "All dependencies are installed ✅️"

check-deps-go:
    @which go >/dev/null 2>&1 || (echo "You don't have \`go\` installed :(" && exit 1)

check-deps-gettext:
    @which gettext >/dev/null 2>&1 || (echo "You don't have \`gettext\` installed :(" && exit 1)

# This extracts the strings from umotd and makes a potfile with them
extract-translations: check-deps-gettext
    #!/bin/bash
    echo "Extracting translations..."
    find . -type f -name '*.go' -not -path './vendor/*' -print0 |
        xargs -0 xgettext \
            --language=Go \
            --from-code=UTF-8 \
            --keyword=Get \
            --msgid-bugs-address=themimolet@proton.me \
            --copyright-holder="The Project Bluefin Contributors" \
            --package-version="{{version}}" \
            --package-name="{{package}}" \
            --output=locales/default.pot \
            || (echo "An error occurred while extracting translations :(" \
            && exit 1);
    echo "Translations extracted ✅️"

# This generates the translation files for the specified language.
translate lang: check-deps-go extract-translations
    #!/bin/bash
    # If the language already exists, update it
    if [ -f "locales/{{lang}}/default.po" ]; then \
        echo "Translation file for \"{{lang}}\" already exists. Updating..." ;\
        msgmerge -U locales/"{{lang}}"/default.po locales/default.pot || (echo "An error occurred while running \`msgmerge\` :(" && exit 1);\
        rm -f locales/"{{lang}}"/default.po~;\
        echo "File updated! You can edit it in locales/{{lang}}/default.po";\
        exit 0;\
    fi
    # If the language does not exist, create it
    echo "Translation file for \"{{lang}}\" do not exist. Creating new file..."
    mkdir -p locales/"{{lang}}"/
    msginit -i locales/default.pot -l "{{lang}}" -o locales/"{{lang}}"/default.po --no-translator
    echo "Translation file generated. You can edit them in locales/{{lang}}/default.po"