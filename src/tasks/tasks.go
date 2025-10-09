/*
    Bridle, a Javascript Project Manager
    Open-Source, WTFPL License.

    Copyright (C) 2025-20xx Neo <neotesk>
*/

package Tasks;

import (
	"github.com/neotesk/bridle/internal/cli"
	"github.com/neotesk/bridle/internal/types"
);

var tasks = []Types.CLITask {
    performAllTask,
    installDependenciesTask,
}

func mapTasks () map [ string ] Types.CLITask {
    taskList := map [ string ] Types.CLITask {}
    for _, task := range tasks {
        taskList[ task.Name ] = task;
    }
    return taskList;
}

var TaskList = mapTasks();

func GetTask ( taskName string ) Types.CLITask {
    task, exists := TaskList[ taskName ];
    if !exists {
        CLI.Die( "Fatal Error! Given task with the name %s does not exist.\n", taskName );
    }
    return task;
}

func RunTask ( taskName string, args Types.CLITaskArguments ) {
    GetTask( taskName ).Action( args );
}