/*
    Bridle, a Javascript Project Manager
    Open-Source, WTFPL License.

    Copyright (C) 2025-20xx Neo <neotesk>
*/

package BridleOTS;

import (
    "github.com/neotesk/bridle/internal/otsfile"
    "github.com/neotesk/bridle/internal/helpers"
)

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
                }
        }
    }

    return output;
}