/*
    Bridle, a Javascript Project Manager
    Open-Source, WTFPL License.

    Copyright (C) 2025-20xx Neo <neotesk>
*/

package Types;

// This is for having a record of arguments and flags
type Argument struct {
    Name string;
    ShortDesc string;
    DefaultValue string;
}

type Flag struct {
    Name string;
    ShortDesc string;
    DefaultValue bool;
}

// This is for storing default arguments
type DefaultArgs struct {
    Arguments []Argument;
    Flags []Flag;
}

// This is for exporting arguments
type ArgumentsList struct {
    Arguments map[ string ] string;
    Flags map[ string ] bool;
    Keywords []string;
}