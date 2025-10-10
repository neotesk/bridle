/*
    Bridle, a Javascript Project Manager
    Open-Source, WTFPL License.

    Copyright (C) 2025-20xx Neo <neotesk>
*/

package BridleOTS;

import (
    "strings"
    "github.com/neotesk/bridle/internal/otsfile"
    "github.com/neotesk/bridle/internal/helpers"
    "github.com/neotesk/bridle/internal/cli"
)

func lookupDescriptor ( doc OTSFile.OTSDocument, name string ) OTSFile.OTSDescriptor {
    for _, item := range doc.Items {
        if item.Name == name {
            return item;
        }
    }
    CLI.Die( "Fatal error! Cannot find the referenced object '%s'\n", name );
    return OTSFile.OTSDescriptor {}; /* Need this to supress the error, it doesn't do anything */
}

func ParseDocument ( doc OTSFile.OTSDocument ) BridleDocument {
    output := BridleDocument {
        Project: BridleProject {
            Name: "null",
            Description: "null",
            Version: "null",
            Source: "null",
            Author: "null",
            License: "null",
        },
        Operations: BridleOperations {},
        Settings: BridleSettings {
            DependenciesPath: "@libs",
        },
        Dependencies: BridleDependencies {},
        Actions: BridleActions {},
    };

    for _, item := range doc.Items {
        switch item.Name {
            case "project":
                output.Project = BridleProject {
                    Name: Helpers.MakeCoalesce( item.Item.Items[ "name" ], "null" ),
                    Description: Helpers.MakeCoalesce( item.Item.Items[ "description" ], "null" ),
                    Version: Helpers.MakeCoalesce( item.Item.Items[ "version" ], "null" ),
                    Source: Helpers.MakeCoalesce( item.Item.Items[ "source" ], "null" ),
                    Author: Helpers.MakeCoalesce( item.Item.Items[ "author" ], "null" ),
                    License: Helpers.MakeCoalesce( item.Item.Items[ "license" ], "null" ),
                };
            case "operation":
                objName := Helpers.MakeCoalesce( item.Item.Items[ 1 ], "_bundleTask" );
                output.Operations = append( output.Operations, BridleOperation {
                    OperationName: Helpers.MakeCoalesce( item.Item.Items[ 0 ], "bundle" ),
                    Description: lookupDescriptor( doc, objName ),
                } );
            case "dependencies":
                if item.Item.Length > 0 {
                    CLI.Die( "Fatal Error! Dependencies object cannot contain non-keyed elements." );
                }
                for key, value := range item.Item.Items {
                    output.Dependencies[ Helpers.Make[ string ]( key ) ] = Helpers.Make[ string ]( value );
                }
            case "defineAction":
                action := BridleAction {};
                for key, value := range item.Item.Items {
                    action[ Helpers.Make[ string ]( key ) ] = value;
                }
                output.Actions[ Helpers.Make[ string ]( action[ "name" ] ) ] = action;
            case "settings":
                output.Settings = BridleSettings {
                    DependenciesPath: Helpers.MakeCoalesce( item.Item.Items[ "dependenciesPath" ], "@libs" ),
                };
            default:
                if !strings.HasPrefix( item.Name, "_" ) {
                    CLI.Die( "Fatal Error! Unknown descriptor '%s' (If you want to use this object as a reference, put an underscore at the start.)\n", item.Name );
                }
        }
    }

    return output;
}