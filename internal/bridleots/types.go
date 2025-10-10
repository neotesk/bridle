/*
    Bridle, a Javascript Project Manager
    Open-Source, WTFPL License.

    Copyright (C) 2025-20xx Neo <neotesk>
*/

package BridleOTS;

import "github.com/neotesk/bridle/internal/otsfile"

type BridleProject struct {
    Name string
    Description string
    Version string
    Source string
    Author string
    License string
};

type BridleSettings struct {
    DependenciesPath string
};

type BridleDependencies map [ string ] string;

type BridleAction map [ string ] any;
type BridleActions map [ string ] BridleAction;

type BridleOperation struct {
    OperationName string
    Description OTSFile.OTSDescriptor
}

type BridleOperations []BridleOperation

type BridleDocument struct {
    Project BridleProject
    Operations BridleOperations
    Settings BridleSettings
    Dependencies BridleDependencies
    Actions BridleActions
}