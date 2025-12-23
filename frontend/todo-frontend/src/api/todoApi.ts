import axios from "axios";

const todoApi = axios.create({
  baseURL: import.meta.env.VITE_TODO_API || "http://localhost:8081",
});

export default todoApi;