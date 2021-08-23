grammar Fate;

fate
	: block EOF
	;

block
	: stat*
	;

stat
	: assignstat
	| ifstat
	| returnstat
	;

assignstat
	: variable '=' expr
	;

returnstat
	: 'return'
	;

ifstat
	: 'if' expr 'then' block ('elseif' expr 'then' block)* ('else' block)? 'end';

expr
	: value                                                                 # valueExpr
	| variable (('.' FUNC '()')+)?                                          # varFuncExpr
	| op=('!'|'-') expr                                                     # unaryExpr
	| '(' expr ')'                                                          # parens
	| '[' expr (',' expr)* ']'                                              # arrayExpr
	| expr op=('*'|'/'|'%') expr                                            # mulDiv
	| expr op=('+'|'-') expr                                                # addSub
	| expr op=('=='|'!='|'>'|'<'|'>='|'<='|'regexpMatch') expr              # simpleBooleanExpr
	| expr op=('contains'|'!contains') expr                                 # setBooleanExpr
	| expr '&&' expr                                                        # andExpr
    | expr '||' expr                                                        # orExpr
    ;

variable
	: field ('.' field)*
	;

field
	:  NAME (index)*
	;

value
	: STRING    # string
	| INT       # int
	| FLOAT     # float
	| TRUE      # true
	| FALSE     # false
	| NIL       # nil
	;

index: '[' INT ']';

LINE_COMMENT
    :   '//' ~[\r\n]* -> skip
    ;
COMMENT
    :   '/*' .*? '*/' -> channel(HIDDEN)
    ;

TRUE: 'true';
FALSE: 'false';
NIL: 'nil';
FUNC: 'count'|'sum'|'avg'|'max'|'min';

NAME
    : [a-zA-Z_][a-zA-Z_0-9]*
    ;
STRING
    : '\'' .*? '\''
    | '"' .*? '"'
    ;

INT: DIGIT+;
FLOAT: DIGIT+ '.' DIGIT*;
DIGIT: [0-9];

WS: [ \t\r\n]+ -> skip;
