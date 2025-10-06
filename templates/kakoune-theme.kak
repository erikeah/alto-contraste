## code
face global value blue+b
face global type bright-blue
face global variable default
face global module cyan
face global function magenta
face global string bright-green
face global keyword red
face global operator bright-green+b
face global attribute rgb:{{ .Orange.ToHex3 }}
face global comment bright-black+i
face global documentation comment
face global meta cyan
face global builtin default+b


## markup
face global title rgb:{{ .Orange.ToHex3 }}+bu
face global header bright-cyan+bu
face global mono green
face global block magenta
face global link bright-cyan
face global bullet cyan
face global list yellow


## builtin
face global Default default,default
face global PrimarySelection blue,default+frb
face global SecondarySelection bright-black,bright-white+fb
face global PrimaryCursor default,default+fbr
face global SecondaryCursor bright-white,bright-black+fg
face global PrimaryCursorEol default,red+g
face global SecondaryCursorEol default,red+g
face global LineNumbers bright-black,default
face global LineNumberCursor rgb:{{ .Orange.ToHex3 }},default+b
face global LineNumbersWrapped white+i
face global MenuForeground PrimarySelection
face global MenuBackground default,bright-white+fg
face global Information bright-blue+br
face global InlineInformation white+fi
face global Error red+br
face global DiagnosticError red
face global DiagnosticWarning rgb:{{ .Orange.ToHex3 }}
face global StatusLine black,bright-white
face global StatusCursor default,default+bu
face global Prompt rgb:{{ .Orange.ToHex3 }}+bf
face global MatchingChar default,bright-white+b
face global BufferPadding cyan
face global Whitespace bright-black+f

## kak-tree-sitter
face global ts_attribute attribute
face global ts_comment comment
# face global ts_comment_block
# face global ts_comment_line
face global ts_conceal yellow+i
face global ts_constant blue
# face global ts_constant_builtin_boolean
# face global ts_constant_character
# face global ts_constant_character_escape
# face global ts_constant_macro
# face global ts_constant_numeric
# face global ts_constant_numeric_float
# face global ts_constant_numeric_integer
face global ts_constructor bright-red
face global ts_diff_plus bright-green
face global ts_diff_minus bright-red
face global ts_diff_delta black
# face global ts_diff_delta_moved
face global ts_error Error
face global ts_function function
face global ts_function_builtin +b@function
# face global ts_function_macro
# face global ts_function_method
# face global ts_function_method_private
# face global ts_function_special
face global ts_hint bright-yellow
face global ts_info bright-blue
face global ts_keyword keyword
# face global ts_keyword_control
# face global ts_keyword_conditional
# face global ts_keyword_control_conditional
# face global ts_keyword_control_directive
face global ts_keyword_control_import magenta
# face global ts_keyword_control_repeat
face global ts_keyword_control_return cyan
# face global ts_keyword_control_except
# face global ts_keyword_control_exception
# face global ts_keyword_directive
# face global ts_keyword_function
# face global ts_keyword_operator
# face global ts_keyword_special
# face global ts_keyword_storage
# face global ts_keyword_storage_modifier
# face global ts_keyword_storage_modifier_mut
# face global ts_keyword_storage_modifier_ref
# face global ts_keyword_storage_type
face global ts_label bright-green
face global ts_markup_bold +b
face global ts_markup_heading title
face global ts_markup_heading_1 header
# face global ts_markup_heading_2
# face global ts_markup_heading_3
# face global ts_markup_heading_4
# face global ts_markup_heading_5
# face global ts_markup_heading_6
# face global ts_markup_heading_marker
face global ts_markup_italic +i
face global ts_markup_list_checked bright-black+s
face global ts_markup_list_numbered default
face global ts_markup_list_unchecked default
face global ts_markup_list_unnumbered default
face global ts_markup_link_label green
face global ts_markup_link_url cyan+u
face global ts_markup_link_uri cyan+u
face global ts_markup_link_text bright-red
face global ts_markup_quote cyan+i
face global ts_markup_raw bright-green
# face global ts_markup_raw_block
# face global ts_markup_raw_inline
face global ts_markup_strikethrough +s
face global ts_namespace module
face global ts_operator operator
face global ts_property bright-blue
face global ts_punctuation default
# face global ts_punctuation_bracket
# face global ts_punctuation_delimiter
# face global ts_punctuation_special
face global ts_special blue
face global ts_spell yellow
face global ts_string string
# face global ts_string_regex
# face global ts_string_regexp
# face global ts_string_escape
# face global ts_string_special
# face global ts_string_special_path
# face global ts_string_special_symbol
# face global ts_string_symbol
face global ts_tag red
face global ts_tag_error +c@ts_tag
face global ts_text default
# face global ts_text_title
face global ts_type type
face global ts_type_builtin +b@type
# face global ts_type_enum_variant
face global ts_variable variable
face global ts_variable_builtin +b@variable
# face global ts_variable_other_member
# face global ts_variable_other_member_private
face global ts_variable_parameter rgb:{{ .AltOrange.ToHex3 }}
face global ts_warning DiagnosticWarning
