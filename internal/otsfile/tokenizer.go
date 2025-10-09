/*
    Bridle, a Javascript Project Manager
    Open-Source, WTFPL License.

    Copyright (C) 2025-20xx Neo <neotesk>
*/

package OTSFile;

import (
	"bufio"
	"io"
	"unicode"
	"github.com/neotesk/bridle/internal/cli"
);

func newTokenizer ( reader io.Reader, filename string ) *Tokenizer {
    return &Tokenizer{
        position: Position { Line: 1, Column: 0 },
        reader: bufio.NewReader( reader ),
        filename: filename,
    };
}

func ( tk *Tokenizer ) iterPos () {
    tk.position.Line++;
    tk.position.Column = 0;
}

func ( tk *Tokenizer ) revIter () {
    if err := tk.reader.UnreadRune(); err != nil {
        panic( err );
    }
    tk.position.Column--;
}

func ( tk *Tokenizer ) Tokenize () Token {
    // Used for iterating a string
    shouldIterString := false;
    shouldIterNumeric := false;
    shouldIterIdentifier := false;
    iterStartPos := Position {};
    iterRecord := "";

    // Iteration
    for {
        r, _, err := tk.reader.ReadRune();
        if err != nil {
            if err == io.EOF {
                return Token { Kind: T_EOF, Value: "", Position: tk.position };
            }
            // This is a different error.
            panic( err );
        }

        // Iterate the column
        tk.position.Column++;

        // Check if we are iterating a string
        if shouldIterString {
            switch r {
                case '"':
                    return Token { Kind: T_String, Value: iterRecord, Position: iterStartPos };
                default:
                    if r == '\n' {
                        tk.iterPos();
                    }
                    iterRecord = iterRecord + string( r );
                    continue;
            }
        }

        // Check if we are iterating a numeric
        if shouldIterNumeric {
            if unicode.IsDigit( r ) {
                iterRecord = iterRecord + string( r );
                continue;
            } else {
                tk.revIter();
                return Token { Kind: T_Number, Value: iterRecord, Position: iterStartPos };
            }
        }

        // Check if we are iterating an identifer
        if shouldIterIdentifier {
            if unicode.IsLetter( r ) || unicode.IsDigit( r ) {
                iterRecord = iterRecord + string( r );
                continue;
            } else {
                tk.revIter();
                return Token { Kind: T_Identifier, Value: iterRecord, Position: iterStartPos };
            }
        }

        switch r {
            // Ignore new line
            case '\n':
                tk.iterPos();

            // Symbols
            case '(':
                return Token { Kind: T_OpenParen, Value: "(", Position: tk.position };
            case ')':
                return Token { Kind: T_CloseParen, Value: ")", Position: tk.position };
            case '=':
                return Token { Kind: T_Equals, Value: "=", Position: tk.position };

            // String Definition
            case '"':
                shouldIterString = true;
                iterStartPos = tk.position;
                continue;

            // Identifiers, numbers and other
            default:
                if unicode.IsSpace( r ) {
                    continue;
                } else if unicode.IsDigit( r ) {
                    shouldIterNumeric = true;
                    iterStartPos = tk.position;
                    iterRecord = iterRecord + string( r );
                    continue;
                } else if unicode.IsLetter( r ) || r == '_' {
                    shouldIterIdentifier = true;
                    iterStartPos = tk.position;
                    iterRecord = iterRecord + string( r );
                    continue;
                } else {
                    CLI.Die( "Fatal Error! Unknown token '%s' in file %s:%d:%d\n", string( r ), tk.filename, tk.position.Line, tk.position.Column );
                }
        }
    }
}