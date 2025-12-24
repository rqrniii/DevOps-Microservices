// api/todo.ts
import todoApi from "./todoApi";
import axios from "axios";

const getToken = () => localStorage.getItem("token") || "";

export const getTodos = () => {
  return todoApi.get("", {
    headers: {
      Authorization: `Bearer ${getToken()}`,
    },
  });
};

export const addTodo = (task: string) => {
  return todoApi.post(
    "",
    { task },
    {
      headers: {
        Authorization: `Bearer ${getToken()}`,
      },
    }
  );
};

export const toggleTodo = (id: number) => {
  return todoApi.put(
    `/${id}/toggle`,
    {},
    {
      headers: {
        Authorization: `Bearer ${getToken()}`,
      },
    }
  );
};

export const addAITasks = async (prompt: string): Promise<void> => {
  const token = getToken();

  const aiBaseUrl = import.meta.env.VITE_AI_API || "http://localhost:8082";
  // 1️⃣ Call AI service
  const aiResponse = await axios.post(
    `${aiBaseUrl}/generate`,
    { prompt },
    { headers: { Authorization: `Bearer ${token}` } }
  );

  const tasks: string[] = aiResponse.data.tasks.map((t: string) =>
    t.replace(/^\d+\.\s*/, "").trim()
  );

  // 2️⃣ Send to backend
  await todoApi.post(
    "/ai",
    { tasks },
    { headers: { Authorization: `Bearer ${token}` } }
  );
};

export const deleteTodo = (id: number) => {
  return todoApi.delete(`/${id}`, {
    headers: {
      Authorization: `Bearer ${getToken()}`,
    },
  });
};