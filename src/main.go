/*
   Bridle, a Javascript Project Manager
   Open-Source, WTFPL License.

   Copyright (C) 2025-20xx Neo <neotesk>
*/

package main

import (
	"fmt"
	"os"

	"github.com/neotesk/bridle/internal/cli"
	"github.com/neotesk/bridle/internal/helpers"
	"github.com/neotesk/bridle/internal/types"
);

var version = "github-flavor";

func main () {
    // Setup default argument parameters
    defaultArgs := Types.DefaultArgs {
        Flags: []Types.Flag {
            {
                Name: "h",
                ShortDesc: "Prints Help Page",
                DefaultValue: false,
            },
            {
                Name: "d",
                ShortDesc: "Enables Debug Mode",
                DefaultValue: false,
            },
            {
                Name: "v",
                ShortDesc: "Prints version",
                DefaultValue: false,
            },
            {
                Name: "s",
                ShortDesc: "Enables Silent Mode",
                DefaultValue: false,
            },
        },
        Arguments: []Types.Argument {},
    };

    // After setting the default arguments, let's
    // feed them inside the Arguments function
    // so we can get a good output of what we have
    // in our hands.
    argsList := CLI.Arguments( defaultArgs );
    CLI.IsDebug = argsList.Flags[ "d" ];

    // Check if we should print the version
    if ( argsList.Flags[ "v" ] ) {
        fmt.Println( version );
    }

    // Check if we should print the help page
    if ( argsList.Flags[ "h" ] ) {
        flags := "-";
        tasks := []string {
            "performAll",
            "installDependencies",
            "initProject",
            "newProject",
        };
        _tasks := "";
        for _, key := range defaultArgs.Flags {
            flags += key.Name;
        }
        for _, key := range tasks {
            _tasks += "   " + key + "\n";
        }
        fmt.Printf(
            "%s %s %s %s\n%s\n%s%s\n   %s\n",
            CLI.Colorify( "Usage:", "fcba03" ),
            CLI.Colorify( "bridle", "32a852" ),
            CLI.Colorify( "[" + flags + "]", "7a7a7a" ),
            CLI.Colorify( "[taskName]", "7a7a7a" ),
            CLI.Colorify( "Available Tasks:", "ab86c2" ),
            _tasks,
            CLI.Colorify( "More info:", "ab86c2" ),
            CLI.Colorify( "https://github.com/neotesk/bridle/wiki", "3283a8" ),
        );
    }

    // Since I want to combine page flags (like v and h)
    // I wanna make sure the program exits afterwards
    if ( argsList.Flags[ "v" ] || argsList.Flags[ "h" ] ) {
        os.Exit( 1 );
    }

    // Get the command
    command := Helpers.ItemCoalesce( argsList.Keywords, 0, "performAll" );
    fmt.Println( command );
}