package token_docopt

// ================================ lexer for docopt language ===========
var (
  State_Prologue = `
  (?P<NEWLINE>\n)
  |(?P<SECTION>^Usage:) => state_Usage_Line
  |(?P<LINE_OF_TEXT>[^\n]+)
  `

  State_Usage = `
  (?P<NEWLINE>\n)
  |(?P<USAGE>^Usage:)
  |(?P<SECTION>^[A-Z][A-Za-z _-]+:) => state_Options
  |(?P<LONG_BLANK>[\t ]{2,}) => state_Usage_Line
  # skip single blank
  |([\t ])
  # Match some kind of comment when not preceded by LongBlank
  |(?P<LINE_OF_TEXT>[^\n]+)
  `

  State_Usage_Line = `
  (?P<NEWLINE>\n) => state_Usage
  |([\t ]+)
  |(?P<SHORT>-[A-Za-z0-9?])
  |(?P<LONG>--[A-Za-z][A-Za-z0-9_-]+|^--$)
  |(?P<ARGUMENT><[A-Za-z][A-Za-z0-9_-]*>|[A-Z_][A-Z0-9_-]+)
  # Punctuation doesn't accept comma but elipsis ...
  |(?P<PUNCT>[\][=()|]|\.{3})
  |(?P<IDENT>[A-Za-z][A-Za-z0-9_-]+)
  `

  State_Options = `
  (?P<NEWLINE>\n)
  # Options: is matched by state_Usage
  |(?P<SECTION>^[A-Z][A-Za-z _-]+:) => state_Free
  |(?P<DEFAULT>^default: )
  |(?P<LONG_BLANK>[\t ]{2,})
  # skip single blank
  |([\t ])
  |(?P<SHORT>-[A-Za-z0-9?])
  |(?P<LONG>--[A-Za-z][A-Za-z0-9_-]+|^--$)
  |(?P<ARGUMENT><[A-Za-z][A-Za-z0-9_-]+>|[A-Z_][A-Z0-9_-]+)
  # Punctuation differe from state_Usage accepts comma
  |(?P<PUNCT>[\][=,()|])
  # LINE_OF_TEXT not matching []
  |(?P<LINE_OF_TEXT>[^\n[\]]+)
  `

  State_Free = `
  (?P<NEWLINE>\n)
  |(?P<SECTION>^[A-Z][A-Za-z _-]+:)
  |(?P<LINE_OF_TEXT>[^\n]+)
  `

  All_states = map[string]string{
    "state_Prologue" : State_Prologue,
    "state_Usage" : State_Usage,
    "state_Usage_Line" : State_Usage_Line,
    "state_Options" : State_Options,
    "state_Free" : State_Free,
  }
)
