/*
    Bridle, a Javascript Project Manager
    Open-Source, WTFPL License.

    Copyright (C) 2025-20xx Neo <neotesk>
*/

package OTSFile;

import "bufio";

type TokenKind int;

const (
    // Start
    T_EOF TokenKind = iota

    // Basics
    T_Identifier
    T_String
    T_Number

    // Parenthesis
    T_OpenParen
    T_CloseParen

    // Symbols
    T_Equals
);

var tokenNames = map [ TokenKind ] string {
    T_EOF: "EndOfFile",
    T_Identifier: "Identifier",
    T_String: "String",
    T_Number: "Number",
    T_OpenParen: "OpenParenthesis",
    T_CloseParen: "CloseParenthesis",
    T_Equals: "EqualsSign",
}

type Token struct {
    Kind TokenKind;
    Value string;
    Position Position;
}

func ( tok Token ) GetType () string {
    return tokenNames[ tok.Kind ];
}

func ( tok TokenKind ) Name () string {
    return tokenNames[ tok ];
}

func ( tok Token ) Debug () {
    println( tokenNames[ tok.Kind ] + "\t->\t" + tok.Value );
}

type Position struct {
	Line int
	Column int
}

type Tokenizer struct {
	position Position
	reader *bufio.Reader
	filename string
}

type OTSObject struct {
    Items map [ any ] any
    Length int
}

type OTSDescriptor struct {
    Name string
    Item OTSObject
}

type OTSDocument struct {
    Items []OTSDescriptor
}