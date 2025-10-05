/*
    Bridle, a Javascript Project Manager
    Open-Source, WTFPL License.

    Copyright (C) 2025-20xx Neo <neotesk>
*/

package Tasks;

import (
	"fmt"
	"github.com/neotesk/bridle/internal/types"
);

var installDependenciesTask = Types.CLITask {
    Name: "installDependencies",
    Action: func ( args Types.CLITaskArguments ) {
        fmt.Println( "Install Dependencies" );
    },
}