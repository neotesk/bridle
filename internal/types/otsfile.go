/*
    Bridle, a Javascript Project Manager
    Open-Source, WTFPL License.

    Copyright (C) 2025-20xx Neo <neotesk>
*/

package Types;

type StringToken struct {
    Type string
    Value string
}

type NumericToken struct {
    Type string
    Value float64
}

type AnyToken struct {
    Type string
    Value any
}