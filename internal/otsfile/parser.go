/*
    Bridle, a Javascript Project Manager
    Open-Source, WTFPL License.

    Copyright (C) 2025-20xx Neo <neotesk>
*/

package OTSFile;

import (
    "os"
    "strings"
    "strconv"
    "github.com/neotesk/bridle/internal/cli"
);

func expect ( filename string, token Token, kind TokenKind ) {
    if ( token.Kind != kind ) {
        CLI.Die( "Fatal Error! Expected token kind '%s' but found '%s' in file %s:%d:%d\n", kind.Name(), token.GetType(), filename, token.Position.Line, token.Position.Column );
    }
}

func expectMulti ( filename string, token Token, kind... TokenKind ) {
    for _, realKind := range kind {
        if ( token.Kind == realKind ) {
            return;
        }
    }
    kinds := []string {};
    for _, k := range kind {
        kinds = append( kinds, "'" + k.Name() + "'" );
    }
    CLI.Die( "Fatal Error! Expected %s but found '%s' in file %s:%d:%d\n", strings.Join( kinds, " or " ), token.GetType(), filename, token.Position.Line, token.Position.Column );
}

func item ( arr *[]Token, idx int ) Token {
    if idx >= len( *arr ) {
        CLI.Die( "Fatal Error! Expected a token but reached end of line.\n" );
    }
    return ( *arr )[ idx ];
}

func parseArray ( filename string, _tk *[]Token, idx *int ) OTSObject {
    // Expect an open parenthesis
    expect( filename, item( _tk, *idx ), T_OpenParen );
    *idx++;

    // Create a data object
    data := OTSObject {
        Items: map [ any ] any {},
        Length: 0,
    };

    // Read Insides
    for {
        tok := item( _tk, *idx );
        if tok.Kind == T_CloseParen {
            *idx++;
            break;
        }
        nextTok := item( _tk, *idx + 1 );
        if nextTok.Kind == T_Equals || ( nextTok.Kind == T_Identifier && nextTok.Value == "is" ) {
            expectMulti( filename, tok, T_Identifier, T_String );
            *idx += 2;
            data.Items[ tok.Value ] = parseExpression( filename, _tk, idx );
        } else {
            data.Items[ data.Length ] = parseExpression( filename, _tk, idx );
            *idx++;
            data.Length++;
        }
    }

    return data;
}

func parseExpression ( filename string, _tk *[]Token, idx *int ) any {
    // Note down certain variables
    tk := *_tk

    // Get the current token
    curTok := tk[ *idx ];

    // Check if it's a string or an Identifier
    switch curTok.Kind {
        case T_String, T_Identifier:
            *idx++;
            return curTok.Value;
        case T_Number:
            *idx++;
            return CLI.HandleError( strconv.Atoi( curTok.Value ) );
        case T_OpenParen:
            return parseArray( filename, _tk, idx );
        default:
            CLI.Die( "Fatal Error! Invalid token '%s' at %s:%d:%d (Expected an expression)\n", curTok.GetType(), filename, curTok.Position.Line, curTok.Position.Column );
    }

    return nil;
}

func Parse ( filename string ) OTSDocument {
    file := CLI.HandleError( os.Open( filename ) );
    tk := newTokenizer( file, filename );

    tokens := []Token {};

    for {
        tok := tk.Tokenize()
        if tok.Kind == T_EOF {
            break
        }
        tokens = append( tokens, tok );
    }

    tkIndex := 0;
    tkLen := len( tokens );

    data := OTSDocument {
        Items: []OTSDescriptor {},
    };

    for {
        curTok := tokens[ tkIndex ];

        // Expect an identifier
        expect( filename, curTok, T_Identifier );
        tkIndex++;

        // Expect an array
        val := parseArray( filename, &tokens, &tkIndex );

        // Put it inside the data
        data.Items = append( data.Items, OTSDescriptor {
            Name: curTok.Value,
            Item: val,
        } );

        if tkIndex >= tkLen {
            break;
        }
    }

    return data;
}