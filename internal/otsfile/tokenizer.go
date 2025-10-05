/*
    Bridle, a Javascript Project Manager
    Open-Source, WTFPL License.

    Copyright (C) 2025-20xx Neo <neotesk>
*/

package OTSFile;

import (
	"os"
	"strings"

	CLI "github.com/neotesk/bridle/internal/cli"
	"github.com/neotesk/bridle/internal/types"
);

var validAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_";
var validNumbers = "1234567890";
var validSymbols = "()@=";
var validIdentifierNext = validAlphabet + validNumbers;
var validIgnore = " \n\t";

func tokenizeIdentifier ( idx int, charList []string ) ( int, Types.StringToken ) {
    output := Types.StringToken {
        Type: "Identifier",
        Value: "",
    };
    curIdx := idx;
    for strings.Contains( validIdentifierNext, charList[ curIdx ] ) {
        output.Value = output.Value + charList[ curIdx ];
        curIdx += 1;
    }
    return curIdx - idx, output;
}

func tokenizeNumber ( idx int, charList []string ) ( int, Types.NumericToken ) {
    output := Types.NumericToken {
        Type: "Number",
        Value: 0,
    };
    output_temp := "";
    curIdx := idx;
    char := charList[ curIdx ];
    for strings.Contains( validNumbers,  char ) || ( char == "." && !strings.Contains( output_temp, "." ) ) {
        output_temp = output_temp + char;
        curIdx += 1;
        char = charList[ curIdx ];
    }
    return curIdx - idx, output;
}

func tokenizeString ( idx int, l int, r int, charList []string ) ( int, int, int, Types.StringToken ) {
    curIdx := idx + 1;
    curChar := charList[ curIdx ];
    output := Types.StringToken {
        Type: "String",
        Value: "",
    };
    for curChar != "\"" {
        if curChar == "\n" {
            l += 1;
            r = 1;
        }
        output.Value = output.Value + curChar;
        curIdx += 1;
        curChar = charList[ curIdx ];
    }
    curIdx += 1;
    return curIdx - idx, l, r, output;
}

func Tokenize ( data string, filename string ) []any {
    charList := strings.Split( data, "" );
    idx := 0;
    line := 1;
    row := 1;
    charLen := len( charList );
    tokenList := []any {};
    for idx < charLen {
        curChar := charList[ idx ];
        var nextIdx = 0;
        var token any;
        if strings.Contains( validAlphabet, curChar ) {
            nextIdx, token = tokenizeIdentifier( idx, charList );
        } else if strings.Contains( validNumbers, curChar ) {
            nextIdx, token = tokenizeNumber( idx, charList );
        } else if strings.Contains( validSymbols, curChar ) {
            nextIdx = 1;
            token = Types.StringToken {
                Type: "Symbol",
                Value: curChar,
            };
        } else if curChar == "\"" {
            nextIdx, line, row, token = tokenizeString( idx, line, row, charList );
        } else if strings.Contains( validIgnore, curChar ) {
            idx += 1;
            if curChar == "\n" {
                row = 1;
                line += 1;
            } else {
                row += 1;
            }
            continue;
        }
        if nextIdx == 0 {
            CLI.ErrPrintf( "Fatal Error! Unknown token '%s' in file %s:%d:%d\n", curChar, filename, line, row );
            os.Exit( 1 );
        }
        idx += nextIdx;
        row += nextIdx;
        tokenList = append( tokenList, token );
    }
    return tokenList;
}