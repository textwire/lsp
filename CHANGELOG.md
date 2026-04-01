# Release Notes

## v0.5.0 (2026-04-01)

- ✨ Update Textwire to the latest version.

## v0.4.2 (2026-03-01)

- ✨ Update Textwire to the latest version.

## v0.4.1 (2026-02-14)

- 🐛 Fixed bug with infinite loop that whould hung your editor when you had some syntax issues. Bug is fixed by updating Textwire to the latest version.

## v0.4.0 (2025-11-08)

- ✨ Loop directives like `@break`, `@continue`, `@breakIf`, `@continueIf` are now getting suggested only inside of loops.
- 🧑‍💻 Added testing GitHub Actions job to run tests.

## v0.3.1 (2025-06-14)

- 🐛 Updated LSP textwire to the latest version which will fix `@break($1)` snippet to `@break`

## v0.3.0 (2025-06-14)

- 🐛 Updated LSP textwire to the latest version which will fix crashing LSP when you make a syntax error in your Textwire code
- ✨ Autocomplete snippets that appear after you hit enter are now more complex. Instead of simple autocomple like `@if` you know get the full if statement and the cursor placed inside condition. It allows you to hit tab to move to the next place in a snippet

## v0.2.0 (2025-05-30)

- ✨ Added autocompletion for loop object. Now if you type `loop.` inside of a loop, it will show available properties on the object

## v0.1.4 (2025-05-17)

- ✨ Autocomplete suggestions show code example with syntax highlighting. Before, it was just displayed as text

## v0.1.3 (2025-05-16)

- 🐛 Fixed proper autocomplete logic

## v0.1.2 (2025-05-15)

- 🐛 Fixed bug with autocompletion not working properly on vscode editor

## v0.1.1 (2025-05-12)

- 🐛 Fixed bug with autocompletion not working properly in some cases

## v0.1.0 (2025-05-03)

- ✨ Added `build.yml` file for GitHub actions to build releases with all LSP binaries

## v0.0.2 (2025-04-25)

- ✨ Improved logger
- ✨ Added generating `build-version` file to `bin` directory

## v0.0.1 (2025-04-16)

- ✨ Added completions for directives
- ✨ Added showing hover information for directives
