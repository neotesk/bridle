/*
    Bridle, a Javascript Project Manager
    Open-Source, WTFPL License.

    Copyright (C) 2025-20xx Neo <neotesk>
*/

package Helpers;

import (
    "github.com/neotesk/bridle/internal/cli"
)

func Item [ T any ] ( arr []T, idx int ) any {
    if len( arr ) <= idx {
        return nil;
    }
    return arr[ idx ];
}

func ItemCoalesce [ T any ] ( arr []T, idx int, def T ) T {
    if len( arr ) <= idx {
        return def;
    }
    return arr[ idx ];
}

func Make [ T any ] ( thing any ) T {
    output, ok := thing.( T );
    if !ok {
        CLI.Die( "Fatal Error: Cannot convert object into desired type.\n" );
    }
    return output;
}

func MakeCoalesce [ T any ] ( thing any, def T ) T {
    if thing == nil {
        return def;
    }
    output, ok := thing.( T );
    if !ok {
        CLI.Die( "Fatal Error: Cannot convert object into desired type.\n" );
    }
    return output;
}