version := `go run . version`

# This returns the current version of uMotd. It is used in the build command to set the version of the binary.
version :
    @echo "Current version: {{version}}"

# This builds the umotd binary for Linux on x86_64 and arm64 architectures.
build :
    @echo "Building umotd..."

    GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o umotd_{{version}}_linux_amd64
    GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o umotd_{{version}}_linux_arm64

    @echo "Built umotd v{{version}} for Linux on x86_64 and arm64"

# This generates the translation files for the specified language.
translate lang:
    #!/bin/bash
    # Translators script for uMotd

    # Checking dependencies

    echo "Checking dependencies..."
    which go >/dev/null 2>&1 || (echo "You don't have \`go\` installed :(" && exit 1)
    which gettext >/dev/null 2>&1 || (echo "You don't have \`gettext\` installed :(" && exit 1)
    [ -f "$(go env GOPATH)/bin/xgotext" ] || (go install github.com/leonelquinteros/gotext/cli/xgotext@latest || (echo "An error occurred while installing \`xgotext\` :(" && exit 1))
    echo "All dependencies are installed ✅️"

    # If the language already exists, update it

    if [ -f "locales/{{lang}}/default.po" ]; then \
        echo "Translation file for \"{{lang}}\" already exists. Updating..." ;\
        $(go env GOPATH)/bin/xgotext -in . -out locales/temp || (echo "An error occurred while running \`xgotext\` :(" && rm -rf locales/temp && exit 1) ;\
        msgmerge --update locales/"{{lang}}"/default.po locales/temp/default.pot || (echo "An error occurred while running \`msgmerge\` :(" && rm -rf locales/temp && exit 1);\
        rm -rf locales/temp;\
        rm -f locales/"{{lang}}"/default.po~;\
        echo "File updated! You can edit it in locales/{{lang}}/default.po";\
        exit 0;\
    fi

    # If the language does not exist, create it

    echo "Translation file for \"{{lang}}\" do not exist. Creating new file..."
    mkdir -p locales/"{{lang}}"
    $(go env GOPATH)/bin/xgotext -in . -out locales/temp || (echo "An error occurred while running \`xgotext\` :(" && rm -rf locales/temp && exit 1)
    cp locales/temp/default.pot locales/"{{lang}}"/default.po
    rm -rf locales/temp
    echo "Translation file generated. You can edit them in locales/{{lang}}/default.po"

