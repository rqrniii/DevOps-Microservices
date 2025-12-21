import { useEffect, useState } from "react";
import { getTodos, addTodo } from "../api/todo";

export default function Todos() {
  const [todos, setTodos] = useState<any[]>([]);
  const [task, setTask] = useState("");

  const fetchTodos = async () => {
    const res = await getTodos();
    setTodos(res.data);
  };

  useEffect(() => {
    fetchTodos();
  }, []);

  const handleAdd = async () => {
    await addTodo(task);
    setTask("");
    fetchTodos();
  };

  return (
    <div>
      <h2>Todos</h2>

      <input
        value={task}
        onChange={(e) => setTask(e.target.value)}
      />
      <button onClick={handleAdd}>Add</button>

      <ul>
        {todos.map((t) => (
          <li key={t.id}>{t.task}</li>
        ))}
      </ul>
    </div>
  );
}