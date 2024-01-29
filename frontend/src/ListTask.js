// ListTask.js
import React from 'react';
import { Grid, Button } from '@mui/material';

const ListTask = ({ tasks, markComplete, deleteTask }) => {
  return (
    <Grid container spacing={2}>
      {tasks.map((task) => (
        <Grid item key={task.id} xs={12} sm={6} md={4}>
          <div>
            <h3>{task.taskName}</h3>
            <p>{task.description}</p>
            <p>Status: {task.status}</p>
            <Button variant="outlined" color="primary" onClick={() => markComplete(task.id)}>
              Mark as Complete
            </Button>
            <Button variant="outlined" color="secondary" onClick={() => deleteTask(task.id)}>
              Delete Task
            </Button>
          </div>
        </Grid>
      ))}
    </Grid>
  );
};

export default ListTask;
