/*
    Bridle, a Javascript Project Manager
    Open-Source, WTFPL License.

    Copyright (C) 2025-20xx Neo <neotesk>
*/

package Helpers;

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