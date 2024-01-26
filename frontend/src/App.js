// App.js
import React, { useState } from 'react';
import CreateTask from './CreateTask';
import ListTask from './ListTask';
import useTaskManager from './useTaskManager';

const App = () => {
  const { tasks, markComplete, deleteTask } = useTaskManager();

  return (
    <div>
      <CreateTask />
      <ListTask tasks={tasks} markComplete={markComplete} deleteTask={deleteTask} />
    </div>
  );
};

export default App;
