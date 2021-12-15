Prism.languages.hike = Prism.languages.extend("clike", {
  keyword:
    /\b(?:str|new|runtime|break|for|while|if|else|co_wait|let|const|import|bool|try|test|except|export|as|pub|int64|int32|int16|int8|int|module|ret|class|struct|co)\b/,
  builtin: /\b(?:print|nil|scan|this|CoroutinesMemoryOverflow)\b/,
  boolean: /\b(?:true|false)\b/,
  operator:
    /(==|!=|<=|>=|<|>|&&|\|\||!|=|\+\=|\-\=|\*\=|\/\=|\+|\-|\*|\/|%|\^|\.\.|\-\-|\+\+|\/\%|\/\%=|\^=|<=>)/,
  number: /(?:\b\d+(\.\d+)?\b)|(\b([0-9]+|\?)[gbci]\b)/,
  string: /[a-z]?"(?:\\.|[^\\"])*"|'(?:\\.|[^\\'])*'/,
  tag: /@([a-zA-Z_][a-zA-Z0-9_]*)/,
})
delete Prism.languages.spwn["class-name"]