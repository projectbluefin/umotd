# uMotd

*uMotd is the translatable MOTD for Universal Blue!*

**WIP** : Some features still need testing and strings need to be checked.

**Contributions are welcome!** If you want to contribute, you're welcome to submit a pull request or open an issue - it's very much appreciated ❤️

[Access the documentation](https://github.com/theMimolet/umotd/tree/main/docs)

## How to try

<!-- ### Universal Blue system images

If you're running Bluefin Testing, it's already included in the image itself ! -->

### Install it with Homebrew

If you have Homebrew installed on your system, you can install Umotd with the following command:
```
brew install themimolet/tap/umotd
```

### Download it from the releases page

You can download the binaries from the [releases page](https://github.com/theMimolet/umotd/releases).

You can then rename it to `umotd` and place it in your usual `/bin` folder.

> Note: You won't receive automatic updates, you will have to download Umotd at each new release.

### Compile from source

You'll need to have [`go`](https://repology.org/project/go/versions) installed on your system to compile Umotd from source.

Then you'll have to simply clone the repository and then build the binary:
```
git clone https://github.com/theMimolet/umotd
cd umotd
go build
./umotd
```

You'll then have the `umotd` binary in the current directory, which you can just drop into your usual `/bin` folder and it will work without any further setup (except for the configuration file if you want to customize it).

## Commands

Umotd supports the following commands:

- `toggle`: toggles the MOTD on or off for the current user
- `enable`: always enables the MOTD for the current user
- `disable`: always disables the MOTD for the current user
- `version`, `--version`, `-v`: displays the version of Umotd you're currently using

- `tags list`: lists all available tags
- `tags add <tag>`: adds one or multiple new tags to the config file
- `tags remove <tag>`: removes one or multiple tags from the config file
