/*
    Bridle, a Javascript Project Manager
    Open-Source, WTFPL License.

    Copyright (C) 2025-20xx Neo <neotesk>
*/

package Tasks

import (
	"fmt"
	"os"
	"github.com/neotesk/bridle/internal/cli"
	"github.com/neotesk/bridle/internal/otsfile"
	"github.com/neotesk/bridle/internal/types"
);

var performAllTask = Types.CLITask {
    Name: "performAll",
    Action: func ( args Types.CLITaskArguments ) {
        fmt.Println( "Perform All" );
        file := CLI.HandleError( os.ReadFile( "bridle.ots" ) );
        tokens := OTSFile.Tokenize( string( file ), "bridle.ots" );
        fmt.Println( tokens );
        installDependenciesTask.Action( args );
    },
}