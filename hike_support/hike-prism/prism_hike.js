Prism.languages.hike = Prism.languages.extend("clike", {
  keyword:
    /\b(?:export|pub|int|module|ret|class|struct|co)\b/,
  builtin: /\b(?:print|nil|scan|this)\b/,
  boolean: /\b(?:true|false)\b/,
  operator:
    /(==|!=|<=|>=|<|>|&&|\|\||!|=|\+\=|\-\=|\*\=|\/\=|\+|\-|\*|\/|%|\^|\.\.|\-\-|\+\+|\/\%|\/\%=|\^=|<=>)/,
  number: /(?:\b\d+(\.\d+)?\b)|(\b([0-9]+|\?)[gbci]\b)/,
  string: /[a-z]?"(?:\\.|[^\\"])*"|'(?:\\.|[^\\'])*'/,
  tag: /@([a-zA-Z_][a-zA-Z0-9_]*)/,
})
delete Prism.languages.spwn["class-name"]