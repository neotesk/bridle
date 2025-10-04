/*
    Bridle, a Javascript Project Manager
    Open-Source, WTFPL License.

    Copyright (C) 2025-20xx Neo <neotesk>
*/

package Types;

type CLITaskArguments struct {
    CWD string
}

type CLITask struct {
    Name string
    Action func ( args CLITaskArguments )
}