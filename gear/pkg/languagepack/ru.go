package languagepack

var Russian = map[string]string{
	"error.filenotfound":     "Файл %s не найден.",
	"error.filecannotberead": "Файл %s не может быть прочитан.",

	"info.fileread":   "Файл считался за %d микросекунд",
	"info.fileparsed": "Файл спарсился за %d микросекунд",

	"error.nomain": "Файл, который вы пытаетесь скомпилировать не содержит модуля `main`.",
	"note.nomain":  "Замените модуль `%s` на `main`.",

	"error.nomainfunc": "Файл, который вы пытаетесь скомпилировать не содержит функцию `main`.",
	"note.nomainfunc":  `Добавьте функцию main.`,

	"info.foundmain":     "Найден модуль `main`",
	"info.foundmainfunc": "Найдена функция `main`",
}
