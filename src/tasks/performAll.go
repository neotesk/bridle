/*
    Bridle, a Javascript Project Manager
    Open-Source, WTFPL License.

    Copyright (C) 2025-20xx Neo <neotesk>
*/

package Tasks;

import (
	"fmt"
	"github.com/neotesk/bridle/internal/otsfile"
	"github.com/neotesk/bridle/internal/bridleots"
	"github.com/neotesk/bridle/internal/types"
);

var performAllTask = Types.CLITask {
    Name: "performAll",
    Action: func ( args Types.CLITaskArguments ) {
        fmt.Println( "Perform All" );
        parsed := OTSFile.Parse( "bridle.ots" );
        fmt.Println( BridleOTS.ParseDocument( parsed ) );
        installDependenciesTask.Action( args );
    },
}