Editor file check tool checks practice exercises for files missing from the editor.
For each practice exercise, the `.meta/config.json` is loaded.
The `files.editor` list is compared to the `.go` files in the exercise directory.
If the lists differ, the tool complains about them.

If a file should not be included for some reason, it can be excluded via the `config.json`.
