/*
    Bridle, a Javascript Project Manager
    Open-Source, WTFPL License.

    Copyright (C) 2025-20xx Neo <neotesk>
*/

package Helpers;

import "strings";

func StringChunk ( str string, totalLength int ) string {
    chunks := strings.Split( str, " " );
    newText := "";
    currentText := []string{};
    for _, chunk := range chunks {
        currentText = append( currentText, chunk );
        if len( strings.Join( currentText, " " ) ) > totalLength {
            newText = newText + strings.Join( currentText[ :len( currentText ) - 1 ], " " ) + "\n";
            currentText = []string{ chunk };
        }
    }
    return newText + strings.Join( currentText, " " );
}