package languagepack

var English = map[string]string{
	"filenotfound":     "File %s not found.",
	"filecannotberead": "File %s cannot be read.",

	"info.fileread":   "File read in %d microseconds.",
	"info.fileparsed": "File parsed in %d microseconds.",

	"error.nomain": "File, that you tried to compile doesn't have `main` module.",
	"note.nomain":  "Change module name from `%s` to `main`.",

	"error.nomainfunc": "File, that you tried to compile doesn't have `main` function.",
	"note.nomainfunc":  `Just add main function main.`,

	"info.foundmain":     "Found module `main`",
	"info.foundmainfunc": "Found function `main`",
}
