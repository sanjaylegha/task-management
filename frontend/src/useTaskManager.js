// src/useTaskManager.js
import { useState, useEffect } from 'react';
import axios from 'axios';

const useTaskManager = () => {
  const [tasks, setTasks] = useState([]);

  // Fetch tasks from the server on component mount
  useEffect(() => {
    fetchTasks();
  }, []);

  const fetchTasks = async () => {
    try {
      const response = await axios.get('http://localhost:8080/api/tasks/list'); // Replace with your API endpoint
      setTasks(response.data || []);
    } catch (error) {
      console.error('Error fetching tasks:', error);
    }
  };

  const addTask = async (newTask) => {
    try {
      const response = await axios.post('http://localhost:8080/api/tasks/create', newTask); // Replace with your API endpoint
      setTasks([...tasks, response.data]);
    } catch (error) {
      console.error('Error adding task:', error);
    }
  };

  const markComplete = async (taskId) => {
    try {
      await axios.put(`http://localhost:8080/api/tasks/update-status?id=${taskId}`)
    } catch (error) {
      console.error('Error updating task status')
    }
  };

  const deleteTask = async (taskId) => {
    try {
      await axios.delete(`http://localhost:8080/api/tasks/delete?id=${taskId}`)
    } catch (error){
      console.error('Error deleting task')
    }
  };

  return { tasks, addTask, markComplete, deleteTask };
};

export default useTaskManager;
