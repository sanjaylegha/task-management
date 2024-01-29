// CreateTask.js
import React from 'react';
import { useForm } from 'react-hook-form';
import { TextField, Button, FormControl, InputLabel, Select, MenuItem } from '@mui/material';
import useTaskManager from './useTaskManager';

const CreateTask = () => {
  const { addTask } = useTaskManager();
  const { register, handleSubmit } = useForm();

  const onSubmit = (data) => {
    addTask(data);
    // Handle form submission logic here
    console.log(data);
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <TextField {...register('taskName')} label="Task Name" required />
      <TextField {...register('description')} label="Description" required />
      <FormControl>
        <InputLabel>Status</InputLabel>
        <Select {...register('status')} required>
          <MenuItem value="pending">Pending</MenuItem>
          <MenuItem value="completed">Completed</MenuItem>
        </Select>
      </FormControl>
      <Button type="submit" variant="contained" color="primary">
        Create Task
      </Button>
    </form>
  );
};

export default CreateTask;
