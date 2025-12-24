import axios from "axios";

const todoApi = axios.create({
  baseURL: import.meta.env.VITE_TODO_API || "/api/todos",
});

export default todoApi;